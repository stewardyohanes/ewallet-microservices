package main

import (
	"ewallet-microservices/cmd"
	"ewallet-microservices/helpers"
)

func main() {
	// Load Config
	helpers.SetupConfig()

	// Load Logger
	helpers.SetupLogger()

	// Load Database
	helpers.SetupMySQL()

	// Run HTTP Server
	cmd.ServeHTTP()

	// Run GRPC Server
	go cmd.ServeGRPC()
}
