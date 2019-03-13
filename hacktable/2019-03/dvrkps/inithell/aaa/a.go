package aaa

var a int

func init() {
	println("aaa/a: init")
	a = 1
}

// A returns value.
func A() int {
	return a
}
