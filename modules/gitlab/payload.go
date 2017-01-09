// Package gitlab is for return a payload parsed.
package gitlab

import (
	"time"
	"encoding/json"
)

// Payload struct
type Payload struct {
	ObjectKind        string      `json:"object_kind"`
	Before            string      `json:"before"`
	After             string      `json:"after"`
	Ref               string      `json:"ref"`
	CheckoutSha       string      `json:"checkout_sha"`
	UserID            int         `json:"user_id"`
	UserName          string      `json:"user_name"`
	UserEmail         string      `json:"user_email"`
	UserAvatar        string      `json:"user_avatar"`
	ProjectID         int         `json:"project_id"`
	Project           Project     `json:"project"`
	Repository        Repository  `json:"repository"`
	Commits           []Commit    `json:"Commits"`
	TotalCommitsCount int          `json:"total_commits_count"`
}

// Project struct
type Project struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	Avatarurl         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int    `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}

// Repository struct
type Repository struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitHTTPURL      string `json:"git_http_url"`
	GitSSHURL       string `json:"git_ssh_url"`
	VisibilityLevel int    `json:"visibility_level"`
}

// Commit struct
type Commit struct {
	ID        string        `json:"id"`
	Message   string        `json:"message"`
	Timestamp time.Time    `json:"timestamp"`
	URL       string        `json:"url"`
	Author    Author        `json:"author"`
	Added     []string      `json:"added"`
	Modified  []string      `json:"modified"`
	Removed   []string      `json:"removed"`
}

// Author struct with information about user
type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// GetName retries the name of Pusher.
func (p *Author) GetName() string {
	if len(p.Name) > 0 {
		return p.Name
	}
	if len(p.Username) > 0 {
		return p.Username
	}
	if len(p.Email) > 0 {
		return p.Email
	}
	return "Unknown"
}

// NewPayload returns a new struct of github slack-bot.
func NewPayload(b []byte) (*Payload, error) {
	p := &Payload{}

	if err := json.Unmarshal(b, p); err != nil {
		return p, ErrGitLabUnmarshal{Err: err}
	}

	return p, nil
}
