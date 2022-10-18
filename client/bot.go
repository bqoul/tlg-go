package client

type Bot struct {
	token   string
	actions map[string][]func()
}
