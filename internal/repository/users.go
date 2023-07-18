package repository

import (
	"github.com/gavrishp/coreapiservicetest/internal/common/models"
	"gorm.io/gorm"
)

type UsersRepository struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{DB: db}
}

func (r *UsersRepository) AddUser(user *models.User) (*models.User, error) {

	if result := r.DB.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UsersRepository) GetUser(userId string) (*models.User, error) {
	var user models.User

	if result := r.DB.First(&user, userId); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UsersRepository) GetUsers() (*[]models.User, error) {
	var users []models.User

	if result := r.DB.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (r *UsersRepository) DeleteUser(userId string) error {
	var user models.User

	if result := r.DB.First(&user, userId); result.Error != nil {
		return result.Error
	}

	r.DB.Delete(&user)
	return nil
}

func (r *UsersRepository) UpdateUser(userId string, user *models.User) (*models.User, error) {
	var currentUser models.User

	if result := r.DB.First(&currentUser, userId); result.Error != nil {
		return nil, result.Error
	}

	currentUser.Name = user.Name
	currentUser.Role = user.Role
	currentUser.Description = user.Description

	r.DB.Save(&currentUser)

	return &currentUser, nil
}
