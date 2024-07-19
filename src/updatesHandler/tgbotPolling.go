package updateshandler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	updateprocesser "github.com/irvin518/tgbotpro/src/updateProcesser"
	updatetypes "github.com/irvin518/tgbotpro/src/updateTypes"
)

type TgbotPolling struct {
	botApi  *tgbotapi.BotAPI
	handler *UpdatesHandler
}

func NewTgbotPolling(bot *tgbotapi.BotAPI) *TgbotPolling {
	return &TgbotPolling{
		botApi:  bot,
		handler: NewUpdatesHandler(),
	}
}

func (m *TgbotPolling) Start(closeChan chan any) error {
	updateChan := m.botApi.GetUpdatesChan(tgbotapi.UpdateConfig{
		Limit:   1,
		Timeout: 0,
		Offset:  0,
	})

	go func() {
		for {
			select {
			case <-closeChan:
				return
			default:
			}
			select {
			case update := <-updateChan:
				m.handler.processUpdates(update)
			default:
			}

		}
	}()

	return nil
}

func (m *TgbotPolling) AddProcesser(t updatetypes.UpdateTypes, processer updateprocesser.UpdateProcesser) {
	m.handler.addProcesser(t, processer)
}
