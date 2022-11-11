package app

import (
	"fmt"
	yl "github.com/gethiox/yeelight-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type YeeLightApp struct {
}

func (a YeeLightApp) Proc(message *tgbotapi.Message) tgbotapi.MessageConfig {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	bulbIp := "192.168.1.240"
	bulb := yl.NewBulb(bulbIp)

	err := bulb.Connect()
	if err != nil {
		panic(err)
	}

	switch message.Text {
	case "/on":
		err = bulb.PowerOn(1000)
		if err != nil {
			fmt.Println(err)
		}
	case "/off":
		err = bulb.PowerOff(1000)
		if err != nil {
			fmt.Println(err)
		}
	case "/sunset":
		err = bulb.Brightness(100, 1000)
		if err != nil {
			fmt.Println(err)
		}
	}

	bulb.Disconnect()
	msg := tgbotapi.NewMessage(message.Chat.ID, "ok")
	//msg.ReplyToMessageID = update.Message.MessageID

	return msg
}
