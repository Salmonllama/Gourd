package main

import (
	"github.com/salmonllama/gourd"
	"os"
)

func main() {
	os.Exit(lifecycle())
}

func lifecycle() int {
	bot := startup()
	if bot != nil {
		err := bot.Connect() // Bot logs in here
		if err != nil {
			panic(err)
		}
	}

	return 0
}

func startup() *gourd.Gourd {
	// Create modules, perform pre-login startup stuff
	// This can include creating modules/commands (not recommended, see the Separate-Modules section)
	GeneralModule := &gourd.Module{
		Name:        "General",
		Description: "The general module",
	}

	PingCommand := gourd.Command{
		Name:        "Ping",
		Description: "Pings the bot",
		Aliases:     []string{"ping"},
		Inhibitor:   gourd.NilInhibitor{}, // Inhibitors are set for each command. See the wiki page for more info on Inhibitors
		Arguments:   nil,
		Private:     true,
		Run:         func(ctx gourd.CommandContext) { ctx.Reply("Pong!") },
	}

	GeneralModule.AddCommand(&PingCommand)

	// Creates a new instance of Gourd
	bot := gourd.New("login-token-here",
		"owner-id-here",
		"default-prefix-here",
	)

	bot.AddModule(GeneralModule)
	return bot
}
