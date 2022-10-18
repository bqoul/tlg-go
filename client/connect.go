package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/bqoul/tlg-go/alias"
	"github.com/bqoul/tlg-go/upd"
)

func (bot Bot) Connect() {
	offset := 1 // 1 here so bot wont respont to the last update twise

	for {
		resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%v", bot.token, offset))
		alias.Check(err)

		body, err := io.ReadAll(resp.Body)
		alias.Check(err)

		var data upd.Responce
		json.Unmarshal(body, &data)

		for _, update := range data.Result {
			values := reflect.ValueOf(update)
			types := values.Type()

			for i := 0; i < values.NumField(); i++ {
				if types.Field(i).Name == "UpdateId" {
					// if offset is not equals to update_id
					if offset != values.Field(i).Interface().(int) {
						// giving offset the value of update_id + 1
						offset += values.Field(i).Interface().(int)
					} else {
						// othervise incrementing offset by one
						// to make sure that bot wont endlessly respond to the same update
						offset++
					}
					continue // skiping "update_id" iteration because we dont need to emit it as event
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
