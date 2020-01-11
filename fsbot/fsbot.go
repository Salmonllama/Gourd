package fsbot

import (
	"context"
	"github.com/andersfylling/disgord"
	"github.com/salmonllama/fsbot_go/lib"
)

type FSBot struct {
	Client *disgord.Client
	Config lib.Configuration
	// TODO: Add Database and Handler to the FSBot struct
}

func (bot *FSBot) isHomeGuild(id string) bool {
	return bot.Config.HomeGuild == id
}

// Connect opens the connection to discord
func (bot *FSBot) Connect() error {
	defer bot.Client.StayConnectedUntilInterrupted(context.Background())
	return nil
}

// New creates a new instance of FSBot
func New(config lib.Configuration) *FSBot {
	dgClient := disgord.New(disgord.Config{
		BotToken: config.Token,
	})

	fsbot := &FSBot{
		Client: dgClient,
		Config: config,
	}

	return fsbot
}
