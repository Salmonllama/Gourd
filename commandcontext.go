package gourd

import (
	"context"
	"github.com/andersfylling/disgord"
)

type CommandContext struct {
	prefix  string
	args    []string
	command string
	message *disgord.Message
	client  *disgord.Client
}

// Prefix returns the prefix used in the command
func (ctx *CommandContext) Prefix() string {
	return ctx.prefix
}

// Args returns the args included with the command, if any
func (ctx *CommandContext) Args() []string {
	return ctx.args
}

func (ctx *CommandContext) Command() string {
	return ctx.command
}

// Message returns the message object for the command message.
// It is not recommended to use this directly, use the convenience methods instead, if possible
func (ctx *CommandContext) Message() *disgord.Message {
	return ctx.message
}

// Client returns the disgord client
// It is not recommended to use this directly, use the convenience methods instead, if possible
func (ctx *CommandContext) Client() *disgord.Client {
	return ctx.client
}

// Reply sends a message to the channel the command was used in.
// Input is any type, see https://github.com/andersfylling/disgord/blob/39ba986ca2e94602ce44f4bf7625063124bdc325/client.go#L705
func (ctx *CommandContext) Reply(data ...interface{}) *disgord.Message {
	msg, err := ctx.message.Reply(context.Background(), ctx.client, data...)
	if err != nil {
		// TODO: error handling within gourd
	}

	return msg
}

func (ctx *CommandContext) IsPrivate() bool {
	return ctx.message.IsDirectMessage()
}

func (ctx *CommandContext) Author() *disgord.User {
	return ctx.message.Author
}

func (ctx *CommandContext) AuthorMember() *disgord.Member {
	return ctx.message.Member
}
