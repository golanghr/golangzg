package main

import (
	"io"
	"os"

	"github.com/dvrkps/dojo/monorepo/internal/both"
)

func main() {
	os.Exit(run(os.Stdout, 1, 2, 3))
}

func run(w io.Writer, major, minor, patch int) int {
	v := both.Version(major, minor, patch)
	println("two:", v)
	return 0
}
