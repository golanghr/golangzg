package both

import "testing"

func TestVersion(t *testing.T) {
	const (
		major = 1
		minor = 2
		patch = 3
		want  = "v1.2.3"
	)
	got := Version(major, minor, patch)
	if got != want {
		t.Errorf("Version( %d, %d, %d ) == %q; want %q",
			major, minor, patch, got, want)
	}
}
