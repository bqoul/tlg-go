package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bqoul/tlg-go/alias"
	"github.com/bqoul/tlg-go/upd"
)

func (bot Bot) Connect() {
	var offset int
	for {
		resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%v", bot.token, offset))
		alias.Check(err)

		body, err := io.ReadAll(resp.Body)
		alias.Check(err)
		resp.Body.Close()

		var data upd.Responce
		json.Unmarshal(body, &data)

		if len(data.Result) == 0 {
			continue // skipping iteration if no new updates received
		}

		for key, value := range data.Result[0] {
			// setting offset to only react on new updates
			if key == "update_id" {
				offset = int(value.(float64)) + 1
				continue
			}
			//emitting general events like message, channel post, callback query, etc
			bot.emit(key)

			//emitting additional events like text, location, dice, photo, etc
			for key = range value.(map[string]interface{}) {
				bot.emit(key)
			}
		}
	}
}
