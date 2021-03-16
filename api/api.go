package api

import (
	"log"

	"github.com/koo04/kdeck-server/proto/data"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *data.SayHelloRequest) (*data.SayHelloResponse, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &data.SayHelloResponse{Body: "Hello From the Server!"}, nil
}
