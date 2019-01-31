package show

import (
	"os"
	"testing"
)

func TestShow(t *testing.T) {
	err := Show(os.Stderr, "show: %v", "abc")
	if err != nil {
		t.Error(err)
	}
}
