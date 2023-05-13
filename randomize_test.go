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
		p string
		i bool
	}{
		{"./test/randomize1.txt", false},
	}
	for _, tt := range tests {
		err := randomizeFile(tt.p, tt.i)
		if err != nil {
			t.Fatalf("path:%v, err:%v", tt.p, err)
		}
	}
}
