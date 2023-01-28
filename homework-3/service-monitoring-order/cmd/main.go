package main

import (
	"flag"
	"log"

	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/config"
	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/internal/app"
)

func main() {

	configPath := flag.String("config", "service-monitoring-order/config/config.yml", "choose config")
	flag.Parse()

	conf, err := config.New(*configPath)
	if err != nil {
		log.Fatalf("failed to get config: %s", err)
	}

	app.Run(conf)
}
