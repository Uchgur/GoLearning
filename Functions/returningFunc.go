package main

import "fmt"

func ReturningFunc() {
	x := foo5()
	fmt.Println(x)

	y := bar3()
	fmt.Println(y())

	fmt.Printf("%T\n", foo5)
	fmt.Printf("%T\n", bar3)
	fmt.Printf("%T\n", y)
}

func foo5() int {
	return 42
}

func bar3() func() int {
	return func() int {
		return 43
	}
}
