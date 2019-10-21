package main

import (
	"log"

	"github.com/golanghr/golangzg/hacktable/2019-10/mbaco/safety-with-types/internal/secure"
	"github.com/golanghr/golangzg/hacktable/2019-10/mbaco/safety-with-types/internal/secure2"
)

func main() {
	// algorithType := secure.algorithm("UNSAFE")
	// output: cannot refer to unexported name secure.algorithm
	algorithType := secure.TypeA
	s := secure.Compute(algorithType, "value")
	log.Println(s)

	// even better

	// s = secure2.Compute("", "value")
	// cannot use "" (type string) as type secure2.algorithm in argument to secure2.Compute
	algorithType2 := secure2.TypeB
	s = secure2.Compute(algorithType2, "value")
	log.Println(s)
}
