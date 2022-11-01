package logger

import "fmt"

type Logger interface {
	Error(args ...interface{})
}

type DefaultLogger struct {
}

func (DefaultLogger) Error(args ...interface{}) {
	fmt.Println(args...)
}
