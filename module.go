package gourd

type Module struct {
	Name     string
	commands []*Command
	Access   AccessType // 0 is all, 1 is none, 2 is home
}

type AccessType int

const (
	All AccessType = iota
	None
	Home
)

func NewModule(name string, access AccessType) *Module {
	return &Module{
		Name:     name,
		commands: nil,
		Access:   access,
	}
}

func (mdl *Module) AddCommands(cmd ...*Command) *Module {
	mdl.commands = append(mdl.commands, cmd...)
	return mdl
}

func (mdl *Module) AddCommand(cmd *Command) *Module {
	mdl.commands = append(mdl.commands, cmd)

	return mdl
}
