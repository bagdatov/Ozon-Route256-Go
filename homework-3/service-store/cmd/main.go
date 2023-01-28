package main

import (
	"flag"
	"log"

	"gitlab.ozon.dev/bagdatov/homework-3/service-store/config"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/app"
)

func main() {

	configPath := flag.String("config", "service-store/config/config.yml", "choose config")
	flag.Parse()

	conf, err := config.New(*configPath)
	if err != nil {
		log.Fatalf("failed to get config: %s", err)
	}

	app.Run(conf)
}
