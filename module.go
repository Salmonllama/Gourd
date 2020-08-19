package gourd

type Module struct {
	Name        string
	Description string
	Commands    []*Command
	Listeners   []*ListenerHandler
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

func (module *Module) AddListeners(l ...*ListenerHandler) *Module {
	module.Listeners = append(module.Listeners, l...)

	return module
}

func (module *Module) AddListener(l *ListenerHandler) *Module {
	module.Listeners = append(module.Listeners, l)

	return module
}
