package main

import (
	"github.com/Salmonllama/Gourd"
	"os"
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
	// Create modules, perform pre-login startup stuff
	// This can include creating modules/commands (not recommended, see the Separate-Modules section
	GeneralModule := gourd.NewModule("General", 0)

	PingCommand := gourd.NewCommand("ping").SetOnAction(func(ctx gourd.CommandContext) {
		ctx.Reply("Pong!")
	})

	GeneralModule.AddCommand(PingCommand)

	// Creates a new instance of Gourd
	bot := gourd.New("your-bot-token-here", "!")
	bot.AddModule(GeneralModule)
	return bot
}
