package client

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"

	redditpb "Go_project/internal/generated"
	"google.golang.org/grpc"
)

type Simulator struct {
	client         redditpb.RedditServiceClient
	userCount      int
	subredditCount int
	mutex          sync.Mutex
	activeUsers    map[string]bool
	stats          SimulationStats
	subredditNames []string // Track created subreddits
	postIDs        []string // Track created posts
	outputMutex    sync.Mutex
}

type SimulationStats struct {
	UsersCreated     int
	SubredditsCreated int
	PostsCreated     int
	RepostsCreated   int
	CommentsCreated  int
	VotesCast        int
	FeedsFetched     int
}

func NewSimulator(serverAddr string, userCount, subredditCount int) *Simulator {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	client := redditpb.NewRedditServiceClient(conn)
	return &Simulator{
		client:         client,
		userCount:      userCount,
		subredditCount: subredditCount,
		activeUsers:    make(map[string]bool),
		stats:          SimulationStats{},
		subredditNames: []string{},
		postIDs:        []string{},
	}
}

func (s *Simulator) Run() {
	start := time.Now()
	s.createSubreddits()
	s.createUsers()
	go s.simulateActivity()
	time.Sleep(2 * time.Minute)
	duration := time.Since(start)
	s.safePrintf("Simulation completed in %v\n", duration)
	s.printStats()
}

func (s *Simulator) safePrintf(format string, args ...interface{}) {
	s.outputMutex.Lock()
	defer s.outputMutex.Unlock()
	fmt.Printf(format, args...)
}

func (s *Simulator) createSubreddits() {
	for i := 0; i < s.subredditCount; i++ {
		name := fmt.Sprintf("subreddit%d", i)
		_, err := s.client.CreateSubreddit(context.Background(), &redditpb.CreateSubredditRequest{
			Name:        name,
			Description: fmt.Sprintf("Description for %s", name),
		})
		if err == nil {
			s.stats.SubredditsCreated++
			s.mutex.Lock()
			s.subredditNames = append(s.subredditNames, name)
			s.mutex.Unlock()
			s.safePrintf("Created subreddit: %s\n", name)
			for j := 0; j < 10; j++ { // Minimum 10 posts per subreddit
				s.createInitialPost(name)
			}
		} else {
			s.safePrintf("Failed to create subreddit %s: %v\n", name, err)
			log.Printf("Error creating subreddit %s: %v", name, err) // Additional logging
		}
	}
}

func (s *Simulator) createInitialPost(subredditName string) {
	username := fmt.Sprintf("user%d", rand.Intn(s.userCount)) // Random user for initial post
	title := fmt.Sprintf("Initial Post in %s", subredditName)
	content := fmt.Sprintf("Content of the initial post in %s", subredditName)
	resp, err := s.client.CreatePost(context.Background(), &redditpb.CreatePostRequest{
		Username:  username,
		Subreddit: subredditName,
		Title:     title,
		Content:   content,
	})
	if err == nil {
		postID := resp.GetPostId()
		s.mutex.Lock()
		s.stats.PostsCreated++
		s.postIDs = append(s.postIDs, postID) // Ensure we store post IDs
		s.mutex.Unlock()
		s.safePrintf("Created initial post in %s by %s with ID: %s\n", subredditName, username, postID)
	} else {
		s.safePrintf("Failed to create initial post in %s: %v\n", subredditName, err)
	}
}

func (s *Simulator) createUsers() {
	zipfWeights := generateZipfDistribution(len(s.subredditNames), 1.1)
	for i := 0; i < s.userCount; i++ {
		username := fmt.Sprintf("user%d", i)
		_, err := s.client.RegisterUser(context.Background(), &redditpb.RegisterRequest{
			Username: username,
		})
		if err == nil {
			s.stats.UsersCreated++
			s.safePrintf("Created user: %s\n", username)
			for j, weight := range zipfWeights {
				if rand.Intn(1000) < weight {
					_, err := s.client.JoinSubreddit(context.Background(), &redditpb.JoinSubredditRequest{
						Username:  username,
						Subreddit: s.subredditNames[j],
					})
					if err != nil {
						s.safePrintf("User %s failed to join subreddit %s: %v\n", username, s.subredditNames[j], err)
					}
				}
			}
		} else {
			s.safePrintf("Failed to create user %s: %v\n", username, err)
		}
	}
}

func (s *Simulator) simulateActivity() {
	for i := 0; i < s.userCount; i++ {
		go s.simulateUser(fmt.Sprintf("user%d", i))
	}
}

