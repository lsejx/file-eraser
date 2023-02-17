package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var readLine = func() func() (string, error) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*64), 1024*1024)
	sc.Split(bufio.ScanLines)
	return func() (string, error) {
		ok := sc.Scan()
		if !ok {
			return "", sc.Err()
		}
		return sc.Text(), nil
	}
}()

type interactT struct{ sync.Mutex }

var interacter *interactT

func (p *interactT) ask(path string) (bool, error) {
	p.Lock()
	defer p.Unlock()
	var ans string
	for {
		fmt.Printf("erase %v? (y/n) > ", path)
		ans, err := readLine()
		if err != nil {
			return false, err
		}
		if ans == "y" || ans == "n" {
			break
		}
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
