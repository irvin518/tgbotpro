package updateshandler

import (
	"encoding/json"

	tgbotapi "github.com/irvin518/telegram-bot-api/v5"
	updateprocesser "github.com/irvin518/tgbotpro/src/updateProcesser"
	updatetypes "github.com/irvin518/tgbotpro/src/updateTypes"
	"github.com/irvin518/tgbotpro/src/utils"
)

type UpdatesHandler struct {
	updateDisposer map[updatetypes.UpdateTypes][]updateprocesser.UpdateProcesser
}

func NewUpdatesHandler() *UpdatesHandler {
	return &UpdatesHandler{
		updateDisposer: make(map[updatetypes.UpdateTypes][]updateprocesser.UpdateProcesser),
	}
}

func (m *UpdatesHandler) processUpdates(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var messageType updatetypes.UpdateTypes = -1
	var message any
	if update.Message != nil {
		messageType = updatetypes.Message
		message = update.Message
	} else if update.EditedMessage != nil {
		messageType = updatetypes.EidtMessage
		message = update.EditedMessage
	} else if update.ChannelPost != nil {
		messageType = updatetypes.ChannelPost
		message = update.ChannelPost
	} else if update.EditedChannelPost != nil {
		messageType = updatetypes.EditedChannelPost
		message = update.EditedChannelPost
	} else if update.InlineQuery != nil {
		messageType = updatetypes.InlineQuery
		message = update.InlineQuery
	} else if update.ChosenInlineResult != nil {
		messageType = updatetypes.ChosenInlineResult
		message = update.ChosenInlineResult
	} else if update.CallbackQuery != nil {
		messageType = updatetypes.CallbackQuery
		message = update.CallbackQuery
	} else if update.ShippingQuery != nil {
		messageType = updatetypes.ShippingQuery
		message = update.ShippingQuery
	} else if update.PreCheckoutQuery != nil {
		messageType = updatetypes.PreCheckoutQuery
		message = update.PreCheckoutQuery
	} else if update.Poll != nil {
		messageType = updatetypes.Poll
		message = update.Poll
	} else if update.PollAnswer != nil {
		messageType = updatetypes.PollAnsewer
		message = update.PollAnswer
	} else if update.ChatMember != nil {
		messageType = updatetypes.ChatMember
		message = update.ChatMember
	} else if update.MyChatMember != nil {
		messageType = updatetypes.MyChatMember
		message = update.MyChatMember
	} else if update.ChatJoinRequest != nil {
		messageType = updatetypes.ChatJoinRequest
		message = update.ChatJoinRequest
	} else if update.ChatBoost != nil {
		messageType = updatetypes.ChatBoost
		message = update.ChatBoost
	} else if update.RemovedChatBoost != nil {
		messageType = updatetypes.RemoveChatBoost
		message = update.RemovedChatBoost
	}

	if disposer, ok := m.updateDisposer[messageType]; ok {
		for _, processer := range disposer {
			err := processer.Process(bot, message)
			if err != nil {
				msg, _ := json.Marshal(update)
				utils.Log().Error("process update message %s error %s", msg, err)
			}
		}
	} else {
		utils.Log().Error("mis dispose update %+v , message type %d message %+v", update, messageType, message)
	}
}

func (m *UpdatesHandler) addProcesser(t updatetypes.UpdateTypes, processer updateprocesser.UpdateProcesser) {
	m.updateDisposer[t] = append(m.updateDisposer[t], processer)
}
