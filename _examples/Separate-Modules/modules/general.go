package modules

import (
	"github.com/Salmonllama/Gourd"
)

var GeneralModule = &gourd.Module{
	Name:      "General",
	Inhibitor: gourd.NilInhibitor{}, // No inhibition required; anyone can use this command
	Commands: []*gourd.Command{
		ping(),
	},
}

func ping() (command *gourd.Command) {
	command = GeneralModule.NewCommand("ping")
	command.SetOnAction(func(ctx gourd.CommandContext) {
		ctx.Reply("Pong!")
	})

	return
}
