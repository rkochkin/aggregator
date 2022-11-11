package agm_telegram

import (
	"aggregator/agm_telegram/app"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/exp/slices"
	"log"
)

type TelegramConfig struct {
	BotApiKey string  `json:"BotAPI"`
	ACL       []int64 `json:"ACL"`
}

type TelegramOnUpdateHandler func(message *tgbotapi.Message) tgbotapi.MessageConfig

type Telegram struct {
	Config TelegramConfig
	App    app.TelegramApp
}

func NewTelegram(config TelegramConfig, handler app.TelegramApp) Telegram {
	return Telegram{Config: config, App: handler}
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
			idx := slices.IndexFunc(t.Config.ACL, func(e int64) bool { return e == update.Message.From.ID })
			if idx != -1 {
				bot.Send(t.App.Proc(update.Message))
			}
		}
	}
}
