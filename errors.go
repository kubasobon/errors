package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type ErrorKind string

const (
	// default
	ExecutionError ErrorKind = "ExecutionError"
	ConfigError    ErrorKind = "ConfigError"
	NotImplemented ErrorKind = "NotImplemented"
)

type Error struct {
	kind  ErrorKind
	msg   string
	stack []string
}

func (e *Error) Error() string {
	stack := strings.Join(e.stack, "\n")
	return fmt.Sprintf("%s: %s\n---\n%s", e.kind, e.msg, stack)
}

// ensure type Error fulfills error interface
var _ error = &Error{}

func New(msg string, args ...interface{}) error {
	return &Error{
		kind:  ExecutionError,
		msg:   fmt.Sprintf(msg, args...),
		stack: []string{identifyCaller()},
	}
}

func NewOfKind(kind ErrorKind, msg string, args ...interface{}) error {
	return &Error{
		kind:  kind,
		msg:   fmt.Sprintf(msg, args...),
		stack: []string{identifyCaller()},
	}
}

func Mask(e error) error {
	switch v := e.(type) {
	case *Error:
		stack := append(v.stack, identifyCaller())
		return &Error{
			kind:  v.kind,
			msg:   v.msg,
			stack: stack,
		}
	default:
		return &Error{
			kind:  ExecutionError,
			msg:   e.Error(),
			stack: []string{identifyCaller()},
		}
	}
}

func Maskf(e error, msg string, args ...interface{}) error {
	switch v := e.(type) {
	case *Error:
		stack := append(v.stack, identifyCaller())
		return &Error{
			kind:  v.kind,
			msg:   fmt.Sprintf(msg, args...) + ":" + v.msg,
			stack: stack,
		}
	default:
		return &Error{
			kind:  ExecutionError,
			msg:   fmt.Sprintf(msg, args...) + ":" + e.Error(),
			stack: []string{identifyCaller()},
		}
	}
}

func identifyCaller() string {
	_, filename, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d", filename, line)
}
