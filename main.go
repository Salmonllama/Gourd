package main

import (
	"context"
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/salmonllama/fsbot_go/lib"
)

var (
	config = lib.Config()
)

func printMessage(session disgord.Session, evt *disgord.MessageCreate) {
	msg := evt.Message
	fmt.Println(msg.Author.String() + ": " + msg.Content)
}

func main() {
	client := disgord.New(disgord.Config{
		BotToken: config.Token,
	})

	defer client.StayConnectedUntilInterrupted(context.Background())

	client.On(disgord.EvtMessageCreate, printMessage)
}