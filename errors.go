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

// Error is exported so you can easily check types, but should be not used
// directly. Call New instead.
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

// New creates an error of ExecutionError kind. A message is mandatory. Stack
// of function calls will be attached automatically.
func New(msg string, args ...interface{}) error {
	return &Error{
		kind:  ExecutionError,
		msg:   fmt.Sprintf(msg, args...),
		stack: []string{identifyCaller()},
	}
}

// NewOfKind creates an error of a predefined or custom kind, much like
// New.
func NewOfKind(kind ErrorKind, msg string, args ...interface{}) error {
	return &Error{
		kind:  kind,
		msg:   fmt.Sprintf(msg, args...),
		stack: []string{identifyCaller()},
	}
}

// Mask helps convert generic errors into *Error type. You can use it to
// pin-point breaking calls in the code - it automatically appends a call
// stack.
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

// Maskf does exactly what Mask does, but you can add a custom message.
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
