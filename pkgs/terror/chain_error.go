package terror

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type caller struct {
	file  string
	line  int
	frame runtime.Frame
}

func (c *caller) String() string {
	i := strings.LastIndex(c.frame.Function, ".")
	_, f := filepath.Split(c.file)
	return fmt.Sprintf("%s:%d:%s", f, c.line, c.frame.Function[i+1:])
}

type chainError struct {
	caller caller
	err    error
}

func (c *chainError) getTraceError(l int, newline, prefix string) string {

	if c.err != nil {
		var cerr *chainError
		if errors.As(c.err, &cerr) {
			return c.caller.String() + newline + strings.Repeat(prefix, l) + cerr.getTraceError(l+1, newline, prefix) // cerr.getTraceError(l+1, newline, prefix)
		}
		return c.caller.String() + " " + c.err.Error()
	}
	return c.caller.String()
}

func (c *chainError) Unwrap() error {
	return c.err
}

func (c *chainError) Error() string {
	return c.err.Error()
}

func Trace(err error) error {
	if err == nil {
		return nil
	}
	pc, f, l, ok := runtime.Caller(1)
	if !ok {
		return err
	}
	frames := runtime.CallersFrames([]uintptr{pc})
	frame, _ := frames.Next()
	chain, ok := err.(*chainError)
	if !ok {
		chain = &chainError{caller: caller{
			file:  f,
			line:  l,
			frame: frame,
		}, err: err}
		return chain
	}
	pChain := &chainError{caller: caller{
		file:  f,
		line:  l,
		frame: frame,
	}, err: chain}
	return pChain
}

// unwrap err until encounter not chain Error
func Unwrap(err error) error {

	for {
		if c, ok := err.(*chainError); ok {
			err = c.err
		} else {
			break
		}
	}
	return err
}

func TraceError(err error) string {
	return PrintIdent(err)
}

// print func call chain
func Print(err error) string {
	return printIdent(err, "", "")
}

// print pretty func call chain
func PrintIdent(err error) string {
	return printIdent(err, "\n", "\t")
}

func printIdent(err error, newline, prefix string) string {
	var cerr *chainError
	if errors.As(err, &cerr) {
		return cerr.getTraceError(1, newline, prefix)
	}
	return err.Error()
}
