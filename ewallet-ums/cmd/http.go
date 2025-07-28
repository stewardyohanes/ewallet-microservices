package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/handlers"
	"ewallet-ums/internal/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)


func ServeHTTP() {
	healthCheckSvc := &services.HealthCheck{}
	healthCheckHandler := &handlers.HealthCheck{
		HealthCheckService: healthCheckSvc,
	}

	r := gin.Default()
	r.GET("/health", healthCheckHandler.HealthCheckHandlerHTTP)

	err := r.Run(fmt.Sprintf(":%s", helpers.GetEnv("HTTP_PORT", "8080")))
	if err != nil {
		log.Fatal("Error to start HTTP server", err)
	}
}
