package repositories

import (
	"ewallet-ums/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(user *models.Users) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) InsertNewUserSession(userSession *models.UserSessions) error {
	return r.DB.Create(userSession).Error
}

func (r *UserRepository) GetUserByUsername(username string) (*models.Users, error) {
	var user models.Users

	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}