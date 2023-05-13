package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	fpath "github.com/lsejx/go-filepath"
)

func eprintf(format string, a ...any) {
	fmt.Fprintf(stdErr, format, a...)
}

var helpMsg = fmt.Sprintf(`%v [option] [path] ...

options:
    %v help
    %c%c recursive (for directory)
    %c%c interactive (confirm before erasing)
    %c%c keep (randomize, but don't remove)
`, filepath.Base(os.Args[0]), helpOpFull, opPre, recOp, opPre, intOp, opPre, keepOp)

var fl = newFileList()

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

		if !fl.isNew(arg) {
			continue
		}

		tp := fpath.GetType(arg)
		switch {
		case tp.IsNotExisting():
			eprintf("%v: not found\n", arg)
		case tp.IsDir():
			// directory
			if !op.recursive {
				eprintf("%v: is a directory\n", arg)
				continue
			}
			wg.Add(1)
			go func(path string, op option) {
				defer wg.Done()
				if err := eraseDir(path, op, stdErr); err != nil {
					if !errors.Is(err, errErrOccurred) {
						eprintf("error: %v\n", err)
					}
				}
			}(arg, op)
		default:
			// file
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
