package main

import (
	"errors"
	"io/fs"
	"testing"
)

func TestIsRealErr(t *testing.T) {
	tests := []struct {
		arg error
		ret bool
	}{
		{nil, false},
		{errErrOccurred, false},
		{errCanceled, false},
		{errors.New("test"), true},
	}
	for _, tt := range tests {
		got := isRealErr(tt.arg)
		if got != tt.ret {
			t.Fatalf("a:%v, got:%v", tt.arg, got)
		}
	}
}

func TestParsePathErr(t *testing.T) {
	tests := []struct {
		err error
		ret string
	}{
		{fs.ErrPermission, "permission denied"},
		{errors.New("test"), "test"},
	}
	for _, tt := range tests {
		err := parsePathErr(tt.err)
		if err.Error() != tt.ret {
			t.Fatalf("err:%v, w:%v", err, tt.ret)
		}
	}
}
