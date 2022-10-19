package upd

type Responce struct {
	Ok     bool                     `json:"ok"`
	Result []map[string]interface{} `json:"result"`
}

// // using pointers for optional structs to check if value is nil in client/connect.go
// type Update struct {
// 	UpdateId           int  `json:"update_id"`
// 	Message            *any `json:"message"`
// 	EditedMessage      *any `json:"edited_message"`
// 	ChannelPost        *any `json:"channel_post"`
// 	EditedChannelPost  *any `json:"edited_channel_post"`
// 	InlineQuery        *any `json:"inline_query"`
// 	ChosenInlineResult *any `json:"chosen_inline_result"`
// 	CallbackQuery      *any `json:"callback_query"`
// 	ShippingQuery      *any `json:"shipping_query"`
// 	PreCheckoutQuery   *any `json:"pre_checkout_query"`
// 	PollStopped        *any `json:"poll"`
// 	PollAnswer         *any `json:"poll_answer"`
// 	MyChatMember       *any `json:"my_chat_member"`
// 	ChatMember         *any `json:"chat_member"`
// 	ChatJoinRequest    *any `json:"chat_join_request"`
// }
