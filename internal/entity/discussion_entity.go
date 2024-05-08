package entity

import (
	"time"

	"gorm.io/gorm"
)

type Discussions struct {
	ID      string `json:"id" gorm:"primaryKey"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type DiscussionImages struct {
	ID           string `json:"id" gorm:"primaryKey"`
	DiscussionID string `json:"discussion_id"`
	ImageURL     string `json:"image_url"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type DiscussionComments struct {
	ID           string `json:"id" gorm:"primaryKey"`
	DiscussionID string `json:"discussion_id"`
	UserID       string `json:"user_id"`
	Comment      string `json:"comment"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type DiscussionReplyComments struct {
	ID                  string `json:"id" gorm:"primaryKey"`
	DiscussionID        string `json:"discussion_id"`
	DiscussionCommentID string `json:"discussion_comment_id"`
	UserID              string `json:"user_id"`
	Comment             string `json:"comment"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
