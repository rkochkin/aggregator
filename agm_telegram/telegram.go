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

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("вкл"),
		tgbotapi.NewKeyboardButton("выкл"),
		tgbotapi.NewKeyboardButton("закат"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ярк 10"),
		tgbotapi.NewKeyboardButton("ярк 50"),
		tgbotapi.NewKeyboardButton("ярк 100"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("выйти"),
	),
)

func (t *Telegram) Run() {
	bot, err := tgbotapi.NewBotAPI(t.Config.BotApiKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	cfg := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{Command: "/light", Description: "Освещение"},
	)

	bot.Request(cfg)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			idx := slices.IndexFunc(t.Config.ACL, func(e int64) bool { return e == update.Message.From.ID })
			if idx != -1 {
				switch update.Message.Text {
				case "/light":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					msg.ReplyMarkup = numericKeyboard
					bot.Send(msg)
				case "выйти":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msg)
				default:
					bot.Send(t.App.Proc(update.Message))
				}

			}
		}
	}
}
