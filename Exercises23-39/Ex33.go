package main

import "fmt"

func Ex33() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println()

	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}
