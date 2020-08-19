package gourd

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Inhibitor   interface{}
	Arguments   []Argument // Can be nil -> no arguments required
	Private     bool
	Run         func(ctx CommandContext)
}
