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
	msgText := "ok"
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	bulbIp := "192.168.1.240"
	bulb := yl.NewBulb(bulbIp)

	err := bulb.Connect()
	if err != nil {
		msgText = fmt.Sprintln(err)
		goto error
	}

	defer func(bulb *yl.Bulb) {
		err := bulb.Disconnect()
		if err != nil {
			fmt.Println(err)
		}
	}(bulb)

	switch message.Text {
	case "вкл":
		err = bulb.PowerOn(1000)

	case "выкл":
		err = bulb.PowerOff(1000)
	case "закат":
		var fss []yl.FlowState

		fss = append(fss, yl.NewFlowState(5000, yl.CF_MODE_TEMP, 2800, 100))
		fss = append(fss, yl.NewFlowState(3*60000, yl.CF_MODE_TEMP, 2000, 20))
		fss = append(fss, yl.NewFlowState(2*60000, yl.CF_MODE_COLOR, 0xFF5202, 1))
		fss = append(fss, yl.NewFlowState(10000, yl.CF_MODE_SLEEP, 0, 0))

		err, fe := yl.NewFlowExpression(fss...)
		if err == nil {
			err = bulb.StartColorFlow(len(fss), yl.CF_ACTION_POWEROFF, fe)
		}
	case "ярк 100":
		err = bulb.Brightness(100, 1000)

	case "ярк 50":
		err = bulb.Brightness(50, 1000)

	case "ярк 10":
		err = bulb.Brightness(10, 1000)
	}
	if err != nil {
		msgText = fmt.Sprintln(err)
	}

error:
	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)

	return msg
}
