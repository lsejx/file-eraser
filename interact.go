package main

import (
	"fmt"
	"os"
	"sync"
)

type interactT struct{ sync.Mutex }

var interacter *interactT

func (p *interactT) ask(path string) (bool, error) {
	p.Lock()
	defer p.Unlock()
	fmt.Printf("erase %v? (y/n) > ", path)
	var ans string
	for {
		_, err := fmt.Scan(&ans)
		if err != nil {
			return false, err
		}
		if ans == "y" || ans == "n" {
			break
		}
		buf := make([]byte, len(path)+8)
		for i := range buf {
			buf[i] = ' '
		}
		fmt.Print(string(buf))
		fmt.Print("(y/n) > ")
	}
	return ans == "y", nil
}

// stderr during interaction
type stdErrT struct{ *interactT }

var stdErr *stdErrT

func (p *stdErrT) Write(b []byte) (n int, err error) {
	p.Lock()
	defer p.Unlock()
	return os.Stderr.Write(b)
}

func init() {
	interacter = new(interactT)
	stdErr = new(stdErrT)
	stdErr.interactT = interacter
}
