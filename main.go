package main

import (
	"fmt"
	"os"
	"path/filepath"

	fpath "example.com/me/filepath"
)

func eprintf(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
}

var cmdName = filepath.Base(os.Args[0])

var helpMes = fmt.Sprintf(`
options:
    %v help
    -%c recursive
    -%c interactive

examples:
    %v %v
    %v file0 file1
    %v -%c dir0 file0
`, helpOpFull, recOp, intOp, cmdName, helpOpFull, cmdName, cmdName, recOp)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("%v: help\n", helpOpFull)
		return
	}
	if args[0] == helpOpFull {
		fmt.Print(helpMes)
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
		case tp.IsDir():
			if !op.recursive {
				eprintf("%v: is a directory\n", arg)
				continue
			}
			eraseDir(arg, op.interactive, stdErr)
		case tp.IsExistingFile():
			err := eraseFile(arg, op.interactive)
			if err != nil {
				eprintf("error: %v\n", err)
			}
		case tp.IsNotExisting():
			eprintf("%v: not found\n", arg)
		default:
			eprintf("%v: unsupported file type\n", arg)
		}
	}
}
