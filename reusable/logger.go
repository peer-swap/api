package reusable

import "log"

type Logger interface {
	Error(interface{})
}

type ConsoleLogger struct {
}

func (l ConsoleLogger) Error(v interface{}) {
	log.Println(v)
}
