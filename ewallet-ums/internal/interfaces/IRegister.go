package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type (
	IRegisterService interface {
		Register(ctx context.Context, req *models.Users) (*models.Users, error)
	}

	IRegisterHandler interface {
		RegisterHandler(c *gin.Context)
	}
)