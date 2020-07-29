package gourd

type Command struct {
	Name        string
	Description string
	Inhibitor   interface{}
	Arguments   []Argument // Can be nil -> no arguments required
	Aliases     []string
	Run         func(ctx CommandContext)
}
