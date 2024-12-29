package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"Go_project/internal/core"
	redditpb "Go_project/internal/generated"
	"google.golang.org/grpc"
)

var (
	server = core.NewServer() // Shared server instance for both gRPC and REST
	mutex  sync.Mutex
)

func main() {
	// Run gRPC and REST servers concurrently
	go startGRPCServer()
	startRESTServer()
}

// gRPC Server
func startGRPCServer() {
	grpcServer := grpc.NewServer()
	redditpb.RegisterRedditServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

// REST Server
func startRESTServer() {
	r := gin.Default()

	// Middleware: Logging and Recovery
	r.Use(gin.Logger(), gin.Recovery())

	// User Routes
	r.POST("/users", registerUser)
	r.GET("/users/:username/feed", fetchFeed)

	// Subreddit Routes
	r.POST("/subreddits", createSubreddit)
	r.POST("/subreddits/:subreddit/join", joinSubreddit)

	// Post Routes
	r.POST("/subreddits/:subreddit/posts", createPost)
	r.POST("/posts/:postId/comments", addComment)

	log.Println("REST API server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run REST API server: %v", err)
	}
}

// REST Handlers

// registerUser registers a new user
func registerUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("registerUser: Invalid payload: %v", err)
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := server.Users[req.Username]; exists {
		log.Printf("registerUser: User '%s' already exists", req.Username)
		c.JSON(409, gin.H{"error": "User already exists"})
		return
	}

	server.Users[req.Username] = &core.User{
		Username:         req.Username,
		JoinedSubreddits: make(map[string]bool),
	}
	log.Printf("registerUser: User '%s' registered successfully", req.Username)
	c.JSON(201, gin.H{"message": "User registered successfully"})
}

// createSubreddit creates a new subreddit
func createSubreddit(c *gin.Context) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("createSubreddit: Invalid payload: %v", err)
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := server.Subreddits[req.Name]; exists {
		log.Printf("createSubreddit: Subreddit '%s' already exists", req.Name)
		c.JSON(409, gin.H{"error": "Subreddit already exists"})
		return
	}

	server.Subreddits[req.Name] = &core.Subreddit{
		Name:        req.Name,
		Description: req.Description,
		Members:     make(map[string]bool),
		Posts:       []*core.Post{},
	}
	log.Printf("createSubreddit: Subreddit '%s' created successfully", req.Name)
	c.JSON(201, gin.H{"message": "Subreddit created successfully"})
}

// joinSubreddit allows a user to join a subreddit
func joinSubreddit(c *gin.Context) {
	username := c.Query("username")
	subredditName := c.Param("subreddit")

	mutex.Lock()
	defer mutex.Unlock()

	subreddit, exists := server.Subreddits[subredditName]
	if !exists {
		log.Printf("joinSubreddit: Subreddit '%s' not found", subredditName)
		c.JSON(404, gin.H{"error": "Subreddit not found"})
		return
	}

	user, exists := server.Users[username]
	if !exists {
		log.Printf("joinSubreddit: User '%s' not found", username)
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	user.JoinedSubreddits[subredditName] = true
	subreddit.Members[username] = true
	log.Printf("joinSubreddit: User '%s' joined subreddit '%s'", username, subredditName)
	c.JSON(200, gin.H{"message": "User joined subreddit successfully"})
}

// createPost allows a user to create a post in a subreddit
func createPost(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Title    string `json:"title"`
		Content  string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("createPost: Invalid payload: %v", err)
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	subredditName := c.Param("subreddit")

	mutex.Lock()
	defer mutex.Unlock()

	subreddit, exists := server.Subreddits[subredditName]
	if !exists {
		log.Printf("createPost: Subreddit '%s' not found", subredditName)
		c.JSON(404, gin.H{"error": "Subreddit not found"})
		return
	}

	postID := fmt.Sprintf("%s-%d", subredditName, len(subreddit.Posts)+1)
	post := &core.Post{
		ID:        postID,
		Author:    req.Username,
		Title:     req.Title,
		Content:   req.Content,
		Subreddit: subredditName,
		CreatedAt: time.Now(),
	}

	subreddit.Posts = append(subreddit.Posts, post)
	server.Posts[postID] = post
	log.Printf("createPost: Post '%s' created successfully in subreddit '%s'", postID, subredditName)
	c.JSON(201, gin.H{"message": "Post created successfully", "postId": postID})
}

// addComment adds a comment to a post
func addComment(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Content  string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("addComment: Invalid payload: %v", err)
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	postID := c.Param("postId")

	mutex.Lock()
	defer mutex.Unlock()

	post, exists := server.Posts[postID]
	if !exists {
		log.Printf("addComment: Post '%s' not found", postID)
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	commentID := fmt.Sprintf("%s-comment-%d", postID, len(post.Comments)+1)
	comment := &core.Comment{
		ID:        commentID,
		Author:    req.Username,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}

	post.Comments = append(post.Comments, comment)
	server.Comments[commentID] = comment
	log.Printf("addComment: Comment '%s' added successfully to post '%s'", commentID, postID)
	c.JSON(201, gin.H{"message": "Comment added successfully", "commentId": commentID})
}

// fetchFeed retrieves the feed for a user
func fetchFeed(c *gin.Context) {
	username := c.Param("username")

	mutex.Lock()
	defer mutex.Unlock()

	user, exists := server.Users[username]
	if !exists {
		log.Printf("fetchFeed: User '%s' not found", username)
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var posts []*core.Post
	for subredditName := range user.JoinedSubreddits {
		subreddit, exists := server.Subreddits[subredditName]
		if !exists {
			continue
		}
		posts = append(posts, subreddit.Posts...)
	}
	log.Printf("fetchFeed: Feed fetched for user '%s'", username)
	c.JSON(200, posts)
}
