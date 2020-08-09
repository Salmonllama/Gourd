package gourd

type Handler struct {
	Commands []*Command
	Modules  []*Module
}

func (handler *Handler) addCommand(cmd *Command) *Command {
	handler.Commands = append(handler.Commands, cmd)
	return cmd
}

func (handler *Handler) GetCommandMap() map[string]*Command {
	cmdMap := make(map[string]*Command)

	for _, cmd := range handler.Commands {
		cmdMap[cmd.Name] = cmd
	}

	return cmdMap
}

func (handler *Handler) AddModule(mdl *Module) *Module {
	handler.Modules = append(handler.Modules, mdl)

	for _, cmd := range mdl.Commands {
		handler.addCommand(cmd)
	}

	return mdl
}

func (handler *Handler) GetModuleMap() map[string]*Module {
	mdlMap := make(map[string]*Module)

	for _, mdl := range handler.Modules {
		mdlMap[mdl.Name] = mdl
	}

	return mdlMap
}
