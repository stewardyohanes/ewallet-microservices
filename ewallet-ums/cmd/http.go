package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/handlers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/repositories"
	"ewallet-ums/internal/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)


func ServeHTTP() {
	dependency := dependencyInject()

	r := gin.Default()
	r.GET("/health", dependency.HealthCheckAPI.HealthCheckHandlerHTTP)

	userV1 := r.Group("/api/v1/users")
	userV1.POST("/register", dependency.RegisterAPI.RegisterHandler)
	userV1.POST("/login", dependency.LoginAPI.LoginHandler)
	userV1.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutAPI.LogoutHandler)

	err := r.Run(fmt.Sprintf(":%s", helpers.GetEnv("HTTP_PORT", "8080")))
	if err != nil {
		log.Fatal("Error to start HTTP server", err)
	}
}

type Dependency struct {
	UserRepository interfaces.IUserRepository

	HealthCheckAPI *handlers.HealthCheck
	RegisterAPI 	interfaces.IRegisterHandler
	LoginAPI 		interfaces.ILoginHandler
	LogoutAPI		interfaces.ILogoutHandler
}

func dependencyInject() Dependency {
	healthCheckSvc := &services.HealthCheck{}
	healthCheckHandler := &handlers.HealthCheck{HealthCheckService: healthCheckSvc}

	userRepo := &repositories.UserRepository{DB: helpers.DB}
	registerSvc := &services.RegisterService{UserRepo: userRepo}
	registerHandler := &handlers.RegisterHandler{RegisterService: registerSvc}
	loginSvc := &services.LoginService{UserRepo: userRepo}
	loginHandler := &handlers.LoginHandler{LoginService: loginSvc}
	logoutSvc := &services.LogoutService{UserRepo: userRepo}
	logoutHandler := &handlers.LogoutHandler{LogoutService: logoutSvc}


	return Dependency{
		UserRepository: userRepo,
		HealthCheckAPI: healthCheckHandler,
		RegisterAPI: registerHandler,
		LoginAPI: loginHandler,
		LogoutAPI: logoutHandler,
	}
}
