package client

import (
	"errors"
)

type Bot struct {
	token   string
	actions map[string][]func()
}

func NewBot(token string) (Bot, error) {
	bot := Bot{token, make(map[string][]func())}

	var err error
	if !bot.GetMe().Ok {
		err = errors.New("tlg-go [ERROR] >>> incorrect token specified")
	}

	return bot, err
}
