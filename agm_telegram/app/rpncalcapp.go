package app

import (
	"aggregator/rpncalc"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang-collections/collections/stack"
	"log"
)

type RpnCalcApp struct {
}

func (r RpnCalcApp) Proc(message *tgbotapi.Message) tgbotapi.MessageConfig {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	st := stack.New()
	result := rpncalc.Rpn(message.Text, st)
	msg := tgbotapi.NewMessage(message.Chat.ID, result)
	//msg.ReplyToMessageID = update.Message.MessageID

	return msg
}
