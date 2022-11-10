package agm_telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type TelegramConfig struct {
	BotApiKey string `json:"BotAPI"`
}

type TelegramOnUpdateHandler func(message *tgbotapi.Message) tgbotapi.MessageConfig

type Telegram struct {
	Config          TelegramConfig
	OnUpdateHandler TelegramOnUpdateHandler
}

func NewTelegram(config TelegramConfig, handler TelegramOnUpdateHandler) Telegram {
	return Telegram{Config: config, OnUpdateHandler: handler}
}

func (t *Telegram) Run() {
	bot, err := tgbotapi.NewBotAPI(t.Config.BotApiKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			bot.Send(t.OnUpdateHandler(update.Message))
		}
	}
}
