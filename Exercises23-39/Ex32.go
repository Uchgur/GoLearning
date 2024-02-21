package main

import "fmt"

func Ex32() {
	for i := 0; ; {
		if i == 52 {
			fmt.Println("it's time to take a break")
			break
		}
		i++
	}
}
