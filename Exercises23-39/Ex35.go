package main

import "fmt"

func Ex35() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("The number is %v on position %v\n", v, i)
	}
}
