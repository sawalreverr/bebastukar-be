package dto

import (
	"mime/multipart"
	"time"
)

type DiscussionCredential struct {
	AuthorID string   `json:"author_id"`
	Content  string   `json:"content"`
	Images   []string `json:"images"`
}

type DiscussionCommentCredential struct {
	AuthorID string `json:"author_id"`
	Content  string `json:"content"`
}

type DiscussionResponse struct {
	ID        string    `json:"id"`
	AuthorID  string    `json:"author_id"`
	Content   string    `json:"content"`
	Images    []string  `json:"images"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type DiscussionCommentResponse struct {
	ID        string    `json:"id"`
	AuthorID  string    `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type DiscussionPaginationResponse struct {
	TotalCount  int                  `json:"total_count"`
	Page        int                  `json:"page"`
	Limit       int                  `json:"limit"`
	Discussions []DiscussionResponse `json:"discussions"`
}

type DiscussionInput struct {
	Content string                  `form:"content" validate:"required,min=1,max=500"`
	Images  []*multipart.FileHeader `form:"images"`
}
