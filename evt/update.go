// events for general updates we receive from telegram api
package evt

const (
	Message            string = "Message"
	EditedMessage      string = "EditedMessage"
	ChannelPost        string = "ChannelPost"
	EditedChannelPost  string = "EditedChannelPost"
	InlineQuery        string = "InlineQuery"
	ChosenInlineResult string = "ChosenInlineResult"
	CallbackQuery      string = "CallbackQuery"
	ShippingQuery      string = "ShippingQuery"
	PreCheckoutQuery   string = "PreCheckoutQuery"
	Poll               string = "Poll"
	PollAnswer         string = "PollAnswer"
	MyChatMember       string = "MyChatMember"
	ChatMember         string = "ChatMember"
	ChatJoinRequest    string = "ChatJoinRequest"
)
