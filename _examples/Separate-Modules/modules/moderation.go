package modules

import (
	"github.com/andersfylling/disgord"
	"github.com/salmonllama/gourd"
	"strings"
)

func Kick(ctx gourd.CommandContext) {
	users := ctx.Message.Mentions
	reason := strings.Join(ctx.Args[len(users):], " ")

	for _, user := range users {
		err := ctx.Client.Guild(ctx.Message.GuildID).Member(user.ID).Kick(reason)
		if err != nil {
			panic(err)
		}
	}
}

func Ban(ctx gourd.CommandContext) {
	users := ctx.Message.Mentions
	reason := strings.Join(ctx.Args[len(users):], " ")

	for _, user := range users {
		params := disgord.BanMemberParams{Reason: reason}
		err := ctx.Client.Guild(ctx.Message.GuildID).Member(user.ID).Ban(&params)
		if err != nil {
			panic(err)
		}
	}
}

var KickCommand = gourd.Command{
	Name:        "Kick",
	Description: "Kicks the mentioned users from the guild. Does not work in DMs",
	Aliases:     []string{"kick", "boot"},
	Inhibitor: gourd.PermissionInhibitor{
		Value:    disgord.PermissionKickMembers,
		Response: "You require permission: Kick members.",
	},
	Arguments: nil, // Not yet implemented
	Private:   false,
	Run:       Kick,
}

var BanCommand = gourd.Command{
	Name:        "Ban",
	Description: "Bans the mentioned users from the guild. Does not work in DMs.",
	Aliases:     []string{"ban", "banana", "hammer"},
	Inhibitor: gourd.PermissionInhibitor{
		Value:    disgord.PermissionBanMembers,
		Response: "You require permission: Ban members.",
	},
	Arguments: nil, // Not yet implemented
	Private:   false,
	Run:       Ban,
}

var ModuleModeration = &gourd.Module{
	Name:        "Moderation",
	Description: "All moderation commands and functions",
	Commands:    []*gourd.Command{&KickCommand, &BanCommand},
}
