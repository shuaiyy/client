package main

import (
	"testing"

	"seelie/cmd/demo/www"
)

func TestCommandArgs(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "TestCommandArgs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CommandArgs()
		})
	}
}

func TestFS(t *testing.T) {

	x := www.FS
	//fs.ReadDir("images")
	//fs.ReadDir("fonts")

	tests := []struct {
		name string
	}{
		{name: "TestFS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, e := x.ReadDir("images")
			if e != nil {
				t.Errorf("TestFS() error = %v", e)
			}
			for _, vv := range v {
				t.Log(vv.Name())
			}
			v, e = x.ReadDir("fonts")
			if e != nil {
				t.Errorf("TestFS() error = %v", e)
			}
			for _, vv := range v {
				t.Log(vv.Name())
			}
			y, e := x.Open("index.html")
			if e != nil {
				t.Errorf("TestFS() error = %v", e)
			}
			t.Log(y.Stat())
		})
	}
}
