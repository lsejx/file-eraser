package main

import (
	"os"

	fpath "example.com/me/filepath"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		tp := fpath.GetType(arg)
		println(tp)
	}
}
