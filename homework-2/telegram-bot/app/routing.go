package app

import (
	"context"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/bagdatov/homework-2/telegram-bot/usecase"
)

func Routing(api *tgbotapi.BotAPI, bot usecase.Bot) {
	api.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		txt := update.Message.Text
		chatID := update.Message.Chat.ID
		user := update.Message.From.UserName
		reply := update.Message.ReplyToMessage

		switch {
		case strings.HasPrefix(txt, "/help"):

			go func() {
				resp := bot.Help()
				m := tgbotapi.NewMessage(update.Message.Chat.ID, resp)
				if _, err := api.Send(m); err != nil {
					log.Println(err)
				}
			}()

		case strings.HasPrefix(txt, "/score"):

			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
				defer cancel()

				resp := bot.Score(ctx, chatID)
				m := tgbotapi.NewMessage(chatID, resp)

				if _, err := api.Send(m); err != nil {
					log.Println(err)
				}
			}()

		case strings.HasPrefix(txt, "/next"):

			go func() {
				resp := bot.Next(chatID)
				m := tgbotapi.NewMessage(chatID, resp)

				if _, err := api.Send(m); err != nil {
					log.Println(err)
				}
			}()

		case strings.HasPrefix(txt, "/random"):

			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
				defer cancel()

				resp := bot.Random(ctx)
				m := tgbotapi.NewMessage(chatID, resp)

				if _, err := api.Send(m); err != nil {
					log.Println(err)
				}
			}()

		case strings.HasPrefix(txt, "/stop"):

			go func() {
				resp := bot.Stop(chatID)
				m := tgbotapi.NewMessage(chatID, resp)

				if _, err := api.Send(m); err != nil {
					log.Println(err)
				}
			}()

		case strings.HasPrefix(txt, "/begin "):

			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
				defer cancel()

				resp := bot.Begin(ctx, chatID, txt[7:])
				m := tgbotapi.NewMessage(chatID, resp)

				if _, err := api.Send(m); err != nil {
					log.Println(err)
				}
			}()

		case strings.HasPrefix(txt, "/desc "):

			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
				defer cancel()

				resp := bot.Details(ctx, txt[6:])
				m := tgbotapi.NewMessage(chatID, resp)

				if _, err := api.Send(m); err != nil {
					log.Println(err)
				}
			}()

		case reply != nil:

			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
				defer cancel()

				resp := bot.Submit(ctx, chatID, reply.MessageID, user, txt)
				m := tgbotapi.NewMessage(chatID, resp)

				if _, err := api.Send(m); err != nil {
					log.Println(err)
				}
			}()
		}
	}
}
