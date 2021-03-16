package api

import (
	"encoding/json"
	"log"

	"github.com/koo04/kdeck-server/proto/data"
	"golang.org/x/net/context"
)

type Button struct {
	Name string `json:"name"`
}

type Server struct {
	Buttons []Button
}

func (s *Server) GetButtons(ctx context.Context, _ *data.Empty) (*data.GetButtonsResponse, error) {
	s.Buttons = []Button{{Name: "test"}, {Name: "test2"}}

	j, err := json.Marshal(s.Buttons)
	if err != nil {
		log.Fatalf("Error marshaling data: %s", err)
	}
	log.Printf("Sending response %s\n", string(j))
	return &data.GetButtonsResponse{Body: string(j)}, nil
}
