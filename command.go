package gourd

import (
	"github.com/andersfylling/disgord"
)

type Command struct {
	name        string
	description string
	help        string
	aliases     []string
	permissions []disgord.PermissionBit
	run         func(ctx CommandContext)
}

func NewCommand(aliases ...string) *Command {
	return &Command{name: aliases[0], aliases: aliases}
}

func (cmd *Command) Name() string {
	return cmd.name
}

func (cmd *Command) Description() string {
	return cmd.description
}

func (cmd *Command) Help() string {
	return cmd.help
}

func (cmd *Command) Aliases() []string {
	return cmd.aliases
}

func (cmd *Command) Permissions() []disgord.PermissionBit {
	return cmd.permissions
}

func (cmd *Command) SetDescription(desc string) *Command {
	cmd.description = desc
	return cmd
}

func (cmd *Command) SetHelp(help string) *Command {
	cmd.help = help
	return cmd
}

func (cmd *Command) SetPermissions(perms ...disgord.PermissionBit) *Command {
	cmd.permissions = perms
	return cmd
}

func (cmd *Command) SetOnAction(run func(ctx CommandContext)) *Command {
	cmd.run = run
	return cmd
}
