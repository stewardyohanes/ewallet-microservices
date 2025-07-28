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