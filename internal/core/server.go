package core

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	redditpb "Go_project/internal/generated"
)

// RedditServiceServer defines the structure for managing Reddit-like services
type RedditServiceServer struct {
	redditpb.UnimplementedRedditServiceServer
	Users      map[string]*User
	Subreddits map[string]*Subreddit
	Posts      map[string]*Post
	Comments   map[string]*Comment
	mutex      sync.Mutex
}

// NewServer initializes a new RedditServiceServer
func NewServer() *RedditServiceServer {
	return &RedditServiceServer{
		Users:      make(map[string]*User),
		Subreddits: make(map[string]*Subreddit),
		Posts:      make(map[string]*Post),
		Comments:   make(map[string]*Comment),
	}
}

// CreateSubreddit allows creating a new subreddit
func (s *RedditServiceServer) CreateSubreddit(ctx context.Context, req *redditpb.CreateSubredditRequest) (*redditpb.CreateSubredditResponse, error) {
	log.Printf("[CreateSubreddit] Request received: %v", req)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.Subreddits[req.Name]; exists {
		log.Printf("[CreateSubreddit] Subreddit '%s' already exists", req.Name)
		return nil, fmt.Errorf("subreddit %s already exists", req.Name)
	}

	subreddit := &Subreddit{
		Name:        req.Name,
		Description: req.Description,
		Members:     make(map[string]bool),
		Posts:       []*Post{},
	}
	s.Subreddits[req.Name] = subreddit

	log.Printf("[CreateSubreddit] Subreddit '%s' created successfully", req.Name)
	return &redditpb.CreateSubredditResponse{
		Message: fmt.Sprintf("Subreddit '%s' created successfully", req.Name),
	}, nil
}

// RegisterUser handles user registration
func (s *RedditServiceServer) RegisterUser(ctx context.Context, req *redditpb.RegisterRequest) (*redditpb.RegisterResponse, error) {
	log.Printf("[RegisterUser] Registering user: %s", req.Username)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.Users[req.Username]; exists {
		log.Printf("[RegisterUser] User '%s' already exists", req.Username)
		return nil, fmt.Errorf("user %s already exists", req.Username)
	}

	s.Users[req.Username] = &User{
		Username:         req.Username,
		JoinedSubreddits: make(map[string]bool),
	}

	log.Printf("[RegisterUser] User '%s' registered successfully", req.Username)
	return &redditpb.RegisterResponse{
		Message: fmt.Sprintf("User '%s' registered successfully", req.Username),
	}, nil
}

// JoinSubreddit allows a user to join a subreddit
func (s *RedditServiceServer) JoinSubreddit(ctx context.Context, req *redditpb.JoinSubredditRequest) (*redditpb.JoinSubredditResponse, error) {
	log.Printf("[JoinSubreddit] User '%s' joining subreddit '%s'", req.Username, req.Subreddit)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	subreddit, exists := s.Subreddits[req.Subreddit]
	if !exists {
		log.Printf("[JoinSubreddit] Subreddit '%s' not found", req.Subreddit)
		return nil, fmt.Errorf("subreddit %s not found", req.Subreddit)
	}

	user, exists := s.Users[req.Username]
	if !exists {
		log.Printf("[JoinSubreddit] User '%s' not found", req.Username)
		return nil, fmt.Errorf("user %s not found", req.Username)
	}

	user.JoinedSubreddits[req.Subreddit] = true
	subreddit.Members[req.Username] = true

	log.Printf("[JoinSubreddit] User '%s' joined subreddit '%s'", req.Username, req.Subreddit)
	return &redditpb.JoinSubredditResponse{
		Message: fmt.Sprintf("User '%s' joined subreddit '%s'", req.Username, req.Subreddit),
	}, nil
}

// CreatePost allows a user to create a post in a subreddit
func (s *RedditServiceServer) CreatePost(ctx context.Context, req *redditpb.CreatePostRequest) (*redditpb.CreatePostResponse, error) {
	log.Printf("[CreatePost] User '%s' creating post in subreddit '%s'", req.Username, req.Subreddit)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	subreddit, exists := s.Subreddits[req.Subreddit]
	if !exists {
		log.Printf("[CreatePost] Subreddit '%s' not found", req.Subreddit)
		return nil, fmt.Errorf("subreddit %s not found", req.Subreddit)
	}

	postID := fmt.Sprintf("%s-%d", req.Subreddit, len(subreddit.Posts)+1)
	post := &Post{
		ID:        postID,
		Author:    req.Username,
		Title:     req.Title,
		Content:   req.Content,
		Subreddit: req.Subreddit,
		CreatedAt: time.Now(),
	}

	subreddit.Posts = append(subreddit.Posts, post)
	s.Posts[postID] = post

	log.Printf("[CreatePost] Post '%s' created successfully", postID)
	return &redditpb.CreatePostResponse{PostId: postID}, nil
}

