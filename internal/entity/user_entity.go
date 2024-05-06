package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	ImageURL    string `json:"image_url"`
	Bio         string `json:"bio"`
	Password    string `json:"password"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
