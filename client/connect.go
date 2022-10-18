package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Update struct {
	Ok     bool                     `json:"ok"`
	Result []map[string]interface{} `json:"result"`
}

func (bot Bot) Connect() {
	for {
		resp, _ := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", bot.token))

		body, _ := io.ReadAll(resp.Body)

		var data Update
		json.Unmarshal(body, &data)

		for _, update := range data.Result {
			for key := range update {
				if key == "update_id" {
					continue
				}

				bot.emit(key)
			}
		}

		resp.Body.Close()
	}
}
