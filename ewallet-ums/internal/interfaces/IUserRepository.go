package interfaces

import "ewallet-ums/internal/models"

type IUserRepository interface {
	InsertNewUser(user *models.Users) error
	InsertNewUserSession(userSession *models.UserSessions) error
	GetUserByUsername(username string) (*models.Users, error)
}