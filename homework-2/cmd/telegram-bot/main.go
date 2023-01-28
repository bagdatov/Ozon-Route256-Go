package main

import (
	"flag"
	"gitlab.ozon.dev/bagdatov/homework-2/telegram-bot/app"
	"gitlab.ozon.dev/bagdatov/homework-2/telegram-bot/config"
	"log"
)

func main() {

	configPath := flag.String("config", "configs/config_telegram_bot.yml", "choose config")

	conf, err := config.New(*configPath)
	if err != nil {
		log.Fatalf("failed to get config: %s", err)
	}

	app.Run(conf)

}
