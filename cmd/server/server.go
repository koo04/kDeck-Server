package main

import (
	"fmt"
	"log"
	"net"

	"github.com/koo04/kdeck-server/api"
	"github.com/koo04/kdeck-server/plugins"
	"github.com/koo04/kdeck-server/proto/data"
	"github.com/koo04/kdeck-server/settings"
	"google.golang.org/grpc"
)

func main() {
	_ = settings.Load()
	plugins.Load()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	data.RegisterDataServiceServer(grpcServer, &api.API{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
