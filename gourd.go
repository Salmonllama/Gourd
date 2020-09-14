package gourd

import (
	"context"
	"github.com/andersfylling/disgord"
	"github.com/salmonllama/gourd/internal"
	"strings"
)

type Gourd struct {
	Client        *disgord.Client // TODO: Embedded type? BOOK
	DefaultPrefix string
	Handler       *Handler
	OwnerId       string
	Keywords      map[string]string
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
	valid, usedPrefix := usesPrefix(msg, bot.DefaultPrefix)
	if !valid {
		return
	}

	// Trim the prefix from the message
	trimmedMessage := trimPrefix(messageContent, usedPrefix)

	// Split the trimmed message into command and command args
	commandString, args := separateCommand(trimmedMessage)

	// Check that it's actually a command
	if len(msg.Content) == 1 {
		return
	}

	// Check if it's an existing command
	for _, cmd := range bot.Handler.Commands {
		for _, alias := range cmd.Aliases {
			if alias == strings.ToLower(commandString) {

				// Create the command's context
				ctx := CommandContext{
					Command:       cmd,
					Prefix:        usedPrefix,
					Args:          args,
					CommandString: commandString,
					Message:       msg,
					Client:        bot.Client,
					Gourd:         bot,
				}

				// Check for DM allowance
				if cmd.Private == false && ctx.IsPrivate() {
					ctx.Reply("This command cannot be used in a DM")
					return
				}

				// Handle argument parsing
				// Check if the supplied arguments match the command's required arguments
				// TODO: Complete argument parsing. It's not ready, and will be available in a future release.
				//isValidArgs, err := parseArgs(args, cmd)
				//if err != nil {
				//	Console.Err(err.Error())
				//	return
				//}

				// Check if the user has permission to pass inhibition
				if !hasPermission(&ctx) {
					// Prevent usage
					return
				}

				// Run the command
				cmd.Run(ctx)
			}
		}
	}
}

// usesPrefix checks to see if the message starts with a registered prefix and therefore will trigger the command
func usesPrefix(msg *disgord.Message, prefix string) (bool, string) {
	if strings.HasPrefix(msg.Content, prefix) {
		return true, prefix
	} else {
		return false, ""
	}
}

// trimPrefix returns the message without the prefix
func trimPrefix(message string, usedPrefix string) string {
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
func separateCommand(message string) (command string, args []string) {
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

func hasPermission(ctx *CommandContext) bool {
	inhibitor := ctx.Command.Inhibitor

	switch i := inhibitor.(type) {
	case NilInhibitor:
		return true
	case RoleInhibitor:
		if ctx.IsPrivate() {
			return false
		}

		r := i.handle(ctx.AuthorMember().Roles)
		if r == false && i.Response != nil {
			ctx.Reply(i.Response)
		}

		return r
	case PermissionInhibitor:
		if ctx.IsPrivate() {
			return false
		}

		guild, err := ctx.Guild()
		if err != nil {
			Console.Err(err)
		}

		userPerms, err := ctx.Client.GetMemberPermissions(context.Background(), guild.ID, ctx.Author().ID)
		if err != nil {
			Console.Err(err)
		}

		r := i.handle(userPerms)
		if r == false && i.Response != nil {
			ctx.Reply(i.Response)
		}

		return r
	case KeywordInhibitor:
		r := ctx.Gourd.HasKeyword(ctx.Author().ID.String(), i.Value)

		if r == false && i.Response != nil {
			ctx.Reply(i.Response)
		}

		return r
	case OwnerInhibitor:
		r := ctx.IsAuthorOwner()

		if r == false && i.Response != nil {
			ctx.Reply(i.Response)
		}

		return r
	default:
		return false
	}
}

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
	bot.Handler.AddModule(mdl)
	return bot
}

func (bot *Gourd) AddModules(modules ...*Module) *Gourd {
	bot.Handler.Modules = append(bot.Handler.Modules, modules...)

	return bot
}

// AddKeyword adds a keyword permission to the given user.
// This keyword is stored in runtime memory and used in the KeywordInhibitor
func (bot *Gourd) AddKeyword(userId string, keyword string) {
	bot.Keywords[userId] = keyword
}

// HasKeyword checks if the user has the given keyword.
// This is automatically checked by the KeywordInhibitor.
func (bot *Gourd) HasKeyword(userId string, keyword string) bool {
	return bot.Keywords[userId] == keyword
}

// Connect opens the connection to discord
func (bot *Gourd) Connect() error {
	err := bot.Client.StayConnectedUntilInterrupted(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// New creates a new instance of Gourd
func New(token string, ownerId string, defaultPrefix string) *Gourd {
	client := disgord.New(disgord.Config{
		BotToken: token,
	})

	handler := Handler{}

	gourd := &Gourd{
		Client:        client,
		DefaultPrefix: defaultPrefix,
		Handler:       &handler,
		OwnerId:       ownerId,
		Keywords:      make(map[string]string, 0),
	}

	client.On(disgord.EvtMessageCreate, gourd.processCommand)

	return gourd
}
