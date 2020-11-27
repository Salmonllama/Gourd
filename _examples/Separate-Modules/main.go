package main

import (
	"os"

	"github.com/salmonllama/gourd"
	"github.com/salmonllama/gourd/_examples/Separate-Modules/modules"
)

func main() {
	os.Exit(lifecycle())
}

func lifecycle() int {
	bot := startup()
	if bot != nil {
		bot.Connect()
	}

	return 0
}

func startup() *gourd.Gourd {
	// Create a new instance of Gourd
	bot := gourd.New(
		"bot-token-here",
		"owner-id-here",
		"default-prefix-here",
	)

	// Add any defined modules
	// Don't forget to add your commands to your modules!
	bot.AddModules(
		modules.GeneralModule,
		modules.ModerationModule,
	)

	return bot
}
