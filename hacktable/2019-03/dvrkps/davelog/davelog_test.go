// Package davelog implements a minimal logging package.
// Inspired by article of Dave Cheney.
// https://dave.cheney.net/2015/11/05/lets-talk-about-logging
package davelog

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	l := New(nil, false)
	if l == nil {
		t.Error("New( ... ) == nil")
	}
	l.Infof("write to nil log")
}

func testSetup(verbose bool) (*Log, *bytes.Buffer) {
	var buf = &bytes.Buffer{}
	return New(buf, verbose), buf
}

func TestLog(t *testing.T) {
	tests := []struct {
		name    string
		verbose bool
		action  func(log *Log)
		want    string
	}{
		{
			name:    "infof",
			verbose: false,
			action: func(log *Log) {
				log.Infof("info %v", 42)
			},
			want: "info 42\n",
		},
		{
			name:    "debugf verbose",
			verbose: true,
			action: func(log *Log) {
				log.Debugf("debug %v", 42)
			},
			want: "debug 42\n",
		},
		{
			name:    "debugf",
			verbose: false,
			action: func(log *Log) {
				log.Debugf("debug %v", 42)
			},
			want: "",
		},
	}
	for _, tt := range tests {
		log, buf := testSetup(tt.verbose)
		tt.action(log)
		got := buf.String()
		if got != tt.want {
			t.Errorf("%s: = %q; want %q",
				tt.name, got, tt.want)
		}
	}
}
