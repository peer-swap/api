package reusable

import "reflect"

type Listener interface {
	Handle(interface{})
}

type Event interface {
	Emitter
	Listen(interface{}, Listener)
}

type Emitter interface {
	Emit(interface{})
}

type Name string

type EventMap struct {
	listeners map[Name][]Listener
}

func NewMapEvent() *EventMap {
	return &EventMap{listeners: map[Name][]Listener{}}
}

func (e EventMap) Listen(event interface{}, listener Listener) {
	name := e.getName(event)
	l := e.listeners[name]
	e.listeners[name] = append(l, listener)
}

func (e EventMap) Emit(event interface{}) {
	name := e.getName(event)

	for _, listener := range e.listeners[name] {
		listener.Handle(event)
	}
}

func (e EventMap) getName(i interface{}) Name {
	var name string

	if t := reflect.TypeOf(i); t.Kind() == reflect.Ptr {
		name = t.Elem().Name()
	} else {
		name = t.Name()
	}

	return Name(name)
}
