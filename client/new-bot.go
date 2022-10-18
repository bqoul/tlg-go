package client

import "errors"

func NewBot(token string) (Bot, error) {
	bot := Bot{token, make(map[string][]func())}

	// preventing user from passing the wrong token value with creating bot instance
	var err error
	if !bot.GetMe().Ok {
		err = errors.New("tlg-go [ERROR] >>> incorrect token specified")
	}

	return bot, err
}
