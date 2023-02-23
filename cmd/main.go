package main

import (
	"fmt"
	"go-employees/pkg/config"
	"go-employees/pkg/pb"
	"go-employees/pkg/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	empployeeServer := services.InitEmployeeService()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.Host, conf.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(grpcServer, empployeeServer)
	log.Printf("Launching grpc server on: %s:%s", conf.Host, conf.Port)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v\n", err)
	}
}
