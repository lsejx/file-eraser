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
				os.Exit(1)
			}
		case tp.IsExistingFile():
			err := eraseFile(arg)
			if err != nil {
				eprintf("%v\n", err)
				os.Exit(1)
			}
		case tp.IsNotExisting():
			eprintf("%v: not found\n", arg)
			os.Exit(1)
		default:
			eprintf("unsupported file type")
			os.Exit(1)
		}
	}
}
