package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
)

type (
	ILogoutService interface {
		Logout(ctx context.Context, token string) error
	}

	ILogoutHandler interface {
		LogoutHandler(c *gin.Context)
	}
)