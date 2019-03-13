package zzz

var a int

func init() {
	println("zzz/a: init")
	a = 3
}

// A returns value.
func A() int {
	return a
}
