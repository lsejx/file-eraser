package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"sync"
)

func eprintf(format string, a ...any) {
	fmt.Fprintf(stdErr, format, a...)
}

var (
	errErrOccurred = errors.New("error occured")
	errCanceled    = errors.New("canceled")
)

func isRealErr(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, errErrOccurred) || errors.Is(err, errCanceled) {
		return false
	}
	return true
}

func parsePathErr(err error) error {
	switch e := err.(type) {
	case *fs.PathError:
		return e.Err
	default:
		return e
	}
}

type stdErrT struct{ sync.Mutex }

var stdErr *stdErrT

func (p *stdErrT) Write(b []byte) (n int, err error) {
	p.Lock()
	defer p.Unlock()
	return os.Stderr.Write(b)
}

func init() {
	stdErr = new(stdErrT)
}
