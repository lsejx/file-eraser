package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewLineScanner(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"\naiueo", ""},
		{"aiueo\nkakikukeko", "aiueo"},
	}
	for _, tt := range tests {
		sc := newLineScanner(strings.NewReader(tt.in))
		ok := sc.Scan()
		if !ok {
			t.Fatalf("in:%q, err:%v", tt.in, sc.Err())
		}
		if got := sc.Text(); got != tt.want {
			t.Fatalf("in:%q, got:%q, w:%q", tt.in, got, tt.want)
		}
	}
}

func TestReadLine(t *testing.T) {
	tests := []struct {
		in  string
		ret string
		isE bool
	}{
		{"", "", true},
		{"\n", "", false},
		{"aiueo\naiueo", "aiueo", false},
	}
	for _, tt := range tests {
		i := &interactT{
			out: os.Stderr,
			sc:  newLineScanner(strings.NewReader(tt.in)),
		}
		got, err := i.readLine()
		if tt.isE {
			if err == nil {
				t.Fatalf("in:%q, nilerr", tt.in)
			}
			t.Logf("id:%q, err:%v", tt.in, err)
		} else {
			if got != tt.ret {
				t.Fatalf("in:%q, got:%q, w:%q", tt.in, got, tt.ret)
			}
		}
	}
}

func TestAsk(t *testing.T) {
	tests := []struct {
		id  string
		o   string
		p   string
		in  string
		ret bool
		isE bool
	}{
		{"EOF", "", "", "", false, true},
		{"nil-yes", "", "", "y\n", true, false},
		{"nil-no", "", "", "n\n", false, false},
		{"non-nil", "testO", "testP", "y\n", true, false},
		{"loop-yes", "testO", "testP", "aaa\ny\n", true, false},
		{"loop-no", "testO", "testP", "bbbbbb\nn\n", false, false},
	}
	for _, tt := range tests {
		i := &interactT{
			out: os.Stderr,
			sc:  newLineScanner(strings.NewReader(tt.in)),
		}
		got, err := i.ask(tt.o, tt.p)
		if tt.isE {
			if err == nil {
				t.Fatalf("id:%v, nilerr", tt.id)
			}
			t.Logf("id:%v, err:%v", tt.id, err)
		} else {
			if got != tt.ret {
				t.Fatalf("id:%v, got:%v", tt.id, got)
			}
		}
	}
}
