package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	pb "Go_project/internal/generated"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRedditServiceClient(conn)

	// Create "golang" subreddit before running clients
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := createSubreddit(client, ctx, "golang", "A subreddit for Go developers"); err != nil {
		log.Fatalf("Failed to create subreddit: %v", err)
	}

	// Simulate multiple clients
	var wg sync.WaitGroup
	clientCount := 5 // Number of clients
	for i := 1; i <= clientCount; i++ {
		wg.Add(1)
		go simulateClient(client, fmt.Sprintf("testuser%d", i), &wg)
	}

	wg.Wait()
	log.Println("All clients finished execution")
}

// Simulate a single gRPC client
func simulateClient(client pb.RedditServiceClient, username string, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := registerUser(client, ctx, username); err != nil {
		log.Printf("Simulation error for user %s: %v", username, err)
		return
	}

	if err := joinSubreddit(client, ctx, username, "golang"); err != nil {
		log.Printf("Simulation error for user %s: %v", username, err)
		return
	}

	if err := createPost(client, ctx, username, "golang", "Post by "+username, "Content from "+username); err != nil {
		log.Printf("Simulation error for user %s: %v", username, err)
		return
	}

	if err := fetchFeed(client, ctx, username); err != nil {
		log.Printf("Simulation error for user %s: %v", username, err)
		return
	}
}

// Function implementations

func registerUser(client pb.RedditServiceClient, ctx context.Context, username string) error {
	res, err := client.RegisterUser(ctx, &pb.RegisterRequest{Username: username})
	if err != nil {
		return fmt.Errorf("RegisterUser failed: %w", err)
	}
	fmt.Printf("RegisterUser response for '%s': %s\n", username, res.GetMessage())
	return nil
}

func createSubreddit(client pb.RedditServiceClient, ctx context.Context, name, description string) error {
	res, err := client.CreateSubreddit(ctx, &pb.CreateSubredditRequest{Name: name, Description: description})
	if err != nil {
		return fmt.Errorf("CreateSubreddit failed: %w", err)
	}
	fmt.Printf("CreateSubreddit response for '%s': %s\n", name, res.GetMessage())
	return nil
}

func joinSubreddit(client pb.RedditServiceClient, ctx context.Context, username, subreddit string) error {
	res, err := client.JoinSubreddit(ctx, &pb.JoinSubredditRequest{Username: username, Subreddit: subreddit})
	if err != nil {
		return fmt.Errorf("JoinSubreddit failed: %w", err)
	}
	fmt.Printf("JoinSubreddit response for user '%s' and subreddit '%s': %s\n", username, subreddit, res.GetMessage())
	return nil
}

func createPost(client pb.RedditServiceClient, ctx context.Context, username, subreddit, title, content string) error {
	res, err := client.CreatePost(ctx, &pb.CreatePostRequest{
		Username:  username,
		Subreddit: subreddit,
		Title:     title,
		Content:   content,
	})
	if err != nil {
		return fmt.Errorf("CreatePost failed: %w", err)
	}
	fmt.Printf("CreatePost response for user '%s': Post ID: %s\n", username, res.GetPostId())
	return nil
}

func fetchFeed(client pb.RedditServiceClient, ctx context.Context, username string) error {
	res, err := client.FetchFeed(ctx, &pb.FetchFeedRequest{Username: username})
	if err != nil {
		return fmt.Errorf("FetchFeed failed: %w", err)
	}

	fmt.Printf("FetchFeed response for user '%s':\n", username)
	for _, post := range res.GetPosts() {
		fmt.Printf(" - Post ID: %s\n   Author: %s\n   Title: %s\n   Content: %s\n   Subreddit: %s\n   Created At: %s\n",
			post.GetId(), post.GetAuthor(), post.GetTitle(), post.GetContent(), post.GetSubreddit(), post.GetCreatedAt())
	}
	return nil
}