// AddComment allows a user to add a comment to a post
func (s *RedditServiceServer) AddComment(ctx context.Context, req *redditpb.CommentRequest) (*redditpb.CommentResponse, error) {
	log.Printf("[AddComment] User '%s' adding comment to post '%s'", req.Username, req.PostId)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	post, exists := s.Posts[req.PostId]
	if !exists {
		log.Printf("[AddComment] Post '%s' not found", req.PostId)
		return nil, fmt.Errorf("post %s not found", req.PostId)
	}

	commentID := fmt.Sprintf("%s-comment-%d", req.PostId, len(post.Comments)+1)
	comment := &Comment{
		ID:        commentID,
		Author:    req.Username,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}

	post.Comments = append(post.Comments, comment)
	s.Comments[commentID] = comment

	log.Printf("[AddComment] Comment '%s' added successfully", commentID)
	return &redditpb.CommentResponse{CommentId: commentID}, nil
}

// FetchFeed retrieves posts for a user
func (s *RedditServiceServer) FetchFeed(ctx context.Context, req *redditpb.FetchFeedRequest) (*redditpb.FetchFeedResponse, error) {
	log.Printf("[FetchFeed] Fetching feed for user '%s'", req.Username)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	user, exists := s.Users[req.Username]
	if !exists {
		log.Printf("[FetchFeed] User '%s' not found", req.Username)
		return nil, fmt.Errorf("user %s not found", req.Username)
	}

	var posts []*redditpb.Post
	for subredditName := range user.JoinedSubreddits {
		subreddit, exists := s.Subreddits[subredditName]
		if !exists {
			continue
		}

		for _, post := range subreddit.Posts {
			posts = append(posts, &redditpb.Post{
				Id:        post.ID,
				Author:    post.Author,
				Title:     post.Title,
				Content:   post.Content,
				Subreddit: post.Subreddit,
				CreatedAt: post.CreatedAt.Format(time.RFC3339),
			})
		}
	}

	log.Printf("[FetchFeed] Feed fetched successfully for user '%s'", req.Username)
	return &redditpb.FetchFeedResponse{Posts: posts}, nil
}

// Repost allows a user to repost content to a different subreddit
func (s *RedditServiceServer) Repost(ctx context.Context, req *redditpb.RepostRequest) (*redditpb.RepostResponse, error) {
	log.Printf("[Repost] User '%s' reposting post '%s' to subreddit '%s'", req.Username, req.PostId, req.TargetSubreddit)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if the user exists
	if _, userExists := s.Users[req.Username]; !userExists {
		log.Printf("[Repost] User '%s' not found", req.Username)
		return nil, fmt.Errorf("user %s not found", req.Username)
	}

	// Check if the target subreddit exists
	targetSubreddit, subredditExists := s.Subreddits[req.TargetSubreddit]
	if !subredditExists {
		log.Printf("[Repost] Subreddit '%s' not found", req.TargetSubreddit)
		return nil, fmt.Errorf("subreddit %s not found", req.TargetSubreddit)
	}

	// Check if the original post exists
	originalPost, postExists := s.Posts[req.PostId]
	if !postExists {
		log.Printf("[Repost] Original post '%s' not found", req.PostId)
		return nil, fmt.Errorf("post %s not found", req.PostId)
	}

	// Create the repost
	repostID := fmt.Sprintf("%s-repost-%d", req.TargetSubreddit, len(targetSubreddit.Posts)+1)
	repost := &Post{
		ID:        repostID,
		Author:    req.Username,
		Title:     originalPost.Title,
		Content:   originalPost.Content,
		Subreddit: req.TargetSubreddit,
		CreatedAt: time.Now(),
	}

	// Add the repost to the target subreddit and the global post map
	targetSubreddit.Posts = append(targetSubreddit.Posts, repost)
	s.Posts[repostID] = repost

	log.Printf("[Repost] Post '%s' reposted successfully to subreddit '%s'", req.PostId, req.TargetSubreddit)
	return &redditpb.RepostResponse{
		Message: fmt.Sprintf("Post '%s' reposted successfully", req.PostId),
	}, nil
}
