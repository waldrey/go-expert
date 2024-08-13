package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventDispatcherInterface) bool
	Dispatch(event EventInterface) error
	Clear() error
}
