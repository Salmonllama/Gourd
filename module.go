package gourd

type Module struct {
	Name        string
	Description string
	Commands    []*Command
}

// Modules by default are initialized with a NilInhibitor. It can be overwritten.
func NewModule(name string) *Module { // Are these methods even necessary?
	return &Module{
		Name:     name,
		Commands: nil,
	}
}

func (module *Module) NewCommand(aliases ...string) *Command {
	return &Command{
		Name:    aliases[0],
		Aliases: aliases,
	}
}

func (module *Module) AddCommands(cmds ...*Command) *Module {
	module.Commands = append(module.Commands, cmds...)
	return module
}

func (module *Module) AddCommand(cmd *Command) *Module {
	module.Commands = append(module.Commands, cmd)

	return module
}
