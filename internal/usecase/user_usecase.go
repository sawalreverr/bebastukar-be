package usecase

import (
	"github.com/google/uuid"
	"github.com/sawalreverr/bebastukar-be/config"
	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/entity"
	"github.com/sawalreverr/bebastukar-be/internal/helper"
	"github.com/sawalreverr/bebastukar-be/internal/repository"
	"github.com/sawalreverr/bebastukar-be/pkg"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (u *userUsecase) RegisterUser(userCred dto.UserCredential) (*entity.User, error) {
	userFound, _ := u.userRepository.FindByEmail(userCred.Email)

	if userFound != nil {
		return nil, pkg.ErrDataAlreadyExist
	}

	hashedPass, _ := helper.GenerateHash(userCred.Password)

	user := entity.User{
		ID:          uuid.NewString(),
		Name:        userCred.Name,
		Email:       userCred.Email,
		PhoneNumber: userCred.PhoneNumber,
		ImageURL:    "",
		Role:        "user",
		Password:    hashedPass,
	}

	newUser, err := u.userRepository.Create(user)
	return newUser, err
}

func (u *userUsecase) LoginUser(email string, password string) (string, error) {
	userFound, err := u.userRepository.FindByEmail(email)

	if err != nil {
		return "", pkg.ErrRecordNotFound
	}

	ok := helper.ComparePassword(userFound.Password, password)

	if !ok {
		return "", pkg.ErrRecordNotFound
	}

	secretKey := config.GetConfig().Server.JWTSecret
	token, err := helper.GenerateTokenJWT(userFound.ID, userFound.Role, secretKey)

	return token, err
}

func (u *userUsecase) UpdateUser(userID string, user dto.UserCredential) error {
	var err error

	_, err = u.userRepository.FindByID(userID)

	if err != nil {
		return err
	}

	userCred := entity.User{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}

	err = u.userRepository.Update(userID, userCred)

	return err
}

func (u *userUsecase) DeleteUser(userID string) error {
	err := u.userRepository.Delete(userID)

	return err
}
