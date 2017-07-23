package gucumber

import (
	"bytes"
	"fmt"
	"runtime"
	"github.com/fatih/color"
)

type Tester interface {
	Errorf(format string, args ...interface{})
	Skip(args ...interface{})
}

type TestingT struct {
	skipped bool
	errors  []TestError
}

type TestError struct {
	message string
	stack   []byte
}

func (t *TestingT) Errorf(format string, args ...interface{}) {
	var buf bytes.Buffer

	str := fmt.Sprintf(format, args...)
	sbuf := make([]byte, 8192)
	for {
		size := runtime.Stack(sbuf, false)
		if size < len(sbuf) {
			break
		}
		buf.Write(sbuf[0:size])
	}

	t.errors = append(t.errors, TestError{message: str, stack: buf.Bytes()})
}

func (t *TestingT) Skip(args ...interface{}) {
	t.skipped = true
}

func (t *TestingT) Skipped() bool {
	return t.skipped
}

func (t *TestingT) Failed() bool {
	return len(t.errors) > 0
}

func (t *TestingT) Error(err error) {
	t.errors = append(t.errors, TestError{message: err.Error()})
}
