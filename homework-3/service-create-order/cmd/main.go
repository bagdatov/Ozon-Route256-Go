package main

import (
	"flag"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/config"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/internal/app"
	"log"
)

func main() {

	configPath := flag.String("config", "service-create-order/config/config.yml", "choose config")
	flag.Parse()

	conf, err := config.New(*configPath)
	if err != nil {
		log.Fatalf("failed to get config: %s", err)
	}

	app.Run(conf)
}
