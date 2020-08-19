package gourd

import (
	"context"
	"github.com/Salmonllama/Gourd/internal"
	"github.com/andersfylling/disgord"
	"strings"
)

type Gourd struct {
	client        *disgord.Client // TODO: Embedded type? BOOK
	defaultPrefix string
	handler       *Handler
	ownerId       string
	keywords      map[string]string
}

// Takes a message and decides if it should be treated as a command or not
func (bot *Gourd) processCommand(_ disgord.Session, evt *disgord.MessageCreate) {
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

	// Handle argument parsing

	// Check if it's an existing command
	for _, cmd := range bot.handler.Commands {
		for _, alias := range cmd.Aliases {
			if alias == strings.ToLower(commandString) {
				// Parse arguments

				// Create the command's context
				ctx := CommandContext{
					Command:       cmd,
					Prefix:        usedPrefix,
					Args:          args,
					CommandString: commandString,
					Message:       msg,
					Client:        bot.client,
					Gourd:         bot,
				}

				// Check if the user has permission to pass inhibition
				if !bot.hasPermission(ctx) {
					// Send the inhibitors no-no response, if not nil
					if ctx.Command.Inhibitor.(Inhibitor).Response != nil {
						_, err := ctx.Reply(ctx.Command.Inhibitor.(Inhibitor).Response)
						if err != nil {
							internal.PrintCheck(err)
						}
					}
				}

				// Run the command
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

// parseArgs parses through given arguments for validity
func parseArgs(args []string, command *Command) (bool, error) {
	requiredArgs := command.Arguments
	if len(requiredArgs) == 0 {
		return true, nil
	}

	if len(args) == 0 && len(requiredArgs) != 0 {
		return false, &ArgumentError{Arguments: requiredArgs}
	}

	for i, reqArg := range requiredArgs {
		if reqArg.IsOptional { // Ignore it if the reqArg is optional
			continue
		}

		if reqArg.UseRemainder {
			// Remainder must be string? break and return?
			return true, nil
		}

		if reqArg.IsQuoted {
			// handle quoted arguments, planned for a later update.
		}

		if !internal.IsSet(args, i) { // Might need to args[len(args) +1] to add as last element; may overwrite elements
			args[i] = reqArg.Default
		}

		// Parse the types now, I guess
		switch reqArg.Type {
		case TextArg:
			continue
		case NumericArg:
			if !internal.IsNumeric(args[i]) {
				return false, &ArgumentError{Arguments: requiredArgs}
			}
		case EmojiArg:
			// EmojiArg is not implemented yet, and kind of a low priority
		}
	}

	return true, nil
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
	inhibitor := ctx.Command.Inhibitor.(InhibitorHandler)

func registerListeners(client *disgord.Client, listeners ...*Listener) {
	for _, l := range listeners {
		if len(l.Middlewares) == 0 {
			client.On(l.Type, l.OnEvent)
		} else {
			client.On(l.Type, l.OnEvent)
		}
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

	client.On(disgord.EvtMessageCreate, gourd.processCommand)

	return gourd
}
