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

// concrete implementation of Error interface
type errorx struct {
	msg     string
	code    int
	err     error
	stack   Stack
	errType ErrType
}

func (e *errorx) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %s", e.msg, e.err.Error())
	}
	return e.msg
}

func (e *errorx) Format(s fmt.State, verb rune) {
	// print code and message
	fmt.Printf("%d: %s\n", e.code, e.Error())
}

func (e *errorx) Unwrap() error { return e.err }
func (e *errorx) Cause() error  { return e.err }
func (e *errorx) Code() int     { return e.code }
func (e *errorx) Type() ErrType { return e.errType }
func (e *errorx) Stack() Stack  { return e.stack }

func Wrap(err error) Error {
	if err == nil {
		return nil
	}
	return &errorx{
		msg:   err.Error(),
		err:   err,
		stack: captureStack(),
	}
}

func New(msg string) Error {
	return &errorx{
		msg:   msg,
		stack: captureStack(),
	}
}

func C(code int, msg string) Error {
	return &errorx{
		msg:   msg,
		code:  code,
		stack: captureStack(),
	}
}

func Cf(code int, format string, args ...interface{}) Error {
	return &errorx{
		msg:   fmt.Sprintf(format, args...),
		code:  code,
		stack: captureStack(),
	}
}

// Helper function to capture the stack trace
func captureStack() Stack {
	// TODO: Implement stack capture using runtime.Callers
	return Stack{}
}
