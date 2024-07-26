package main

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/irvin518/telegram-bot-api/v5"
	"github.com/irvin518/tgbotpro"
)

type MessageDisposser struct {
}

func (m *MessageDisposser) Process(api *tgbotapi.BotAPI, msg *tgbotapi.Message) {

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Errorf("err: %s", err)
	}

	fmt.Print(string(data))
}

func main() {
	bot, err := tgbotpro.NewTgbotPolling("6368243461:AAEsK6jbxpA05YpGicleSxe6NatjtR6DnAg")
	if err != nil {
		fmt.Errorf("new error %s", err)
	}

	bot.OnMessage(new(MessageDisposser))

	bot.Start(10, nil, nil)

	defer func() {
		bot.Stop()
	}()

	select {}
}
