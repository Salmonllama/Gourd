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
	hour, min, sec := time.Now().Clock()
	fmt.Printf("%v:%v:%v | DEBUG - ", hour, min, sec)
	fmt.Println(m...)
}

func (ConsoleLogger) Info(m ...interface{}) {
	hour, min, sec := time.Now().Clock()
	fmt.Printf("%v:%v:%v | INFO - ", hour, min, sec)
	fmt.Println(m...)
}

func (ConsoleLogger) Err(m ...interface{}) {
	hour, min, sec := time.Now().Clock()
	fmt.Printf("%v:%v:%v | Error - ", hour, min, sec)
	fmt.Println(m...)
}
