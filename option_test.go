package main

import "testing"

func TestNewOption(t *testing.T) {
	o := newOption()
	if o.recursive != false {
		t.Fatal(o)
	}
	if o.interactive != false {
		t.Fatal(o)
	}
	if o.keep != false {
		t.Fatal(o)
	}
}

func TestOptionRead(t *testing.T) {
	tests := []struct {
		id   string
		ini  option
		arg  string
		ret  bool
		want option
	}{
		{"nil", option{false, false, false}, "", false, option{false, false, false}},
		{"false", option{false, false, false}, "a", false, option{false, false, false}},
		{"rec-change", option{false, false, false}, "-r", true, option{true, false, false}},
		{"int-change", option{false, false, false}, "-i", true, option{false, true, false}},
		{"kep-change", option{false, false, false}, "-k", true, option{false, false, true}},
		{"rec-true-false", option{true, false, false}, "-a", false, option{true, false, false}},
		{"rec-true-true", option{true, false, false}, "-r", true, option{true, false, false}},
		{"change-all", option{false, false, false}, "-rik", true, option{true, true, true}},
	}
	for _, tt := range tests {
		o := &option{tt.ini.recursive, tt.ini.interactive, tt.ini.keep}
		ret := o.read(tt.arg)
		if ret != tt.ret {
			t.Fatalf("id:%v, ret:%v", tt.id, ret)
		}
		if *o != tt.want {
			t.Fatalf("id:%v, o:%v", tt.id, *o)
		}
	}
}
