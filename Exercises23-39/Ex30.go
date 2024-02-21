package main

import (
	"fmt"
	"math/rand"
)

func Ex30() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)

		fmt.Printf("Iteration number %v\n", i)

		switch x {
		case 0:
			fmt.Printf("The x of type %T is zero\n", x)
		case 1:
			fmt.Printf("The x of type %T is one\n", x)
		case 2:
			fmt.Printf("The x of type %T is two\n", x)

		case 3:
			fmt.Printf("The x of type %T is three\n", x)

		case 4:
			fmt.Printf("The x of type %T is four\n", x)

		case 5:
			fmt.Printf("The x of type %T is five\n", x)
		}
	}
}
