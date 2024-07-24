package updateprocesser

import (
	"fmt"

	tgbotapi "github.com/irvin518/telegram-bot-api/v5"
)

func checkMessage[T any](message any) error {
	if _, ok := message.(T); !ok {
		return fmt.Errorf("message type does not match interface method parameter type")
	}
	return nil
}

type UpdateProcesser interface {
	Process(*tgbotapi.BotAPI, any) error
}

type MessageProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.Message)
}

type MessageProcesserAdpter struct {
	processer MessageProcesser
}

func NewMessageProcesserAdpter(processer MessageProcesser) *MessageProcesserAdpter {
	return &MessageProcesserAdpter{
		processer: processer,
	}
}

func (m *MessageProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.Message](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.Message))
	}
	return nil
}

type InlineQueryProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.InlineQuery)
}

type InlineQueryProcesserAdpter struct {
	processer InlineQueryProcesser
}

func NewInlineQueryProcesserAdpter(processer InlineQueryProcesser) *InlineQueryProcesserAdpter {
	return &InlineQueryProcesserAdpter{
		processer: processer,
	}
}

func (m *InlineQueryProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.InlineQuery](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.InlineQuery))
	}
	return nil
}

// ChosenInlineResult
type ChosenInlineResultProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.ChosenInlineResult)
}

type ChosenInlineResultProcesserAdpter struct {
	processer ChosenInlineResultProcesser
}

func NewChosenInlineResultProcesserAdpter(processer ChosenInlineResultProcesser) *ChosenInlineResultProcesserAdpter {
	return &ChosenInlineResultProcesserAdpter{
		processer: processer,
	}
}

func (m *ChosenInlineResultProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.ChosenInlineResult](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.ChosenInlineResult))
	}
	return nil
}

// CallbackQuery
type CallbackQueryProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.CallbackQuery)
}

type CallbackQueryProcesserAdpter struct {
	processer CallbackQueryProcesser
}

func NewCallbackQueryProcesserAdpter(processer CallbackQueryProcesser) *CallbackQueryProcesserAdpter {
	return &CallbackQueryProcesserAdpter{
		processer: processer,
	}
}

func (m *CallbackQueryProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.CallbackQuery](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.CallbackQuery))
	}
	return nil
}

// ShippingQuery
type ShippingQueryProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.ShippingQuery)
}

type ShippingQueryProcesserAdpter struct {
	processer ShippingQueryProcesser
}

func NewShippingQueryProcesserAdpter(processer ShippingQueryProcesser) *ShippingQueryProcesserAdpter {
	return &ShippingQueryProcesserAdpter{
		processer: processer,
	}
}

func (m *ShippingQueryProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.ShippingQuery](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.ShippingQuery))
	}
	return nil
}

// PreCheckoutQuery
type PreCheckoutQueryProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.PreCheckoutQuery)
}

type PreCheckoutQueryProcesserAdpter struct {
	processer PreCheckoutQueryProcesser
}

func NewPreCheckoutQueryProcesserAdpter(processer PreCheckoutQueryProcesser) *PreCheckoutQueryProcesserAdpter {
	return &PreCheckoutQueryProcesserAdpter{
		processer: processer,
	}
}

func (m *PreCheckoutQueryProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.PreCheckoutQuery](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.PreCheckoutQuery))
	}
	return nil
}

//Poll

type PollProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.Poll)
}

type PollProcesserAdpter struct {
	processer PollProcesser
}

func NewPollProcesserAdpter(processer PollProcesser) *PollProcesserAdpter {
	return &PollProcesserAdpter{
		processer: processer,
	}
}

func (m *PollProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.Poll](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.Poll))
	}
	return nil
}

// PollAnswer
type PollAnswerProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.PollAnswer)
}

type PollAnswerProcesserAdpter struct {
	processer PollAnswerProcesser
}

func NewPollAnswerProcesserAdpter(processer PollAnswerProcesser) *PollAnswerProcesserAdpter {
	return &PollAnswerProcesserAdpter{
		processer: processer,
	}
}

func (m *PollAnswerProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.Poll](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.PollAnswer))
	}
	return nil
}

// ChatMemberUpdated
type ChatMemberUpdatedProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.ChatMemberUpdated)
}

type ChatMemberUpdatedProcesserAdpter struct {
	processer ChatMemberUpdatedProcesser
}

func NewChatMemberUpdatedProcesserAdpter(processer ChatMemberUpdatedProcesser) *ChatMemberUpdatedProcesserAdpter {
	return &ChatMemberUpdatedProcesserAdpter{
		processer: processer,
	}
}

func (m *ChatMemberUpdatedProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.ChatMemberUpdated](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.ChatMemberUpdated))
	}
	return nil
}

// ChatJoinRequest
type ChatJoinRequestProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.ChatJoinRequest)
}

type ChatJoinRequestProcesserAdpter struct {
	processer ChatJoinRequestProcesser
}

func NewChatJoinRequestProcesserAdpter(processer ChatJoinRequestProcesser) *ChatJoinRequestProcesserAdpter {
	return &ChatJoinRequestProcesserAdpter{
		processer: processer,
	}
}

func (m *ChatJoinRequestProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.ChatJoinRequest](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.ChatJoinRequest))
	}
	return nil
}

// chart_boost
type ChatBoostProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.ChatBoostUpdated)
}

type ChatBoostProcesserAdpter struct {
	processer ChatBoostProcesser
}

func NewChatBoostProcesserAdpter(processer ChatBoostProcesser) *ChatBoostProcesserAdpter {
	return &ChatBoostProcesserAdpter{
		processer: processer,
	}
}

func (m *ChatBoostProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.ChatJoinRequest](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.ChatBoostUpdated))
	}
	return nil
}

// remove_chart_boost
type RemoveChatBoostProcesser interface {
	Process(*tgbotapi.BotAPI, *tgbotapi.ChatBoostRemoved)
}

type RemoveChatBoostProcesserAdpter struct {
	processer RemoveChatBoostProcesser
}

func NewRemoveChatBoostProcesserAdpter(processer RemoveChatBoostProcesser) *RemoveChatBoostProcesserAdpter {
	return &RemoveChatBoostProcesserAdpter{
		processer: processer,
	}
}

func (m *RemoveChatBoostProcesserAdpter) Process(api *tgbotapi.BotAPI, msg any) error {
	err := checkMessage[*tgbotapi.ChatJoinRequest](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(api, msg.(*tgbotapi.ChatBoostRemoved))
	}
	return nil
}
