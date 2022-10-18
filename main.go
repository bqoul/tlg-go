// this file is temporary
// need it to easy test functionality
// while the framework is in development
package main

import (
	"fmt"

	"github.com/bqoul/tlg-go/client"
	"github.com/bqoul/tlg-go/evt"
)

const token = "5681924286:AAHR49s5CMYFarcwjECeWfOWJFChS2iF-m4"

func main() {
	bot, _ := client.NewBot(token)

	bot.On(evt.Message, func() {
		fmt.Println("message")
	})

	bot.On(evt.EditedMessage, func() {
		fmt.Println("edited message")
	})

	bot.Connect()
}
