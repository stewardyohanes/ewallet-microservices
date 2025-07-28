package interfaces

import "ewallet-ums/internal/models"

type IUserRepository interface {
	InsertNewUser(user *models.Users) error
}