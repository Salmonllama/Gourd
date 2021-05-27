package modules

import (
	"github.com/andersfylling/disgord"
	"github.com/salmonllama/gourd"
)

func Ping(ctx gourd.CommandContext) {
	ctx.Reply("Pong!")
}

func Info(ctx gourd.CommandContext) {
	self, err := ctx.Client.CurrentUser().Get()
	if err != nil {
		panic(err)
	}

	embed := disgord.Embed{
		Title:       "Bot Info",
		Description: "This embed shows info about the bot!",
		URL:         "https://github.com/Salmonllama/gourd",
		Timestamp:   disgord.Time{},
		Author: &disgord.EmbedAuthor{
			Name:    self.Username,
			IconURL: self.Avatar,
		},
	}

	ctx.Reply(embed)
}

var PingCommand = gourd.Command{
	Name:        "Ping",
	Description: "Pings the bot",
	Aliases:     []string{"ping"},
	Inhibitor:   gourd.NilInhibitor{},
	Arguments:   nil, // Not implemented yet
	Private:     true,
	Run:         Ping,
}

var InfoCommand = gourd.Command{
	Name:        "Info",
	Description: "Shows info about the bot",
	Aliases:     []string{"info"},
	Inhibitor:   gourd.NilInhibitor{},
	Arguments:   nil, // Not implemented yet
	Private:     true,
	Run:         Info,
}

var ModuleGeneral = &gourd.Module{
	Name:        "General",
	Description: "Holds all the general commands",
	Commands:    []*gourd.Command{&PingCommand, &InfoCommand},
	Listeners:   nil,
}
