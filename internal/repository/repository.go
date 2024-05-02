package repository

import "github.com/sawalreverr/bebastukar-be/internal/entity"

type UserRepository interface {
	Create(user entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByID(userID string) (*entity.User, error)
	FindAll() (*[]entity.User, error)
	Update(userID string, user entity.User) error
	Delete(userID string) error
}
