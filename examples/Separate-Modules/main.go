package main

import (
	"github.com/Salmonllama/Gourd"
	"github.com/Salmonllama/Gourd/examples/Separate-Modules/modules"
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
	// Create a new instance of Gourd
	bot := gourd.New("your-bot-token-here", "!")

	// Add any defined modules
	// Don't forget to add your commands to your modules!
	bot.AddModules(
		modules.ModuleGeneral(),
		modules.ModuleModeration(),
	)

	return bot
}
