package repository

import "github.com/sawalreverr/bebastukar-be/internal/entity"

type UserRepository interface {
	Create(user entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByID(userID string) (*entity.User, error)
	FindAll(page int, limit int, sortBy string, sortType string) (*[]entity.User, error)
	Update(user *entity.User) error
	Delete(userID string) error
}
