package main

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestRandomize(t *testing.T) {
	tests := []struct {
		id  string
		buf []byte
		l   int64
	}{
		{"nil", []byte{}, 0},
		{"64", make([]byte, 0, 64), 64},
		{"aiueo", []byte("aiueo")[:0], 5},
	}
	for _, tt := range tests {
		r := bytes.NewBuffer(tt.buf)
		err := randomize(r, tt.l)
		if err != nil {
			t.Fatalf("id:%v, err:%v", tt.id, err)
		}
		if int64(r.Len()) != tt.l {
			t.Fatalf("id:%v, len:%v", tt.id, r.Len())
		}
		t.Logf("id:%v, buf:%v", tt.id, hex.EncodeToString(r.Bytes()))
	}
}

func TestRandomizeFile(t *testing.T) {
	tests := []struct {
		p   string
		err bool
	}{
		{"./test/randomize1.txt", false},
		{"test/randomize2.txt", true},
		{"test/randomize3.txt", true},
		{"test/absent", true},
	}
	for _, tt := range tests {
		err := randomizeFile(tt.p)
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
