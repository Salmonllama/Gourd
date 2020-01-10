package lib

import (
	"github.com/andersfylling/disgord"
	"context"
)

type FSBot struct {
	Client *disgord.Client
	Config Configuration
	// Add Database
}

func (bot *FSBot) Connect() error {
	defer bot.Client.StayConnectedUntilInterrupted(context.Background())
	return nil
}

// New creates a new instance of FSBot
func New(config Configuration) *FSBot {
	dgClient := disgord.New(disgord.Config{
		BotToken: config.Token,
	})

	fsbot := &FSBot{
		Client: dgClient,
		Config: config,
	}

	return fsbot
}
