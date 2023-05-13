package main

import "sync"

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
	_, ok := fl.m[path]
	if ok {
		return false
	}
	fl.m[path] = 0
	return true
}
