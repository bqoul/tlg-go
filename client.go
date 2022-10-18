package tlggo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bqoul/tlg-go/alias"
	"github.com/bqoul/tlg-go/upd"
)

type Bot struct {
	token  string
	events map[string][]func()
}

func NewBot(token string) (Bot, error) {
	bot := Bot{token, make(map[string][]func())}

	var err error
	if !bot.GetMe().Ok {
		err = errors.New("tlg-go [ERROR] >>> incorrect token specified")
	}

	return bot, err
}

func (bot Bot) GetMe() upd.GetMe {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getMe", bot.token))
	alias.Check(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	alias.Check(err)

	var data upd.GetMe
	err = json.Unmarshal(body, &data)
	alias.Check(err)
	return data
}
