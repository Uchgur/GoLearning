package main

import "fmt"

func Ex42() {
	arr := [5]int{}
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4

	for i, v := range arr {
		fmt.Printf("Number %v have index %v\n", v, i)
	}
}
