package gourd

import (
	"fmt"
	"time"
)

type Logger interface {
	Debug(m ...interface{})
	Info(m ...interface{})
	Err(m ...interface{})
}

type ConsoleLogger struct{}

var Console = ConsoleLogger{}

func (ConsoleLogger) Debug(m ...interface{}) {
	fmt.Printf("%v | DEBUG - ", time.Now().Clock())
	fmt.Println(m)
}

func (ConsoleLogger) Info(m ...interface{}) {
	fmt.Printf("%v | INFO - ", time.Now().Clock())
	fmt.Println(m)
}

func (ConsoleLogger) Err(m ...interface{}) {
	fmt.Printf("%v | Error - ", time.Now().Clock())
	fmt.Println(m)
}
