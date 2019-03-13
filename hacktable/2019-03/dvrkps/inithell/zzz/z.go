package zzz

var z int

func init() {
	println("zzz/z: init")
	z = 4
}

// Z returns value.
func Z() int {
	return z
}
