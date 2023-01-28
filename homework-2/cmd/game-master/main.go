package main

import (
	"flag"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/app"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/config"
	"log"
)

func main() {

	configPath := flag.String("config", "configs/config_game_master.yml", "choose config")

	conf, err := config.New(*configPath)
	if err != nil {
		log.Fatalf("failed to get config: %s", err)
	}

	app.Run(conf)
}
