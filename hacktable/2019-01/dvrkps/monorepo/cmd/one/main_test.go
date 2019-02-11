package main

import (
	"io/ioutil"
	"testing"
)

func TestRun(t *testing.T) {
	got := run(ioutil.Discard, 1, 2, 3)
	if got != 0 {
		t.Errorf("run(...) == %v, want %v", got, 0)
	}
}
