package main

import "github.com/golanghr/golangzg/hacktable/2019-03/dvrkps/inithell/zzz"

var za int

func init() {
	println("main/z: init za")
	za = zzz.A()
}

var zz int

func init() {
	println("main/z: init zz")
	zz = zzz.Z()
}
