package gourd

import (
	"context"
	"github.com/andersfylling/disgord"
)

type CommandContext struct {
	Prefix        string
	Args          []string
	CommandString string
	Message       *disgord.Message
	Client        *disgord.Client
	Gourd         *Gourd
	Command       *Command
}

// Reply sends a message to the channel the command was used in.
// Input is any type, see https://github.com/andersfylling/disgord/blob/39ba986ca2e94602ce44f4bf7625063124bdc325/client.go#L705
func (ctx *CommandContext) Reply(data ...interface{}) (*disgord.Message, error) {
	msg, err := ctx.Message.Reply(context.Background(), ctx.Client, data...)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ctx *CommandContext) IsPrivate() bool {
	return ctx.Message.IsDirectMessage()
}

func (ctx *CommandContext) Guild() (*disgord.Guild, error) {
	guild, err := ctx.Client.GetGuild(context.Background(), ctx.Message.GuildID)
	if err != nil {
		return nil, err
	}

	return guild, nil
}

func (ctx *CommandContext) Author() *disgord.User {
	return ctx.Message.Author
}

func (ctx *CommandContext) AuthorMember() *disgord.Member {
	return ctx.Message.Member
}

func (ctx *CommandContext) IsAuthorOwner() bool {
	return ctx.Author().ID.String() == ctx.Gourd.ownerId
}
