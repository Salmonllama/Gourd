package modules

import (
	"strings"

	"github.com/andersfylling/disgord"

	"github.com/salmonllama/gourd"
)

var ModerationModule = &gourd.Module{
	Name: "Moderation",
	Inhibitor: gourd.PermissionInhibitor{
		Value:    disgord.PermissionManageServer,                               // Manage Server is required to use any commands in this module
		Response: "You need the Manage Server permission to use this command!", // Can also be any interface{} supported by disgord -> client.SendMsg()
	},
	Commands: []*gourd.Command{kick(), ban()},
}

func kick() (command *gourd.Command) {
	command = ModerationModule.NewCommand("kick", "boot", "bai") // Aliases are mandatory
	command.SetDescription("Kicks the mentioned users")          // Description is optional
	command.SetHelp("[]users, reason for the kick")              // Help is optional

	command.SetOnAction(func(ctx gourd.CommandContext) {
		if ctx.IsPrivate() {
			ctx.Reply("This command can only be used within a server!")
			return
		}

		users := ctx.Message.Mentions                      // Get any mentioned users
		reason := strings.Join(ctx.Args[len(users):], " ") // Reason starts after the user mentions

		kickBuilder := ctx.Client.Guild(ctx.Message.GuildID)
		for _, user := range users {
			if err := kickBuilder.Member(user.ID).Kick(reason); err != nil {
				// TODO: handle the error
			}
		}
	})

	return
}

func ban() (command *gourd.Command) {
	command = ModerationModule.NewCommand("ban", "banana", "hammertime")
	command.SetDescription("Bans the mentioned users")
	command.SetHelp("[]users, reason for the ban")

	command.SetOnAction(func(ctx gourd.CommandContext) {
		if ctx.IsPrivate() {
			ctx.Reply("This command can only be used within a server!")
			return
		}

		users := ctx.Message.Mentions                      // Get any mentioned users
		reason := strings.Join(ctx.Args[len(users):], " ") // Reason starts after the user mentions

		banBuilder := ctx.Client.Guild(ctx.Message.GuildID)
		for _, user := range users {
			banParams := &disgord.BanMemberParams{
				DeleteMessageDays: 30,
				Reason:            reason,
			}

			if err := banBuilder.Member(user.ID).Ban(banParams); err != nil {
				// TODO: handle the error
			}
		}
	})

	return
}
