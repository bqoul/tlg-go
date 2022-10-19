package evt

const (
	Message            Event = "message"
	EditedMessage      Event = "edited_message"
	ChannelPost        Event = "channel_post"
	EditedChannelPost  Event = "edited_channel_post"
	InlineQuery        Event = "inline_query"
	ChosenInlineResult Event = "chosen_inline_result"
	CallbackQuery      Event = "callback_query"
	ShippingQuery      Event = "shipping_query"
	PreCheckoutQuery   Event = "pre_checkout_query"
	PollStopped        Event = "poll"
	PollAnswer         Event = "poll_answer"
	MyChatMember       Event = "my_chat_member"
	ChatMember         Event = "chat_member"
	ChatJoinRequest    Event = "chat_join_request"
)
