package main

import "fmt"

func FuncExpression() {
	z := foo4

	x := func() {
		fmt.Println("Anonymous func ran")
	}

	y := func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}

	z()
	x()
	y("Todd")
}

func foo4() {
	fmt.Println("Foo ran")
}
