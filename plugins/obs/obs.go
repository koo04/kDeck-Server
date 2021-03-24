package obs

import (
	"log"
	"strings"
	"time"

	obsws "github.com/christopher-dG/go-obs-websocket"
	"github.com/koo04/kdeck-server/settings"
)

type PluginSettings struct {
	settings.Base
	Host string
	Port int
}

var pluginSettings PluginSettings
var client obsws.Client
var recording bool

func Initialize() {
	pluginSettings = PluginSettings{}
	pluginSettings.Host = "localhost"
	pluginSettings.Port = 4444
	settings.LoadSettings("obs", &pluginSettings)
}

func Run() {
	if !client.Connected() {
		client = obsws.Client{Host: pluginSettings.Host, Port: pluginSettings.Port}
		if err := client.Connect(); err != nil {
			log.Println(err)
		}
	}

	obsws.SetReceiveTimeout(time.Second * 2)
}

func Stop() {
	defer client.Disconnect()
}

func IsEnabled() bool {
	return pluginSettings.Enabled
}

func RequestHandler(call string) {
	var action = strings.Split(call, ":")

	if action[0] == "scene" {
		if action[1] == "change" {
			changeScene(action[2])
		}
	}

	if action[0] == "mic" {
		if action[1] == "mute" {
			toggleMicMute(action[2])
		}
	}

	if action[0] == "record" {
		toggleRecording()
	}
}

func changeScene(sceneName string) {
	req := obsws.NewSetCurrentSceneRequest(sceneName)
	if err := req.Send(client); err != nil {
		log.Fatal(err)
	}

	resp, err := req.Receive()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Scene Change Status:", resp.Status())
}

func toggleMicMute(source string) {
	req := obsws.NewToggleMuteRequest(source)
	if err := req.Send(client); err != nil {
		log.Fatal(err)
	}

	req2 := obsws.NewGetMuteRequest(source)
	if err := req2.Send(client); err != nil {
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

func getScenes() {
	req := obsws.NewGetSceneListRequest()
	if err := req.Send(client); err != nil {
		log.Fatal(err)
	}

	res, err := req.Receive()
	if err != nil {
		log.Fatal(err)
	}

	for _, scene := range res.Scenes {
		log.Println(scene.Name)
	}
}

func toggleRecording() bool {
	req := obsws.NewStartStopRecordingRequest()
	if err := req.Send(client); err != nil {
		log.Fatal(err)
	}

	recording = !recording
	return recording
}
