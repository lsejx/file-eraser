package main

import (
	"fmt"
	"os"

	fpath "example.com/me/filepath"
)

func eprintf(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
}

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		tp := fpath.GetType(arg)
		switch {
		case tp.IsDir():
			err := eraseDir(arg, os.Stderr)
			if err != nil {
				eprintf("%v\n", err)
				return
			}
		case tp.IsExistingFile():
			err := eraseFile(arg)
			if err != nil {
				eprintf("%v\n", err)
				return
			}
		case tp.IsNotExisting():
			eprintf("%v: not found\n", arg)
			return
		default:
			panic("unsupported file type")
		}
	}
}
