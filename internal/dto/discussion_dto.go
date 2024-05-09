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
	AuthorID     string `json:"author_id"`
	DiscussionID string `json:"discussion_id"`
	Comment      string `json:"comment"`
}

type DiscussionReplyCommentCredential struct {
	AuthorID            string `json:"author_id"`
	DiscussionID        string `json:"discussion_id"`
	DiscussionCommentID string `json:"discussion_comment_id"`
	ReplyComment        string `json:"reply_comment"`
}

type DiscussionResponse struct {
	ID        string                      `json:"id"`
	AuthorID  string                      `json:"author_id"`
	Content   string                      `json:"content"`
	Images    []string                    `json:"images"`
	Comment   []DiscussionCommentResponse `json:"comments"`
	CreatedAt time.Time                   `json:"created_at,omitempty"`
	UpdatedAt time.Time                   `json:"updated_at,omitempty"`
}

type DiscussionCommentResponse struct {
	ID           string                           `json:"id"`
	AuthorID     string                           `json:"author_id"`
	DiscussionID string                           `json:"discussion_id"`
	Comment      string                           `json:"comment"`
	ReplyComment []DiscussionReplyCommentResponse `json:"reply_comments,omitempty"`
	CreatedAt    time.Time                        `json:"created_at,omitempty"`
	UpdatedAt    time.Time                        `json:"updated_at,omitempty"`
}

type DiscussionReplyCommentResponse struct {
	ID                  string    `json:"id"`
	AuthorID            string    `json:"author_id"`
	DiscussionID        string    `json:"discussion_id"`
	DiscussionCommentID string    `json:"discussion_comment_id"`
	ReplyComment        string    `json:"comment"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
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

type DiscussionCommentInput struct {
	Comment string `form:"comment" validate:"required,min=1,max=500"`
}
