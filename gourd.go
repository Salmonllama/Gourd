package gourd

import (
	"context"
	"github.com/Salmonllama/Gourd/internal"
	"github.com/andersfylling/disgord"
	"strings"
)

type Gourd struct { // TODO: Add support for server specific prefixes
	Client        *disgord.Client
	DefaultPrefix string
	handler       *Handler
}

// Takes a message and decides if it should be treated as a command or not
func (bot *Gourd) ProcessCommand(session disgord.Session, evt *disgord.MessageCreate) {
	msg := evt.Message
	messageContent := msg.Content

	// Ignore bot users
	if msg.Author.Bot {
		return
	}

	// Check validity and find the prefix that was used
	valid, usedPrefix := bot.usesPrefix(msg)
	if !valid {
		return
	}

	// Trim the prefix from the message
	trimmedMessage := bot.trimPrefix(messageContent, usedPrefix)

	// Split the trimmed message into command and command args
	commandUsed, args := bot.separateCommand(trimmedMessage)

	// Check that it's actually a command
	if len(msg.Content) == 1 {
		return
	}

	ctx := CommandContext{
		prefix:  usedPrefix,
		args:    args,
		command: commandUsed,
		message: msg,
		client:  bot.Client,
	}

	// Check if it's an existing command
	for _, cmd := range bot.handler.Commands {
		for _, alias := range cmd.aliases {
			if alias == strings.ToLower(commandUsed) {
				cmd.run(ctx)
			}
		}
	}
}

// usesPrefix checks to see if the message starts with a registered prefix and therefore will trigger the command
func (bot *Gourd) usesPrefix(msg *disgord.Message) (bool, string) {
	if strings.HasPrefix(msg.Content, bot.DefaultPrefix) { // TODO: else if server specific prefix
		return true, bot.DefaultPrefix
	} else {
		return false, ""
	}
}

// trimPrefix returns the message without the prefix
func (bot *Gourd) trimPrefix(message string, usedPrefix string) string {
	m := 0
	n := len(usedPrefix)
	for i := range message {
		if m >= n {
			return message[i:]
		}
		m++
	}
	return message[:n]
}

// separateCommand splits the message into a command and a slice of command arguments, if present
func (bot *Gourd) separateCommand(message string) (command string, args []string) {
	split := strings.Split(message, " ")
	command = split[0]
	args = removeSpaces(split[1:])
	return
}

func removeSpaces(slice []string) (ret []string) {
	for i, s := range slice {
		if s != "" {
			ret = append(ret, slice[i])
		}
	}
	return
}

func (bot *Gourd) AddModule(mdl *Module) *Gourd {
	bot.handler.AddModule(mdl)
	return bot
}

func (bot *Gourd) AddModules(modules ...*Module) *Gourd {
	bot.handler.Modules = append(bot.handler.Modules, modules...)

	return bot
}

// Connect opens the connection to discord
func (bot *Gourd) Connect() error {
	err := bot.Client.StayConnectedUntilInterrupted(context.Background())
	internal.Check(err)
	return nil
}

// New creates a new instance of FSBot
func New(token string, defaultPrefix string) *Gourd {
	client := disgord.New(disgord.Config{
		BotToken: token,
	})

	cmd := Handler{}

	gourd := &Gourd{
		Client:        client,
		DefaultPrefix: defaultPrefix,
		handler:       &cmd,
	}

	client.On(disgord.EvtMessageCreate, gourd.ProcessCommand)

	return gourd
}
