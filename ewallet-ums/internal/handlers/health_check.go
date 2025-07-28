package handlers

import (
	"ewallet-microservices/helpers"
	"ewallet-microservices/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	HealthCheckService interfaces.IHealthCheckService
}

func (h *HealthCheck) HealthCheckHandlerHTTP(c *gin.Context) {
	msg, err := h.HealthCheckService.HealthCheckService()
	if err != nil {
		helpers.SendResposneHTTP(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.SendResposneHTTP(c, http.StatusOK, msg, nil)
}
