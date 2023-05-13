package main

import (
	"testing"
)

func TestEraseFile(t *testing.T) {
	tests := []struct {
		p   string
		o   option
		err bool
	}{
		{"test/erase1.txt", option{false, false, false}, false},
		{"test/erase2.txt", option{false, false, true}, false},
		{"test/erase3.txt", option{false, false, false}, true},
	}
	for _, tt := range tests {
		err := eraseFile(tt.p, tt.o)
		if tt.err {
			if err == nil {
				t.Fatalf("path:%v, nilerr", tt.p)
			}
			t.Logf("path:%v, err:%v", tt.p, err)
		} else {
			if err != nil {
				t.Fatalf("path:%v, err:%v", tt.p, err)
			}
		}
	}
}

func TestEraseDir(t *testing.T) {
	tests := []struct {
		p string
		o option
		e bool
	}{
		{"test/eraseDir1", option{true, false, false}, false},
		{"test/eraseDir2", option{true, false, true}, false},
		{"test/eraseDir3", option{true, false, false}, true},
	}
	for _, tt := range tests {
		err := eraseDir(tt.p, tt.o, stdErr)
		if tt.e {
			if err == nil {
				t.Fatalf("path:%v, nilerr", tt.p)
			}
			t.Logf("path:%v, err:%v", tt.p, err)
		} else {
			if err != nil {
				t.Fatalf("path:%v, err:%v", tt.p, err)
			}
		}
	}
}
