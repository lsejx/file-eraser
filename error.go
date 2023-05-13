package main

import (
	"io/fs"
	"os"
	"sync"
)

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
