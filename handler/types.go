package handler

type Command struct {
	Name string
	Description string
	Category string
	Aliases []string
	Usage string
	Permissions []string
}

type CommandContext struct {
	Prefix string
	Args []string
	Command string
}
