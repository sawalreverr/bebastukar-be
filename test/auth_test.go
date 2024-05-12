package tests

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/entity"
	"github.com/sawalreverr/bebastukar-be/internal/helper"
	"github.com/sawalreverr/bebastukar-be/internal/usecase"
	"github.com/sawalreverr/bebastukar-be/internal/usecase/mocks"
	"github.com/sawalreverr/bebastukar-be/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	req := dto.UserCredential{
		Name:        "John Doe",
		Email:       "john.doe@gmail.com",
		PhoneNumber: "089511223344",
		Password:    "password@123",
	}

	user := entity.User{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Role:        "user",
		ImageURL:    "",
		Password:    "$2a$12$Tq.WcrxYy8u0fYJNLDjajukx.6cBtpXsPkNpx.ZglWbOo/dnzRfRy",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindByEmail", req.Email).Return(nil, nil)
		mockRepo.On("Create", mock.Anything).Return(&user, nil)

		newUser, err := uc.RegisterUser(req)
		assert.NoError(t, err)
		assert.NotNil(t, newUser)
		assert.Equal(t, newUser.Email, req.Email)
	})

	t.Run("Email already exists", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindByEmail", req.Email).Return(&user, nil)

		newUser, err := uc.RegisterUser(req)
		assert.Error(t, err)
		assert.Nil(t, newUser)
		assert.Equal(t, pkg.ErrDataAlreadyExist, err)
	})

	t.Run("Register failed", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindByEmail", req.Email).Return(nil, nil)
		mockRepo.On("Create", mock.Anything).Return(nil, errors.New("error on database"))

		newUser, err := uc.RegisterUser(req)
		assert.Error(t, err)
		assert.Nil(t, newUser)
		assert.Equal(t, "error on database", err.Error())
	})
}

func TestLoginUser(t *testing.T) {
	req := dto.Login{
		Email:    "john.doe@gmail.com",
		Password: "password@123",
	}

	hashPassword, _ := helper.GenerateHash(req.Password)

	user := entity.User{
		ID:          uuid.NewString(),
		Name:        "John Doe",
		Email:       req.Email,
		PhoneNumber: "089511223344",
		Role:        "user",
		ImageURL:    "",
		Password:    hashPassword,
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindByEmail", req.Email).Return(&user, nil)

		resp, err := uc.LoginUser(req.Email, req.Password)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, err, nil)
	})

	t.Run("Email or password invalid", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindByEmail", req.Email).Return(nil, errors.New("email or password invalid"))

		resp, err := uc.LoginUser(req.Email, req.Password)
		assert.Error(t, err)
		assert.Equal(t, "", resp)
		assert.Equal(t, pkg.ErrRecordNotFound, err)
	})
}
