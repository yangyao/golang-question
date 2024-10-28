package errorx

import "fmt"

type Error interface {
	error
	fmt.Formatter
	Unwrap() error
	Cause() error
	Code() int
	Type() ErrType
	Stack() Stack
}

type ErrType string

const (
	ErrTypeNotFound ErrType = "not_found"
	ErrTypeTimeout  ErrType = "timeout"
	// TODO: add more error types
)

type Frame struct {
	Name string
	File string
	Line int
}

type Stack []Frame

func Wrap(err error) Error {
	//TODO: implement
	return nil
}

func New(msg string) Error {
	//TODO: implement
	return nil
}

func C(code int, msg string) Error {
	//TODO: implement
	return nil
}

func Cf(code int, format string, args ...interface{}) Error {
	//TODO: implement
	return nil
}
