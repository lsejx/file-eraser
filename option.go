package main

import "strings"

const (
	opPre      = '-'
	helpOpFull = "-h"
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

func (p *option) read(arg string) (changed bool) {
	if len(arg) == 0 {
		return false
	}
	changed = false
	if arg[0] != opPre {
		return false
	}
	if strings.ContainsRune(arg, recOp) {
		p.recursive = true
		changed = true
	}
	if strings.ContainsRune(arg, intOp) {
		p.interactive = true
		changed = true
	}
	if strings.ContainsRune(arg, keepOp) {
		p.keep = true
		changed = true
	}
	return changed
}
