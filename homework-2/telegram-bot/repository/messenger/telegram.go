package messenger

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type messenger struct {
	api *tgbotapi.BotAPI
}

func New(api *tgbotapi.BotAPI) *messenger {

	return &messenger{
		api: api,
	}
}

func (m *messenger) Send(chatID int64, text string) (int, error) {
	convert := tgbotapi.NewMessage(chatID, text)
	message, err := m.api.Send(convert)

	return message.MessageID, err
}
