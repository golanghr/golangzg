package show

import (
	"fmt"
	"io"
)

// Show prints formated text.
func Show(w io.Writer, format string, v ...interface{}) error {
	_, err := fmt.Fprintf(w, format, v...)
	return err
}
