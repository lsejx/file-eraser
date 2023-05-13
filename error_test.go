package main

import (
	"errors"
	"io/fs"
	"testing"
)

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
