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
		fmt.Println("got a message")
	})

	bot.On(evt.Message, func() {
		fmt.Println("test")
	})

	bot.Connect()
}
