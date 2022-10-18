package tlggo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bqoul/tlg-go/upd"
)

type Bot struct {
	token  string
	events map[string][]func()
}

func NewBot(token string) (Bot, error) {
	bot := Bot{token, make(map[string][]func())}
	_, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getMe", token))
	return bot, err
}

func (bot Bot) GetMe() upd.GetMe {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getMe", bot.token))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data upd.GetMe
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return data
}
