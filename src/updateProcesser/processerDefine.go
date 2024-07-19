package updateprocesser

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func checkMessage[T any](message any) error {
	if _, ok := message.(T); !ok {
		return fmt.Errorf("message type does not match interface method parameter type")
	}
	return nil
}

type UpdateProcesser interface {
	Process(any) error
}

type MessageProcesser interface {
	Process(*tgbotapi.Message)
}

type MessageProcesserAdpter struct {
	processer MessageProcesser
}

func NewMessageProcesserAdpter(processer MessageProcesser) *MessageProcesserAdpter {
	return &MessageProcesserAdpter{
		processer: processer,
	}
}

func (m *MessageProcesserAdpter) Process(msg any) error {
	err := checkMessage[*tgbotapi.Message](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(msg.(*tgbotapi.Message))
	}
	return nil
}

type InlineQueryProcesser interface {
	Process(*tgbotapi.InlineQuery)
}

type InlineQueryProcesserAdpter struct {
	processer InlineQueryProcesser
}

func NewInlineQueryProcesserAdpter(processer InlineQueryProcesser) *InlineQueryProcesserAdpter {
	return &InlineQueryProcesserAdpter{
		processer: processer,
	}
}

func (m *InlineQueryProcesserAdpter) Process(msg any) error {
	err := checkMessage[*tgbotapi.InlineQuery](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(msg.(*tgbotapi.InlineQuery))
	}
	return nil
}

// ChosenInlineResult
type ChosenInlineResultProcesser interface {
	Process(*tgbotapi.ChosenInlineResult)
}

type ChosenInlineResultProcesserAdpter struct {
	processer ChosenInlineResultProcesser
}

func NewChosenInlineResultProcesserAdpter(processer ChosenInlineResultProcesser) *ChosenInlineResultProcesserAdpter {
	return &ChosenInlineResultProcesserAdpter{
		processer: processer,
	}
}

func (m *ChosenInlineResultProcesserAdpter) Process(msg any) error {
	err := checkMessage[*tgbotapi.ChosenInlineResult](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(msg.(*tgbotapi.ChosenInlineResult))
	}
	return nil
}

// CallbackQuery
type CallbackQueryProcesser interface {
	Process(*tgbotapi.CallbackQuery)
}

type CallbackQueryProcesserAdpter struct {
	processer CallbackQueryProcesser
}

func NewCallbackQueryProcesserAdpter(processer CallbackQueryProcesser) *CallbackQueryProcesserAdpter {
	return &CallbackQueryProcesserAdpter{
		processer: processer,
	}
}

func (m *CallbackQueryProcesserAdpter) Process(msg any) error {
	err := checkMessage[*tgbotapi.CallbackQuery](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(msg.(*tgbotapi.CallbackQuery))
	}
	return nil
}

// ShippingQuery
type ShippingQueryProcesser interface {
	Process(*tgbotapi.ShippingQuery)
}

type ShippingQueryProcesserAdpter struct {
	processer ShippingQueryProcesser
}

func NewShippingQueryProcesserAdpter(processer ShippingQueryProcesser) *ShippingQueryProcesserAdpter {
	return &ShippingQueryProcesserAdpter{
		processer: processer,
	}
}

func (m *ShippingQueryProcesserAdpter) Process(msg any) error {
	err := checkMessage[*tgbotapi.ShippingQuery](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(msg.(*tgbotapi.ShippingQuery))
	}
	return nil
}

// PreCheckoutQuery
type PreCheckoutQueryProcesser interface {
	Process(*tgbotapi.PreCheckoutQuery)
}

type PreCheckoutQueryProcesserAdpter struct {
	processer PreCheckoutQueryProcesser
}

func NewPreCheckoutQueryProcesserAdpter(processer PreCheckoutQueryProcesser) *PreCheckoutQueryProcesserAdpter {
	return &PreCheckoutQueryProcesserAdpter{
		processer: processer,
	}
}

func (m *PreCheckoutQueryProcesserAdpter) Process(msg any) error {
	err := checkMessage[*tgbotapi.PreCheckoutQuery](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(msg.(*tgbotapi.PreCheckoutQuery))
	}
	return nil
}

//Poll

type PollProcesser interface {
	Process(*tgbotapi.Poll)
}

type PollProcesserAdpter struct {
	processer PollProcesser
}

func NewPollProcesserAdpter(processer PollProcesser) *PollProcesserAdpter {
	return &PollProcesserAdpter{
		processer: processer,
	}
}

func (m *PollProcesserAdpter) Process(msg any) error {
	err := checkMessage[*tgbotapi.Poll](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(msg.(*tgbotapi.Poll))
	}
	return nil
}

// PollAnswer
type PollAnswerProcesser interface {
	Process(*tgbotapi.PollAnswer)
}

type PollAnswerProcesserAdpter struct {
	processer PollAnswerProcesser
}

func NewPollAnswerProcesserAdpter(processer PollAnswerProcesser) *PollAnswerProcesserAdpter {
	return &PollAnswerProcesserAdpter{
		processer: processer,
	}
}

func (m *PollAnswerProcesserAdpter) Process(msg any) error {
	err := checkMessage[*tgbotapi.Poll](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(msg.(*tgbotapi.PollAnswer))
	}
	return nil
}

// ChatMemberUpdated
type ChatMemberUpdatedProcesser interface {
	Process(*tgbotapi.ChatMemberUpdated)
}

type ChatMemberUpdatedProcesserAdpter struct {
	processer ChatMemberUpdatedProcesser
}

func NewChatMemberUpdatedProcesserAdpter(processer ChatMemberUpdatedProcesser) *ChatMemberUpdatedProcesserAdpter {
	return &ChatMemberUpdatedProcesserAdpter{
		processer: processer,
	}
}

func (m *ChatMemberUpdatedProcesserAdpter) Process(msg any) error {
	err := checkMessage[*tgbotapi.ChatMemberUpdated](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(msg.(*tgbotapi.ChatMemberUpdated))
	}
	return nil
}

// ChatJoinRequest
type ChatJoinRequestProcesser interface {
	Process(*tgbotapi.ChatJoinRequest)
}

type ChatJoinRequestProcesserAdpter struct {
	processer ChatJoinRequestProcesser
}

func NewChatJoinRequestProcesserAdpter(processer ChatJoinRequestProcesser) *ChatJoinRequestProcesserAdpter {
	return &ChatJoinRequestProcesserAdpter{
		processer: processer,
	}
}

func (m *ChatJoinRequestProcesserAdpter) Process(msg any) error {
	err := checkMessage[*tgbotapi.ChatJoinRequest](msg)
	if err != nil {
		return err
	}
	if m.processer != nil {
		m.processer.Process(msg.(*tgbotapi.ChatJoinRequest))
	}
	return nil
}
