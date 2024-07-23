package tgbotpro

import (
	tgbotapi "github.com/irvin518/telegram-bot-api/v5"
	updateprocesser "github.com/irvin518/tgbotpro/src/updateProcesser"
	updatetypes "github.com/irvin518/tgbotpro/src/updateTypes"
	updateshandler "github.com/irvin518/tgbotpro/src/updatesHandler"
	"github.com/irvin518/tgbotpro/src/utils"
)

type TgbotPro struct {
	apiImpl    *tgbotapi.BotAPI
	updateImpl BotUpdate
	shutdown   chan any
}

type BotUpdate interface {
	Start(closeChan chan any) error

	AddProcesser(t updatetypes.UpdateTypes, processer updateprocesser.UpdateProcesser)
}

func (m *TgbotPro) Start(logInfo utils.LogInfoFunc, logError utils.LogErrorFunc) error {
	utils.Log().Init(logInfo, logError)
	return m.updateImpl.Start(m.shutdown)
}

func (m *TgbotPro) Stop() {
	close(m.shutdown)
}

func (m *TgbotPro) GetBotApi() *tgbotapi.BotAPI {
	return m.apiImpl
}

func (m *TgbotPro) OnMessage(processer updateprocesser.MessageProcesser) {
	m.updateImpl.AddProcesser(updatetypes.Message, updateprocesser.NewMessageProcesserAdpter(processer))
}

func (m *TgbotPro) OnEditMessage(processer updateprocesser.MessageProcesser) {
	m.updateImpl.AddProcesser(updatetypes.EidtMessage, updateprocesser.NewMessageProcesserAdpter(processer))
}

func (m *TgbotPro) OnChannelPost(processer updateprocesser.MessageProcesser) {
	m.updateImpl.AddProcesser(updatetypes.ChannelPost, updateprocesser.NewMessageProcesserAdpter(processer))
}

func (m *TgbotPro) OnEdithannelPost(processer updateprocesser.MessageProcesser) {
	m.updateImpl.AddProcesser(updatetypes.EditedChannelPost, updateprocesser.NewMessageProcesserAdpter(processer))
}

func (m *TgbotPro) OnInlineQuery(processer updateprocesser.InlineQueryProcesser) {
	m.updateImpl.AddProcesser(updatetypes.InlineQuery, updateprocesser.NewInlineQueryProcesserAdpter(processer))
}

func (m *TgbotPro) OnChosenInlineResult(processer updateprocesser.ChosenInlineResultProcesser) {
	m.updateImpl.AddProcesser(updatetypes.ChosenInlineResult, updateprocesser.NewChosenInlineResultProcesserAdpter(processer))
}

func (m *TgbotPro) OnCallbackQuery(processer updateprocesser.CallbackQueryProcesser) {
	m.updateImpl.AddProcesser(updatetypes.CallbackQuery, updateprocesser.NewCallbackQueryProcesserAdpter(processer))
}

func (m *TgbotPro) OnShippingQuery(processer updateprocesser.ShippingQueryProcesser) {
	m.updateImpl.AddProcesser(updatetypes.ShippingQuery, updateprocesser.NewShippingQueryProcesserAdpter(processer))
}

func (m *TgbotPro) OnPreCheckoutQuery(processer updateprocesser.PreCheckoutQueryProcesser) {
	m.updateImpl.AddProcesser(updatetypes.PreCheckoutQuery, updateprocesser.NewPreCheckoutQueryProcesserAdpter(processer))
}

func (m *TgbotPro) OnPoll(processer updateprocesser.PollProcesser) {
	m.updateImpl.AddProcesser(updatetypes.Poll, updateprocesser.NewPollProcesserAdpter(processer))
}

func (m *TgbotPro) OnPollAnswer(processer updateprocesser.PollAnswerProcesser) {
	m.updateImpl.AddProcesser(updatetypes.PollAnsewer, updateprocesser.NewPollAnswerProcesserAdpter(processer))
}

func (m *TgbotPro) OnChatMemberUpdated(processer updateprocesser.ChatMemberUpdatedProcesser) {
	m.updateImpl.AddProcesser(updatetypes.ChatMember, updateprocesser.NewChatMemberUpdatedProcesserAdpter(processer))
}

func (m *TgbotPro) OnMyChatMemberUpdated(processer updateprocesser.ChatMemberUpdatedProcesser) {
	m.updateImpl.AddProcesser(updatetypes.MyChatMember, updateprocesser.NewChatMemberUpdatedProcesserAdpter(processer))
}

func (m *TgbotPro) OnChatJoinRequest(processer updateprocesser.ChatJoinRequestProcesser) {
	m.updateImpl.AddProcesser(updatetypes.ChatJoinRequest, updateprocesser.NewChatJoinRequestProcesserAdpter(processer))
}

func (m *TgbotPro) OnChatBoost(processer updateprocesser.ChatBoostProcesser) {
	m.updateImpl.AddProcesser(updatetypes.MyChatMember, updateprocesser.NewChatBoostProcesserAdpter(processer))
}

func (m *TgbotPro) OnRemoveChatBoost(processer updateprocesser.RemoveChatBoostProcesser) {
	m.updateImpl.AddProcesser(updatetypes.ChatJoinRequest, updateprocesser.NewRemoveChatBoostProcesserAdpter(processer))
}

func NewTgbotPolling(botToken string) (*TgbotPro, error) {
	api, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}
	return &TgbotPro{
		apiImpl:    api,
		updateImpl: updateshandler.NewTgbotPolling(api),
		shutdown:   make(chan any),
	}, nil
}
