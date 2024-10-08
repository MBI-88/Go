package github

import "time"

// IssuesURL for Github
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult model
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

// Issue model
type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string 
	User *User 
	CreatedAt time.Time `json:"created_at"`
	Body string // in Markdown format
}

// User model
type User struct {
	Login string 
	HTMLURL string `json:"html_url"`
}