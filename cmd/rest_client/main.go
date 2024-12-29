package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	baseURL        = "http://localhost:8080" // REST API base URL
	clientCount    = 100                    // Number of simulated clients
	requestTimeout = 5 * time.Second        // Timeout for HTTP requests
)

// Structs for request payloads and responses
type User struct {
	Username string `json:"username"`
}

type Subreddit struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Post struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type Comment struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

// Helper function to make HTTP requests
func makeRequest(method, endpoint string, payload interface{}) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	client := &http.Client{}
	var req *http.Request
	var err error

	if payload != nil {
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
		req, err = http.NewRequestWithContext(ctx, method, baseURL+endpoint, bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequestWithContext(ctx, method, baseURL+endpoint, nil)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Log the response status and body
	log.Printf("Response: StatusCode=%d, Endpoint=%s, Body=%s", resp.StatusCode, endpoint, string(body))

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("server responded with error %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// REST API operations
func registerUser(username string) {
	payload := User{Username: username}
	res, err := makeRequest("POST", "/users", payload)
	if err != nil {
		log.Printf("[RegisterUser] Failed for username '%s': %v", username, err)
		return
	}
	log.Printf("[RegisterUser] Successful for username '%s': %s", username, res)
}

func createSubreddit(name, description string) {
	payload := Subreddit{Name: name, Description: description}
	res, err := makeRequest("POST", "/subreddits", payload)
	if err != nil {
		log.Printf("[CreateSubreddit] Failed for subreddit '%s': %v", name, err)
		return
	}
	log.Printf("[CreateSubreddit] Successful for subreddit '%s': %s", name, res)
}

func joinSubreddit(username, subreddit string) {
	res, err := makeRequest("POST", fmt.Sprintf("/subreddits/%s/join?username=%s", subreddit, username), nil)
	if err != nil {
		log.Printf("[JoinSubreddit] Failed for username '%s' and subreddit '%s': %v", username, subreddit, err)
		return
	}
	log.Printf("[JoinSubreddit] Successful for username '%s' and subreddit '%s': %s", username, subreddit, res)
}

func createPost(username, subreddit, title, content string) {
	payload := Post{Username: username, Title: title, Content: content}
	res, err := makeRequest("POST", fmt.Sprintf("/subreddits/%s/posts", subreddit), payload)
	if err != nil {
		log.Printf("[CreatePost] Failed for username '%s' in subreddit '%s': %v", username, subreddit, err)
		return
	}
	log.Printf("[CreatePost] Successful for username '%s' in subreddit '%s': %s", username, subreddit, res)
}

func fetchFeed(username string) {
	res, err := makeRequest("GET", fmt.Sprintf("/users/%s/feed", username), nil)
	if err != nil {
		log.Printf("[FetchFeed] Failed for username '%s': %v", username, err)
		return
	}
	log.Printf("[FetchFeed] Successful for username '%s': %s", username, res)
}

// Simulate multiple REST API clients
func simulateRESTClient(username string, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("[SimulateRESTClient] Starting client for user: %s", username)
	registerUser(username)
	joinSubreddit(username, "golang")
	createPost(username, "golang", "Post by "+username, "Content from "+username)
	fetchFeed(username)
	log.Printf("[SimulateRESTClient] Completed client for user: %s", username)
}

func main() {
	// Create the "golang" subreddit before running clients
	createSubreddit("golang", "A subreddit for Go developers")

	// Simulate multiple clients
	var wg sync.WaitGroup
	for i := 1; i <= clientCount; i++ {
		wg.Add(1)
		go simulateRESTClient(fmt.Sprintf("testuser%d", i), &wg)
	}

	wg.Wait()
	log.Println("All clients finished execution")
}
