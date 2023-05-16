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

func (p *interactT) ask(operation string, path string) (bool, error) {
	p.Lock()
	defer p.Unlock()
	var ans string
	for {
		fmt.Printf("%v %v? (y/n) > ", operation, path)
		var err error
		ans, err = readLine()
		if err != nil {
			return false, err
		}
		if ans == "y" || ans == "n" {
			break
		}
	}
	return ans == "y", nil
}

func init() {
	interacter = new(interactT)
}
