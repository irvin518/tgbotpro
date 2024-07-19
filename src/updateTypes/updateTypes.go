package updatetypes

type UpdateTypes int

// const message = update.message;
//     const editedMessage = update.edited_message;
//     const channelPost = update.channel_post;
//     const editedChannelPost = update.edited_channel_post;
//     const inlineQuery = update.inline_query;
//     const chosenInlineResult = update.chosen_inline_result;
//     const callbackQuery = update.callback_query;
//     const shippingQuery = update.shipping_query;
//     const preCheckoutQuery = update.pre_checkout_query;
//     const poll = update.poll;
//     const pollAnswer = update.poll_answer;
//     const chatMember = update.chat_member;
//     const myChatMember = update.my_chat_member;
//     const chatJoinRequest = update.chat_join_request;

const (
	Message UpdateTypes = 0
	EidtMessage
	ChannelPost
	EditedChannelPost
	InlineQuery
	ChosenInlineResult
	CallbackQuery
	ShippingQuery
	PreCheckoutQuery
	Poll
	PollAnsewer
	ChatMember
	MyChatMember
	ChatJoinRequest
)
