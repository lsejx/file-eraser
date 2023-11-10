package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	fpath "github.com/lsejx/go-filepath"
)

const version = "v0.2.0"

var helpMsg = fmt.Sprintf(`%v [option] [path] ...

options:
    %v help
    %v version
    %c%c recursive (for directory)
    %c%c interactive (confirm before erasing)
    %c%c keep (randomize, seeks, truncates, but don't remove)
`, filepath.Base(os.Args[0]), helpOpFull, verOpFull, opPre, recOp, opPre, intOp, opPre, keepOp)

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
	if args[0] == verOpFull {
		fmt.Println(version)
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
			eprintf("%v: is not existing\n", arg)
		case tp.IsDir():
			// directory
			if !op.recursive {
				eprintf("%v: is a directory\n", arg)
				continue
			}
			wg.Add(1)
			go func(path string, op option) {
				defer wg.Done()
				if err := eraseDir(path, op, stdErr); err != nil && isRealErr(err) {
					eprintf("%v\n", err)
				}
			}(arg, op)
		case tp.IsRegularFile():
			// file
			wg.Add(1)
			go func(path string, op option) {
				defer wg.Done()
				if err := eraseFile(path, op); err != nil && isRealErr(err) {
					eprintf("%v\n", err)
				}
			}(arg, op)
		default:
			eprintf("%v: is neither a regular file nor a directory\n", arg)
		}
	}
	wg.Wait()
}
