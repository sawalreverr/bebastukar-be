package dto

import (
	"time"
)

type DiscussionCredential struct {
	AuthorID string   `json:"author_id"`
	Content  string   `json:"content"`
	Images   []string `json:"images"`
}

type DiscussionResponse struct {
	AuthorID  string    `json:"author_id"`
	Content   string    `json:"content"`
	Images    []string  `json:"images,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
