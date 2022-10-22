package client

func NewBot(token string) Bot {
	bot := Bot{token, make(map[string][]func())}

	// preventing user from passing the wrong token value when creating bot instance
	if !bot.GetMe().Ok {
		panic("incorrect token specified")
	}

	return bot
}
