package core

import "time"

// User represents a user in the Reddit-like platform
type User struct {
	Username         string            `json:"username"`
	JoinedSubreddits map[string]bool   `json:"joined_subreddits"`
	Messages         []DirectMessage   `json:"messages"`
}

// Subreddit represents a subreddit on the platform
type Subreddit struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Members     map[string]bool `json:"members"`
	Posts       []*Post         `json:"posts"`
}

// Post represents a post in a subreddit
type Post struct {
	ID        string     `json:"id"`
	Author    string     `json:"author"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Subreddit string     `json:"subreddit"`
	CreatedAt time.Time  `json:"created_at"`
	Comments  []*Comment `json:"comments"`
}

// Comment represents a comment on a post
type Comment struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// DirectMessage represents a private message between users
type DirectMessage struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
