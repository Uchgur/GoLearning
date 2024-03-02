package main

import "fmt"

func Ex45() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	sl = append(sl, 52)

	fmt.Printf("%v\n", sl)

	sl = append(sl, 53, 54, 55)

	fmt.Printf("%v\n", sl)

	y := []int{56, 57, 58, 59, 60}

	sl = append(sl, y...)

	fmt.Printf("%v\n", sl)
}
