package main

import "fmt"

func Ex44() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Printf("%v\n", sl[:5])
	fmt.Printf("%v\n", sl[5:])
	fmt.Printf("%v\n", sl[2:7])
	fmt.Printf("%v\n", sl[1:6])
}
