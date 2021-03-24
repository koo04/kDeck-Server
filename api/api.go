package api

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/koo04/kdeck-server/plugins/obs"
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

type API struct {
	buttons []Button
}

func (api *API) GetButtons(ctx context.Context, _ *data.Empty) (*data.GetButtonsResponse, error) {
	api.buttons = []Button{
		{Name: "Desktop Scene", Type: "text", Plugin: "obs", Action: "scene:change:Desktop"},
		{Name: "Scene 2", Type: "text", Plugin: "obs", Action: "scene:change:Scene 2"},
		{Name: "Test", Type: "text", Plugin: "obs", Action: "test"},
		{Name: "Mute Desktop Audio", Type: "image", ImageURL: "https://www.svgrepo.com/show/107080/mute.svg", Plugin: "obs", Action: "mic:mute:Desktop Audio"},
	}

	j, err := json.Marshal(api.buttons)
	if err != nil {
		log.Fatalf("Error marshaling data: %s", err)
	}
	log.Printf("Sending response %s\n", string(j))
	return &data.GetButtonsResponse{Body: string(j)}, nil
}

func (api *API) PressButton(ctx context.Context, request *data.PressButtonRequest) (*data.Empty, error) {
	if request.Plugin == "obs" {
		var action = strings.Split(request.Action, ":")

		if action[0] == "scene" {
			if action[1] == "change" {
				obs.ChangeScene(action[2])
			}
		}

		if action[0] == "mic" {
			if action[1] == "mute" {
				obs.ToggleMicMute(action[2])
			}
		}

		if action[0] == "test" {
			obs.Test()
		}
	}

	return &data.Empty{}, nil
}
