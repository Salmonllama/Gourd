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
		bot.Connect() // Bot logs in here
	}

	return 0
}

func startup() *gourd.Gourd {
	// Create modules, perform pre-login startup stuff
	// This can include creating modules/commands (not recommended, see the Separate-Modules section)
	GeneralModule := &gourd.Module{
		Name:      "General",
		Inhibitor: gourd.NilInhibitor{}, // Permissions are set on a per-module basis, through inhibitors, see the wiki for more info
	}

	PingCommand := GeneralModule.NewCommand("ping") // Set Aliases here. Command.name also becomes the first alias

	PingCommand.SetOnAction(func(ctx gourd.CommandContext) { // Command logic is set here
		ctx.Reply("Pong!")
	})

	GeneralModule.AddCommand(PingCommand)

	// Creates a new instance of Gourd
	bot := gourd.New("login-token-here",
		"owner-id-here",
		"default-prefix-here",
	)

	bot.AddModule(GeneralModule)
	return bot
}
