package usecase

import (
	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/entity"
)

type UserUsecase interface {
	RegisterUser(user dto.UserCredential) (*entity.User, error)
	LoginUser(email string, password string) (string, error)
	UpdateUser(userID string, user dto.UserCredential) error
	DeleteUser(userID string) error
}
