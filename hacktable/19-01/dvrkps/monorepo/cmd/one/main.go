package main

import (
	"io"
	"os"

	"github.com/dvrkps/dojo/monorepo/internal/both"
	"github.com/dvrkps/dojo/monorepo/internal/show"
)

func main() {
	os.Exit(run(os.Stdout, 1, 2, 3))
}

func run(w io.Writer, major, minor, patch int) int {
	v := both.Version(major, minor, patch)
	err := show.Show(w, "run: %v", v)
	if err != nil {
		return 1
	}
	return 0
}
