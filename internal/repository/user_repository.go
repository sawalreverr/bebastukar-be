package repository

import (
	"github.com/sawalreverr/bebastukar-be/internal/database"
	"github.com/sawalreverr/bebastukar-be/internal/entity"
)

type userRepository struct {
	DB database.Database
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Create(user entity.User) (*entity.User, error) {
	if err := r.DB.GetDB().Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User
	if err := r.DB.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByID(userID string) (*entity.User, error) {
	var user *entity.User
	if err := r.DB.GetDB().Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindAll() (*[]entity.User, error) {
	var users *[]entity.User
	if err := r.DB.GetDB().Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Update(userID string, user entity.User) error {
	if err := r.DB.GetDB().Where("id = ?", userID).Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(userID string) error {
	if err := r.DB.GetDB().Where("id = ?", userID).Delete(&entity.User{}).Error; err != nil {
		return err
	}

	return nil
}
