package gourd

import "fmt"

type ArgumentError struct {
	Arguments []Argument
}

func (argError *ArgumentError) Error() string {
	return fmt.Sprintf("Provided arguments were not valid. Proper arguments: %v", argError.Arguments)
}
