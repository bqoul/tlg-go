package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bqoul/tlg-go/alias"
	"github.com/bqoul/tlg-go/ctx"
)

func (bot Bot) GetMe() ctx.GetMe {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getMe", bot.token))
	alias.Check(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	alias.Check(err)

	var data ctx.GetMe
	err = json.Unmarshal(body, &data)
	alias.Check(err)
	return data
}
