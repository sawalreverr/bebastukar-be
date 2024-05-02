package dto

type UserCredential struct {
	Name        string `json:"name" validate:"required,min=2"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10"`
	Password    string `json:"password" validate:"required,min=8"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
