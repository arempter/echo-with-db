package errors

import "github.com/sirupsen/logrus"

type Op string
type Message string
type Severity string
type UnknownErr Error

type Error struct {
	Op    Op
	Msg   Message
	Err   error
	Level logrus.Level
}

func (e Error) Error() string {
	return string(e.Msg)
}

func E(args ...interface{}) *Error {
	e := &Error{}

	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case Message:
			e.Msg = arg
		case error:
			e.Err = arg
		case logrus.Level:
			e.Level = arg
		default:
			e = &Error{}
		}
	}
	return e
}
