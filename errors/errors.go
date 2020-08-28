package errors

import "github.com/sirupsen/logrus"

type Op string
type Msg string
type Severity string

type Error struct {
	Op    Op
	Msg   Msg
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
		case Msg:
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
