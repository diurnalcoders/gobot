package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Client struct {
	Bot *tgbotapi.BotAPI
}

func NewClient(token string) *Client {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	return &Client{Bot: bot}
}
