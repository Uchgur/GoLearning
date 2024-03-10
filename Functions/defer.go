package main

import "fmt"

func Defer() {
	defer foo2()
	bar2()
}

func foo2() {
	fmt.Println("foo")
}

func bar2() {
	fmt.Println("bar")
}
