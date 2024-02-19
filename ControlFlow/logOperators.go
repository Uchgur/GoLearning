package main

import "fmt"

func LogOperators() {
	x := 40
	y := 5
	if x < 42 && y < 42 {
		// && requries both to be true to run
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		// || requires one of them to be true to run
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 && y != 10 {
		fmt.Println("x is not 42 & y is not 10")
	}

}
