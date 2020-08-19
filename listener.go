package gourd

import "github.com/andersfylling/disgord"

// Listener is the Gourd implementation of event handlers.
// Type is a required field, and it is suggested to use disgord supplied events (ex: disgord.EvtMessageCreate).
// OnEvent is the handler function, and takes an interface that should be cast appropriately (ex: event.(*disgord.MessageCreate) ).
type Listener struct {
	Type        string
	Middlewares []MiddlewareHandler
	OnEvent     func(session disgord.Session, event struct{})
}

type ListenerHandler interface { // Interface that all listeners must implement
	OnEvent(disgord.Session, interface{})
}

// Middleware is the Gourd representation of Disgord Middlewares
type Middleware struct{}

type MiddlewareHandler interface { // Interface that all middlewares must implement
	onProcess(evt interface{}) interface{}
}
