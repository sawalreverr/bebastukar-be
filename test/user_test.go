package tests

import (
	"errors"
	"testing"

	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/entity"
	"github.com/sawalreverr/bebastukar-be/internal/usecase"
	"github.com/sawalreverr/bebastukar-be/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUser(t *testing.T) {
	req := dto.UpdateUser{
		Name:        "John Doe",
		PhoneNumber: "089511223344",
		Bio:         "Updated bio",
	}

	userID := "user_id"
	user := &entity.User{
		ID:          userID,
		Name:        "Old Name",
		PhoneNumber: "1234567890",
		Bio:         "Old Bio",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindByID", userID).Return(user, nil)
		mockRepo.On("Update", mock.Anything).Return(nil)

		err := uc.UpdateUser(userID, req)
		assert.NoError(t, err)
	})

	t.Run("User not found", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindByID", userID).Return(nil, errors.New("user not found"))

		err := uc.UpdateUser(userID, req)
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestUpdateUserAvatar(t *testing.T) {
	userID := "user_id"
	imageURL := "http://example.com/avatar.jpg"

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindByID", userID).Return(&entity.User{}, nil)
		mockRepo.On("Update", mock.Anything).Return(nil)

		err := uc.UpdateUserAvatar(userID, imageURL)
		assert.NoError(t, err)
	})

	t.Run("User not found", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindByID", userID).Return(nil, errors.New("user not found"))

		err := uc.UpdateUserAvatar(userID, imageURL)
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestFindUserByID(t *testing.T) {
	userID := "user_id"
	user := &entity.User{
		ID: userID,
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindByID", userID).Return(user, nil)

		result, err := uc.FindUserByID(userID)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("User not found", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindByID", userID).Return(nil, errors.New("user not found"))

		result, err := uc.FindUserByID(userID)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestFindAllUser(t *testing.T) {
	page := 1
	limit := 10
	sortBy := "name"
	sortType := "asc"

	users := &[]entity.User{{}, {}}

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindAll", page, limit, sortBy, sortType).Return(users, nil)

		result, err := uc.FindAllUser(page, limit, sortBy, sortType)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("FindAll", page, limit, sortBy, sortType).Return(nil, errors.New("error"))

		result, err := uc.FindAllUser(page, limit, sortBy, sortType)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "error", err.Error())
	})
}

func TestDeleteUser(t *testing.T) {
	userID := "user_id"

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("Delete", userID).Return(nil)

		err := uc.DeleteUser(userID)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecase.NewUserUsecase(mockRepo)

		mockRepo.On("Delete", userID).Return(errors.New("error"))

		err := uc.DeleteUser(userID)
		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
	})
}
