package handlers

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	UserService interfaces.IRegisterService
}

func (h *RegisterHandler) RegisterHandler(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	req := &models.Users{}

	if err := c.ShouldBindJSON(req); err != nil {
		log.Error("Error binding JSON", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Error validating request", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	resp, err := h.UserService.Register(c.Request.Context(), req)
	if err != nil {
		log.Error("Error registering user", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, "User created successfully", resp)
}

