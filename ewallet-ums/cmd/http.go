package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/handlers"
	"ewallet-ums/internal/repositories"
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

	registerRepo := &repositories.UserRepository{
		DB: helpers.DB,
	}
	registerSvc := &services.RegisterService{
		UserRepo: registerRepo,
	}
	registerHandler := &handlers.RegisterHandler{
		UserService: registerSvc,
	}

	userV1 := r.Group("/api/v1/users")
	userV1.POST("/register", registerHandler.RegisterHandler)

	err := r.Run(fmt.Sprintf(":%s", helpers.GetEnv("HTTP_PORT", "8080")))
	if err != nil {
		log.Fatal("Error to start HTTP server", err)
	}
}
