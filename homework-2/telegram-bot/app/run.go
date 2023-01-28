package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"gitlab.ozon.dev/bagdatov/homework-2/telegram-bot/config"
	"gitlab.ozon.dev/bagdatov/homework-2/telegram-bot/repository/game"
	"gitlab.ozon.dev/bagdatov/homework-2/telegram-bot/repository/messenger"
	"gitlab.ozon.dev/bagdatov/homework-2/telegram-bot/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func Run(conf *config.Config) {

	log.Println("Dialing with grpc game server...")
	conn, err := grpc.Dial(conf.GRPC.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect gRPC: %s", err)
	}
	log.Println("Dialed")

	defer conn.Close()

	client := pb.NewChgkClient(conn)

	log.Println("Dialing with Telegram API...")
	botAPI, err := tgbotapi.NewBotAPI(conf.Telegram.Key)
	if err != nil {
		log.Fatalf("failed to connect to messenger %s", err)
	}
	log.Println("Dialed")

	m := messenger.New(botAPI)
	g := game.New(client)
	bot := usecase.New(g, m)

	// Start listening to messages
	log.Println("Ready")
	Routing(botAPI, bot)
}
