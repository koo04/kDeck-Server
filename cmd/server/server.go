package main

import (
	"fmt"
	"log"
	"net"

	"github.com/koo04/kdeck-server/api"
	"github.com/koo04/kdeck-server/proto/data"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	data.RegisterDataServiceServer(grpcServer, &api.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
