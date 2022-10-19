package client

import "github.com/bqoul/tlg-go/evt"

func (bot *Bot) emit(event string) {
	if actions, ok := bot.actions[event]; ok {
		for _, action := range actions {
			go action()
		}
	}
}

func (bot *Bot) On(event evt.Event, action func()) {
	if actions, ok := bot.actions[string(event)]; ok {
		bot.actions[string(event)] = append(actions, action)
	} else {
		bot.actions[string(event)] = []func(){action}
	}
}
