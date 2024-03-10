package main

import "fmt"

func AnonFunc() {
	foo3()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}("Todd")

	s := func(s string) string {
		fmt.Println("This function has a return")
		return fmt.Sprint("The book is ", s)
	}("Yoga for everyone")
	fmt.Println(s)
}

func foo3() {
	fmt.Println("Foo ran")
}
