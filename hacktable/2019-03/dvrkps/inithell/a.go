package main

import "github.com/golanghr/golangzg/hacktable/2019-03/dvrkps/inithell/aaa"

var aa int

func init() {
	println("main/a: init aa")
	aa = aaa.A()
}

var az int

func init() {
	println("main/a: init az")
	az = aaa.Z()
}
