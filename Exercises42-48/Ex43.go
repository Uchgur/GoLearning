package main

import "fmt"

func Ex43() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range sl {
		fmt.Printf("Value: %v, Type: %T\n", v, v)
	}
}
