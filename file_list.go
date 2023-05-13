package main

import (
	"path/filepath"
	"sync"
)

type fileList struct {
	sync.Mutex
	m map[string]byte
}

func newFileList() *fileList {
	return &fileList{
		m: make(map[string]byte),
	}
}

// true: go (you are the first)
// false: stop (someone else is processing or processed it)
func (fl *fileList) isNew(path string) bool {
	fl.Lock()
	defer fl.Unlock()
	abs, err := filepath.Abs(path)
	if err != nil {
		// failed to get working dir
		abs = filepath.Clean(path)
	}
	_, ok := fl.m[abs]
	if ok {
		return false
	}
	fl.m[path] = 0
	return true
}
