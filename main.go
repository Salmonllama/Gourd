package main

import (
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/salmonllama/fsbot_go/fsbot"
	"github.com/salmonllama/fsbot_go/lib"
	"os"
)

var (
	config = lib.Config()
)

func printMessage(session disgord.Session, evt *disgord.MessageCreate) {
	msg := evt.Message
	fmt.Println(msg.Author.String() + ": " + msg.Content)
}

func main() {
	os.Exit(lifecycle())
}

func lifecycle() int {
	bot := fsbot.New(config)
	if bot != nil {
		err := bot.Connect()
		lib.Check(err)
	}
	return 0
}