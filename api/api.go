package api

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/koo04/kdeck-server/config"
	"github.com/koo04/kdeck-server/proto/data"
	"golang.org/x/net/context"
)

type Button struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	ImageURL string `json:"img"`
	Plugin   string `json:"plugin"`
	Action   string `json:"action"`
}

type Server struct {
	Buttons []Button
}

func api() {
	log.Println(*config.GetConfig())
}

func (s *Server) GetButtons(ctx context.Context, _ *data.Empty) (*data.GetButtonsResponse, error) {
	s.Buttons = []Button{
		{Name: "Desktop Scene", Type: "text", Plugin: "obs", Action: "scene:change:Desktop"},
		{Name: "Scene 2", Type: "text", Plugin: "obs", Action: "scene:change:Scene 2"},
		{Name: "Mute Desktop Audio", Type: "image", ImageURL: "https://www.svgrepo.com/show/107080/mute.svg", Plugin: "obs", Action: "mic:mute:Desktop Audio"},
	}

	j, err := json.Marshal(s.Buttons)
	if err != nil {
		log.Fatalf("Error marshaling data: %s", err)
	}
	log.Printf("Sending response %s\n", string(j))
	return &data.GetButtonsResponse{Body: string(j)}, nil
}

func (s *Server) PressButton(ctx context.Context, request *data.PressButtonRequest) (*data.Empty, error) {
	if request.Plugin == "obs" {
		var action = strings.Split(request.Action, ":")

		if action[0] == "scene" {
			if action[1] == "change" {
				s.obs.ChangeScene(action[2])
			}
		}

		if action[0] == "mic" {
			if action[1] == "mute" {
				s.obs.ToggleMicMute(action[2])
			}
		}
	}

	return &data.Empty{}, nil
}
