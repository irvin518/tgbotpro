package updateshandler

import (
	"time"

	tgbotapi "github.com/irvin518/telegram-bot-api/v5"
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

func (m *TgbotPolling) Start(msgLimit int, closeChan chan any) error {
	updateChan := m.botApi.GetUpdatesChan(tgbotapi.UpdateConfig{
		Limit:   msgLimit,
		Timeout: 0,
		Offset:  0,
	})
	for i := 0; i < msgLimit; i++ {
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
					time.Sleep(time.Millisecond)
				}

			}
		}()
	}
	return nil
}

func (m *TgbotPolling) AddProcesser(t updatetypes.UpdateTypes, processer updateprocesser.UpdateProcesser) {
	m.handler.addProcesser(t, processer)
}
