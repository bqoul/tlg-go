package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/bqoul/tlg-go/upd"
)

func (bot Bot) Connect() {
	offset := 0

	for {
		resp, _ := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%v", bot.token, offset))

		body, _ := io.ReadAll(resp.Body)

		var data upd.Responce
		json.Unmarshal(body, &data)

		for _, update := range data.Result {
			values := reflect.ValueOf(update)
			types := values.Type()

			for i := 0; i < values.NumField(); i++ {
				if types.Field(i).Name == "UpdateId" {
					// if offset is not equals to update_id
					if offset != values.Field(i).Interface().(int) {
						// giving offset same value as update_id
						offset = values.Field(i).Interface().(int)
					} else {
						// othervise incrementing offset by one
						// to make sure that bot wont endlessly respond to he same update
						offset++
					}
					continue
				}

				// checking for struct fileds that have information inside
				if values.Field(i).Interface().(*any) != nil {
					// emitting event with the same name as the struct field
					bot.emit(types.Field(i).Name)
				}
			}
		}

		resp.Body.Close()
	}
}
