package both

import "fmt"

// Version returns command version.
func Version(major, minor, patch int) string {
	return fmt.Sprintf("v%d.%d.%d", major, minor, patch)
}
