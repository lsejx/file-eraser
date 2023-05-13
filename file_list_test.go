package main

import (
	"sync"
	"testing"
)

func TestNewFileList(t *testing.T) {
	fl := newFileList()
	if fl.m == nil {
		t.FailNow()
	}
	if len(fl.m) != 0 {
		t.FailNow()
	}
}

func TestIsNew(t *testing.T) {
	fl := &fileList{m: make(map[string]byte)}
	if !fl.isNew("test") {
		t.Fatal(fl.m)
	}
	if fl.isNew("test") {
		t.Fatal(fl.m)
	}
	if len(fl.m) != 1 {
		t.Fatal(len(fl.m), fl.m)
	}
}

func TestFileList(t *testing.T) {
	fl := newFileList()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if fl.isNew("test") {
				t.Log("set", i)
			} else {
				t.Log(i)
			}
		}(i)
	}
	wg.Wait()
}
