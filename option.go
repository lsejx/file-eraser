package main

import "strings"

const (
	opPre      = '-'
	helpOpFull = "-h"
	verOpFull  = "-v"
	recOp      = 'r'
	intOp      = 'i'
	keepOp     = 'k'
)

type option struct {
	recursive   bool
	interactive bool
	keep        bool
}

func newOption() option {
	return option{false, false, false}
}

func (p *option) read(arg string) (isOption bool) {
	if len(arg) == 0 {
		return false
	}
	isOption = false
	if arg[0] != opPre {
		return false
	}
	if strings.ContainsRune(arg, recOp) {
		p.recursive = true
		isOption = true
	}
	if strings.ContainsRune(arg, intOp) {
		p.interactive = true
		isOption = true
	}
	if strings.ContainsRune(arg, keepOp) {
		p.keep = true
		isOption = true
	}
	return isOption
}
