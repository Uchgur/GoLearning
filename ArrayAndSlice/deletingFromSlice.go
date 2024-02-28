package main

import "fmt"

func Deleting() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")

	xi = append(xi[:4], xi[5:]...) //There we delete fourth element
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")
}
