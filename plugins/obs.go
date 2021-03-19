package plugins

import (
	"log"
	"time"

	obsws "github.com/christopher-dG/go-obs-websocket"
	"github.com/koo04/kdeck-server/config"
)

type OBS struct {
	settings *config.OBS
	client   obsws.Client
}

func (obs *OBS) OBS(config *config.OBS) (obsws.Client, error) {
	obs.settings = config

	if !obs.client.Connected() {
		obs.client = obsws.Client{Host: "localhost", Port: 4444}
		if err := obs.client.Connect(); err != nil {
			return obs.client, err
		}
	}

	obsws.SetReceiveTimeout(time.Second * 2)
	return obs.client, nil
}

func (obs *OBS) CloseClient() {
	defer obs.client.Disconnect()
}

func (obs *OBS) ChangeScene(sceneName string) {
	req := obsws.NewSetCurrentSceneRequest(sceneName)
	if err := req.Send(obs.client); err != nil {
		log.Fatal(err)
	}

	resp, err := req.Receive()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Scene Change Status:", resp.Status())
}

func (obs *OBS) ToggleMicMute(source string) {
	req := obsws.NewToggleMuteRequest(source)
	if err := req.Send(obs.client); err != nil {
		log.Fatal(err)
	}

	req2 := obsws.NewGetMuteRequest(source)
	if err := req2.Send(obs.client); err != nil {
		log.Fatal(err)
	}

	resp, err := req2.Receive()
	if err != nil {
		log.Fatal(err)
	}
	if resp.Muted {
		log.Printf("%s is muted", source)
	} else {
		log.Printf("%s is unmuted", source)
	}
}