func (s *Simulator) simulateUser(username string) {
	for {
		s.simulateConnection(username)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

func (s *Simulator) simulateConnection(username string) {
	s.mutex.Lock()
	s.activeUsers[username] = true
	s.mutex.Unlock()
	defer func() {
		s.mutex.Lock()
		delete(s.activeUsers, username)
		s.mutex.Unlock()
	}()
	duration := time.Duration(rand.Intn(30)) * time.Second
	timer := time.NewTimer(duration)

	// Simulate disconnection
	disconnectChance := rand.Intn(100) < 10 // 10% chance to disconnect
	if disconnectChance {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Second) // Disconnect for a random time
		log.Printf("%s disconnected\n", username)
	}

	for {
		select {
		case <-timer.C:
			return
		default:
			s.performRandomAction(username)
			time.Sleep(time.Second)
		}
	}
}

func (s *Simulator) performRandomAction(username string) {
	action := rand.Intn(6) // Log the action being performed
	s.safePrintf("User %s performing action %d\n", username, action)
	switch action {
	case 0:
		s.safePrintf("%s is joining a random subreddit\n", username)
		s.joinRandomSubreddit(username)
	case 1:
		s.safePrintf("%s is creating a random post\n", username)
		s.createRandomPost(username)
	case 2:
		s.safePrintf("%s is creating a random comment\n", username)
		s.createRandomComment(username)
	case 3:
		s.safePrintf("%s is voting randomly\n", username)
		s.voteRandomly(username)
	case 4:
		s.safePrintf("%s is fetching their feed\n", username)
		s.getFeed(username)
	case 5:
		s.safePrintf("%s is reposting content\n", username)
		s.createRandomRepost(username)
	}
}

func (s *Simulator) joinRandomSubreddit(username string) {
	if len(s.subredditNames) == 0 {
		s.safePrintf("No subreddits available for %s to join\n", username)
		return
	}
	subredditName := s.subredditNames[rand.Intn(len(s.subredditNames))]
	if _, err := s.client.JoinSubreddit(context.Background(), &redditpb.JoinSubredditRequest{
		Username:  username,
		Subreddit: subredditName,
	}); err == nil {
		s.safePrintf("%s joined %s\n", username, subredditName)
	} else {
		log.Printf("%s failed to join subreddit %s: %v", username, subredditName, err)
	}
}

func (s *Simulator) createRandomPost(username string) {
	if len(s.subredditNames) == 0 {
		log.Printf("No subreddits available for %s to create a post\n", username)
		return
	}
	subredditName := s.subredditNames[rand.Intn(len(s.subredditNames))]
	title := "Random Post by " + username
	content := "Content of random post by " + username
	resp, err := s.client.CreatePost(context.Background(), &redditpb.CreatePostRequest{
		Username:  username,
		Subreddit: subredditName,
		Title:     title,
		Content:   content,
	})
	if err == nil {
		postID := resp.GetPostId()
		log.Printf("%s created post in %s with ID: %v\n", username, subredditName, postID)
		s.stats.PostsCreated++ // Increment Posts Created
	} else {
		log.Printf("Failed to create post for %s in %s: %v\n", username, subredditName, err)
	}
}

func (s *Simulator) createRandomComment(username string) {
	if len(s.postIDs) == 0 {
		log.Printf("No posts available for %s to comment on\n", username)
		return
	}
	postID := s.postIDs[rand.Intn(len(s.postIDs))]
	content := "Comment by " + username
	if _, err := s.client.AddComment(context.Background(), &redditpb.CommentRequest{
		Username: username,
		PostId:   postID,
		Content:  content,
	}); err == nil {
		log.Printf("%s commented on post %v\n", username, postID)
		s.stats.CommentsCreated++ // Increment Comments Created
	} else {
		log.Printf("Failed to comment on post %v by %s: %v", postID, username, err)
	}
}

func (s *Simulator) voteRandomly(username string) {
	if len(s.postIDs) == 0 {
		log.Printf("No items available for %s to vote on\n", username)
		return
	}
	itemID := s.postIDs[rand.Intn(len(s.postIDs))]
	upvote := rand.Intn(2) == 0
	if upvote {
		_, _ = s.client.UpvotePost(context.Background(), &redditpb.UpvoteRequest{
			Username: username,
			PostId:   itemID,
		})
		s.stats.VotesCast++ // Increment Votes Cast
	} else {
		_, _ = s.client.DownvotePost(context.Background(), &redditpb.DownvoteRequest{
			Username: username,
			PostId:   itemID,
		})
		s.stats.VotesCast++ // Increment Votes Cast
	}
	log.Printf("%s voted on item %v\n", username, itemID)
}

func (s *Simulator) getFeed(username string) {
	if _, err := s.client.FetchFeed(context.Background(), &redditpb.FetchFeedRequest{
		Username: username,
	}); err == nil {
		log.Printf("%s fetched their feed\n", username)
		s.stats.FeedsFetched++ // Increment Feeds Fetched
	} else {
		log.Printf("Failed to fetch feed for %s: %v", username, err)
	}
}

func (s *Simulator) createRandomRepost(username string) {
	if len(s.postIDs) == 0 || len(s.subredditNames) == 0 {
		log.Printf("No posts or subreddits available for %s to repost\n", username)
		return
	}
	postID := s.postIDs[rand.Intn(len(s.postIDs))]
	targetSubreddit := s.subredditNames[rand.Intn(len(s.subredditNames))]
	_, err := s.client.Repost(context.Background(), &redditpb.RepostRequest{
		Username:       username,
		PostId:         postID,
		TargetSubreddit: targetSubreddit,
	})
	if err == nil {
		log.Printf("%s reposted content to %v\n", username, targetSubreddit)
		s.stats.RepostsCreated++
	} else {
		log.Printf("Failed to repost content by %v:%v \n ", username, err)
	}
}

func (s *Simulator) printStats() {
	s.safePrintf("Simulation Statistics:\n")
	s.safePrintf("Users Created: %d\n", s.stats.UsersCreated)
	s.safePrintf("Subreddits Created: %d\n", s.stats.SubredditsCreated)
	s.safePrintf("Posts Created: %d\n", s.stats.PostsCreated)
	s.safePrintf("Reposts Created: %d\n", s.stats.RepostsCreated)
	s.safePrintf("Comments Created: %d\n", s.stats.CommentsCreated)
	s.safePrintf("Votes Cast: %d\n", s.stats.VotesCast)
	s.safePrintf("Feeds Fetched: %d\n", s.stats.FeedsFetched)
}

func generateZipfDistribution(n int, s float64) []int {
	weights := make([]int, n)

	var normalizationFactor float64
	for k := 1; k <= n; k++ {
		normalizationFactor += 1 / math.Pow(float64(k), s)
	}

	for k := 1; k <= n; k++ {
		weights[k-1] = int((1 / math.Pow(float64(k), s)) / normalizationFactor * 1000)
	}

	return weights
}
