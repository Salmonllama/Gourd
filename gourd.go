package gourd

import (
	"context"
	"github.com/Salmonllama/Gourd/internal"
	"github.com/andersfylling/disgord"
	"strings"
)

type Gourd struct { // TODO: Add support for server specific prefixes
	client        *disgord.Client
	defaultPrefix string
	handler       *Handler
	ownerId       string
	keywords      map[string]string
}

// Takes a message and decides if it should be treated as a command or not
func (bot *Gourd) ProcessCommand(_ disgord.Session, evt *disgord.MessageCreate) {
	msg := evt.Message
	messageContent := msg.Content

	// Ignore bot users
	if msg.Author.Bot {
		return
	}

	// Blacklist checks go here

	// Check validity and find the prefix that was used
	valid, usedPrefix := bot.usesPrefix(msg)
	if !valid {
		return
	}

	// Trim the prefix from the message
	trimmedMessage := bot.trimPrefix(messageContent, usedPrefix)

	// Split the trimmed message into command and command args
	commandString, args := bot.separateCommand(trimmedMessage)

	// Check that it's actually a command
	if len(msg.Content) == 1 {
		return
	}

	ctx := CommandContext{
		Prefix:        usedPrefix,
		Args:          args,
		CommandString: commandString,
		Message:       msg,
		Client:        bot.client,
		Gourd:         bot,
	}

	// Check if it's an existing command
	for _, cmd := range bot.handler.Commands {
		for _, alias := range cmd.Aliases {
			if alias == strings.ToLower(commandString) {
				ctx.Command = cmd

				// Check for permission
				if !bot.hasPermission(ctx) {
					return
				}

				cmd.Run(ctx)
			}
		}
	}
}

// usesPrefix checks to see if the message starts with a registered prefix and therefore will trigger the command
func (bot *Gourd) usesPrefix(msg *disgord.Message) (bool, string) {
	if strings.HasPrefix(msg.Content, bot.defaultPrefix) { // TODO: else if server specific prefix
		return true, bot.defaultPrefix
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

func (bot *Gourd) hasPermission(ctx CommandContext) bool {
	inhibitorInterface := ctx.Command.Module.Inhibitor // TODO: Maybe move inhibitors to commands?

	switch inhibitorInterface.(type) {
	case NilInhibitor:
		return bot.handler.handleNilInhibitor()
	case RoleInhibitor:
		return bot.handler.handleRoleInhibitor(ctx)
	case PermissionInhibitor:
		return bot.handler.handlePermissionInhibitor(ctx)
	case KeywordInhibitor:
		return bot.handler.handleKeywordInhibitor(ctx)
	case OwnerInhibitor:
		return bot.handler.handleOwnerInhibitor(ctx)
	default:
		return false
	}
}

func (bot *Gourd) AddModule(mdl *Module) *Gourd {
	bot.handler.AddModule(mdl)
	return bot
}

func (bot *Gourd) AddModules(modules ...*Module) *Gourd {
	bot.handler.Modules = append(bot.handler.Modules, modules...)

	return bot
}
func (bot *Gourd) AddKeyword(userId string, keyword string) {
	bot.keywords[userId] = keyword
}

func (bot *Gourd) HasKeyword(userId string, keyword string) bool {
	return bot.keywords[userId] == keyword
}

// Connect opens the connection to discord
func (bot *Gourd) Connect() error {
	err := bot.client.StayConnectedUntilInterrupted(context.Background())
	internal.Check(err)
	return nil
}

// New creates a new instance of Gourd
func New(token string, ownerId string, defaultPrefix string) *Gourd {
	client := disgord.New(disgord.Config{
		BotToken: token,
	})

	handler := Handler{}

	gourd := &Gourd{
		client:        client,
		defaultPrefix: defaultPrefix,
		handler:       &handler,
		ownerId:       ownerId,
		keywords:      make(map[string]string, 0),
	}

	client.On(disgord.EvtMessageCreate, gourd.ProcessCommand)

	return gourd
}
