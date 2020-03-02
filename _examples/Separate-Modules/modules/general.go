package modules

import (
	"github.com/Salmonllama/Gourd"
)

func ModuleGeneral() (module *gourd.Module) {
	module = gourd.NewModule("General", 0).AddCommands(ping())

	return module
}

func ping() (command *gourd.Command) {
	command = gourd.NewCommand("ping")
	command.SetOnAction(func(ctx gourd.CommandContext) {
		ctx.Reply("boop")
	})

	return
}
