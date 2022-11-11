package app

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelegramApp interface {
	Proc(message *tgbotapi.Message) tgbotapi.MessageConfig
}
