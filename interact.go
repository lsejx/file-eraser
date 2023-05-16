package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

type interactT struct {
	sync.Mutex
	out io.Writer
	sc  *bufio.Scanner
}

func (p *interactT) readLine() (string, error) {
	ok := p.sc.Scan()
	if !ok {
		err := p.sc.Err()
		if err == nil {
			return "", errors.New("EOF")
		}
		return "", err
	}
	return p.sc.Text(), nil
}

func (p *interactT) ask(operation string, path string) (bool, error) {
	p.Lock()
	defer p.Unlock()
	var ans string
	for {
		fmt.Fprintf(p.out, "%v %v? (y/n) > ", operation, path)
		var err error
		ans, err = p.readLine()
		if err != nil {
			return false, err
		}
		if ans == "y" || ans == "n" {
			break
		}
	}
	return ans == "y", nil
}

func newLineScanner(in io.Reader) *bufio.Scanner {
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 1024*64), 1024*1024)
	sc.Split(bufio.ScanLines)
	return sc
}

var interacter = &interactT{
	out: os.Stdout,
	sc:  newLineScanner(os.Stdin),
}
