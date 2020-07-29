package gourd

type Module struct {
	Name      string
	Inhibitor interface{}
	Commands  []*Command
}

// Modules by default are initialized with a NilInhibitor. It can be overwritten.
func NewModule(name string) *Module {
	return &Module{
		Name:      name,
		Commands:  nil,
		Inhibitor: NilInhibitor{},
	}
}

func (module *Module) NewCommand(aliases ...string) *Command {
	return &Command{
		Name:    aliases[0],
		Aliases: aliases,
		Module:  module,
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
