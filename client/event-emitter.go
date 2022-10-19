package client

func (bot *Bot) emit(event string) {
	if actions, ok := bot.actions[event]; ok {
		for _, action := range actions {
			go action()
		}
	}
}

func (bot *Bot) On(event string, action func()) {
	if actions, ok := bot.actions[event]; ok {
		bot.actions[event] = append(actions, action)
	} else {
		bot.actions[event] = []func(){action}
	}
}
