package main

// https://play.golang.org/p/UYS9aKfzRnL

func main() {
	println(Double(2))
}

func Double(n int) int {
	if n < 2 {
		return n
	}
	return multi(n)
}

func multi(n int) int {
	if n < 2 {
		return n
	}
	return n * n
}
