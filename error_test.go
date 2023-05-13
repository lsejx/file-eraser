package main

import (
	"errors"
	"testing"
)

func TestCatPathAndErr(t *testing.T) {
	tests := []struct {
		p   string
		s   string
		e   error
		ret string
	}{
		{"/root", "test-err", errors.New("test"), "/root: test-err, test"},
	}
	for _, tt := range tests {
		err := catPathAndErr(tt.p, tt.s, tt.e)
		if err == nil {
			t.Fatalf("path:%v, nil", tt.p)
		}
		if err.Error() != tt.ret {
			t.Fatalf("err:%v, e:%v", err, tt.ret)
		}
	}
}
