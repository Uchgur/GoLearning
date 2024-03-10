package main

import "fmt"

func Unfurling() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...) //This is what we call unfurling
	fmt.Println("The sum is", x)
}
