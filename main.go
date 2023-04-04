package main

import (
	"fmt"
	"os"

	fpath "github.com/lsejx/go-filepath"
)

func eprintf(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
}

var helpMsg = fmt.Sprintf(`file-eraser [option] [path] ...

options:
    %v help
    %c%c recursive
    %c%c interactive
`, helpOpFull, opPre, recOp, opPre, intOp)

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
			eraseDir(arg, op.interactive, stdErr)
		default:
			err := eraseFile(arg, op.interactive)
			if err != nil {
				eprintf("error: %v\n", err)
			}
		}
	}
}
