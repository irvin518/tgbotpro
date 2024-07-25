package updateshandler

import (
	tgbotapi "github.com/irvin518/telegram-bot-api/v5"
	updateprocesser "github.com/irvin518/tgbotpro/src/updateProcesser"
	updatetypes "github.com/irvin518/tgbotpro/src/updateTypes"
)

type TgbotPolling struct {
	botApi  *tgbotapi.BotAPI
	handler *UpdatesHandler
}

const (
	Limits = 10
)

func NewTgbotPolling(bot *tgbotapi.BotAPI) *TgbotPolling {
	return &TgbotPolling{
		botApi:  bot,
		handler: NewUpdatesHandler(),
	}
}

func (m *TgbotPolling) Start(closeChan chan any) error {
	updateChan := m.botApi.GetUpdatesChan(tgbotapi.UpdateConfig{
		Limit:   Limits,
		Timeout: 0,
		Offset:  0,
	})
	for i := 0; i < Limits; i++ {
		go func() {
			for {
				select {
				case <-closeChan:
					return
				default:
				}
				select {
				case update := <-updateChan:
					m.handler.processUpdates(m.botApi, update)
				default:
				}

			}
		}()
	}
	return nil
}

func (m *TgbotPolling) AddProcesser(t updatetypes.UpdateTypes, processer updateprocesser.UpdateProcesser) {
	m.handler.addProcesser(t, processer)
}
