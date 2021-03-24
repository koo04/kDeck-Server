package plugins

import (
	"log"

	"github.com/koo04/kdeck-server/plugins/obs"
)

func Load() {
	obs.Initialize()

	log.Println(obs.IsEnabled())

	if obs.IsEnabled() {
		obs.Run()
	}
}
