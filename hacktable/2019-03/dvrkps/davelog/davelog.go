package davelog

import (
	"io"
	"log"
)

// A Log represents an active logger.
type Log struct {
	output  *log.Logger
	verbose bool
}

// New creates logger.
func New(w io.Writer, verbose bool) *Log {
	l := Log{
		verbose: verbose,
	}
	if w != nil {
		l.output = log.New(w, "", 0)
	}
	return &l
}

// Infof logs formatted info message.
func (l *Log) Infof(format string, v ...interface{}) {
	l.logf(format, v...)
}

// Debugf logs formatted debug message.
func (l *Log) Debugf(format string, v ...interface{}) {
	if !l.verbose {
		return
	}
	l.logf(format, v...)
}

func (l *Log) logf(format string, v ...interface{}) {
	if l.output == nil {
		return
	}
	l.output.Printf(format, v...)
}
