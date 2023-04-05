package main

import (
	"fmt"
	"os"
	"sync"

	fpath "github.com/lsejx/go-filepath"
)

func eprintf(format string, a ...any) {
	fmt.Fprintf(stdErr, format, a...)
}

var helpMsg = fmt.Sprintf(`file-eraser [option] [path] ...

options:
    %v help
    %c%c recursive (for directory)
    %c%c interactive (confirm before erasing)
    %c%c keep (randomize, but don't remove)
`, helpOpFull, opPre, recOp, opPre, intOp, opPre, keepOp)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("%v: help\n", helpOpFull)
		return
	}
	if args[0] == helpOpFull {
		fmt.Print(helpMsg)
		return
	}

	var wg sync.WaitGroup
	op := newOption()
	for _, arg := range args {
		argIsOp := op.read(arg)
		if argIsOp {
			continue
		}

		tp := fpath.GetType(arg)
		switch {
		case tp.IsNotExisting():
			eprintf("%v: not found\n", arg)
		case tp.IsDir():
			if !op.recursive {
				eprintf("%v: is a directory\n", arg)
				continue
			}
			wg.Add(1)
			go func(path string, op option) {
				defer wg.Done()
				eraseDir(path, op, stdErr)
			}(arg, op)
		default:
			wg.Add(1)
			go func(path string, op option) {
				defer wg.Done()
				if err := eraseFile(path, op); err != nil {
					eprintf("%v\n", err)
				}
			}(arg, op)
		}
	}
	wg.Wait()
}
