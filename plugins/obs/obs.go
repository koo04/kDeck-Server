package obs

import (
	"log"
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

func ChangeScene(sceneName string) {
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

func ToggleMicMute(source string) {
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

func GetScenes() {
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
