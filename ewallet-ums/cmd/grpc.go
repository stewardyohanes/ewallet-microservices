package cmd

import (
	"ewallet-ums/helpers"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	listener, err := net.Listen("tcp", helpers.GetEnv("GRPC_PORT", ":7000"))
	if err != nil {
		log.Fatal("Error to listen GRPC server", err)
	}

	grpcServer := grpc.NewServer()

	logrus.Info("Listening GRPC server on port", helpers.GetEnv("GRPC_PORT", ":7000"))
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Error starting GRPC server", err)
	}
}
