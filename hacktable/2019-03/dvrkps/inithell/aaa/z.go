package aaa

var z int

func init() {
	println("aaa/z: init")
	z = 2
}

// Z returns value.
func Z() int {
	return z
}
