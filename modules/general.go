package modules

import (
	"context"
	"github.com/salmonllama/fsbot_go/handler"
)

func ModuleGeneral() *handler.Module {
	module := handler.InitModule("General", 0).AddCommands(ping(), test())
	return module
}

func ping() *handler.Command {
	PingCommand := handler.NewCommand("Ping", "ping")
	PingCommand.SetOnAction(func(ctx handler.CommandContext) {
		ctx.Message.Reply(context.Background(), ctx.Client, "boop")
	})

	return PingCommand
}

func test() *handler.Command {
	TestCommand := handler.NewCommand("Test", "test", "t")
	TestCommand.SetOnAction(func(ctx handler.CommandContext) {
		ctx.Message.Reply(context.Background(), ctx.Client, "test")
	})

	return TestCommand
}