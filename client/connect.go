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
	var offset int
	for {
		resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%v", bot.token, offset))
		alias.Check(err)

		body, err := io.ReadAll(resp.Body)
		alias.Check(err)

		var data upd.Responce
		json.Unmarshal(body, &data)

		if len(data.Result) == 0 {
			continue // skip iteration if no new updates received
		}

		values := reflect.ValueOf(data.Result[0])
		types := values.Type()

		for i := 0; i < values.NumField(); i++ {
			if types.Field(i).Name == "UpdateId" {
				if offset == values.Field(i).Interface().(int) {
					offset++
					continue
				}
				offset = values.Field(i).Interface().(int) + 1
				continue // skiping "update_id" iteration because we dont need to emit it as event
			}

			// checking for struct fields that have information inside
			if values.Field(i).Interface().(*any) != nil {
				// emitting general event with the same name as the struct field
				bot.emit(types.Field(i).Name)
			}
		}

		resp.Body.Close()
	}
}
