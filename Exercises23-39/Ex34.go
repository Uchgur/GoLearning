package main

import "fmt"

func Ex34() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %v iteration of the first loop\n", i)
		fmt.Println("---")
		for j := 0; j < 5; j++ {
			fmt.Printf("This is the %v iteration of the second loop\n", j)
		}
	}
	fmt.Println()
}
