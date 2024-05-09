package dto

import "mime/multipart"

type UpdateUser struct {
	Name        string `json:"name" validate:"required,min=2"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10"`
	Bio         string `json:"bio"`
}

type ProfileUser struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	ImageURL    string `json:"image_url"`
	Bio         string `json:"bio"`
}

type AvatarUpload struct {
	Image *multipart.File `form:"image" validate:"required,image"`
}
