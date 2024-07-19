package updateshandler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	updateprocesser "github.com/irvin518/tgbotpro/src/updateProcesser"
	updatetypes "github.com/irvin518/tgbotpro/src/updateTypes"
)

type UpdatesHandler struct {
	updateDisposer map[updatetypes.UpdateTypes][]updateprocesser.UpdateProcesser
}

func NewUpdatesHandler() *UpdatesHandler {
	return &UpdatesHandler{
		updateDisposer: make(map[updatetypes.UpdateTypes][]updateprocesser.UpdateProcesser),
	}
}

func (m *UpdatesHandler) processUpdates(update tgbotapi.Update) {
	var messageType updatetypes.UpdateTypes = -1
	if update.Message != nil {
		messageType = updatetypes.Message
	} else if update.EditedMessage != nil {
		messageType = updatetypes.EidtMessage
	} else if update.ChannelPost != nil {
		messageType = updatetypes.ChannelPost
	} else if update.EditedChannelPost != nil {
		messageType = updatetypes.EditedChannelPost
	} else if update.InlineQuery != nil {
		messageType = updatetypes.InlineQuery
	} else if update.ChosenInlineResult != nil {
		messageType = updatetypes.ChosenInlineResult
	} else if update.CallbackQuery != nil {
		messageType = updatetypes.CallbackQuery
	} else if update.ShippingQuery != nil {
		messageType = updatetypes.ShippingQuery
	} else if update.PreCheckoutQuery != nil {
		messageType = updatetypes.PreCheckoutQuery
	} else if update.Poll != nil {
		messageType = updatetypes.Poll
	} else if update.PollAnswer != nil {
		messageType = updatetypes.PollAnsewer
	} else if update.ChatMember != nil {
		messageType = updatetypes.ChatMember
	} else if update.MyChatMember != nil {
		messageType = updatetypes.MyChatMember
	} else if update.ChatJoinRequest != nil {
		messageType = updatetypes.ChatJoinRequest
	}
	if disposer, ok := m.updateDisposer[messageType]; ok {
		for _, processer := range disposer {
			processer.Process(update.Message)
		}
	}
}

func (m *UpdatesHandler) addProcesser(t updatetypes.UpdateTypes, processer updateprocesser.UpdateProcesser) {
	m.updateDisposer[t] = append(m.updateDisposer[t], processer)
}
