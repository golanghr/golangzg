package main

import "fmt"

type myEnum string

const (
    MY_ENUM_FOO myEnum = "foo"
    MY_ENUM_BAR myEnum = "bar"
)

func main() {
    a := MY_ENUM_FOO
    s := "foo"

    // cannot use s (variable of type string) as myEnum value in assignment
    // a = s
    a = myEnum(s)
    fmt.Printf("a is %T with value %v\n", a, a)
}
