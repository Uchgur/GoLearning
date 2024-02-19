package main

import "fmt"

func IfStatments() {
	x := 45
	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("equal to, or greater than, the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to the meaning of life")
	} else {
		fmt.Println("greater than the meaning of life")
	}
}
