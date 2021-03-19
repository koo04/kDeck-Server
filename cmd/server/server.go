package main

import (
	"fmt"
	"log"
	"net"

	"github.com/koo04/kdeck-server/api"
	"github.com/koo04/kdeck-server/config"
	"github.com/koo04/kdeck-server/proto/data"
	"google.golang.org/grpc"
)

func main() {
	var config = config.LoadConfig()

	// if config.Obs.Enabled {
	// 	plugins.
	// }

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	data.RegisterDataServiceServer(grpcServer, &api.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
