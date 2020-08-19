package gourd

type ArgType int

const (
	TextArg    ArgType = iota // 0
	NumericArg                // 1
	EmojiArg                  // 2
)

type Argument struct {
	Name         string
	Descriptor   string
	Type         ArgType
	IsOptional   bool
	UseRemainder bool
	IsQuoted     bool
	Default      string
}
