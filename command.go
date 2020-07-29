package gourd

type Command struct {
	Name        string
	Description string
	Help        string
	Aliases     []string
	Module      *Module // TODO: Remove this. Replace this with the Inhibitor. Remove Inhibitor from Modules
	Run         func(ctx CommandContext)
}
