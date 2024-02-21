package main

import (
	"fmt"
	"math/rand"
)

func Ex25() {
	x := rand.Intn(251)
	fmt.Printf("The value of %T is %v\n", x, x)
	switch {
	case x < 100:
		fmt.Println("Between 0 and 100")
	case (x > 100 && x < 200):
		fmt.Println("Between 101 and 200")
	case (x > 200 && x < 250):
		fmt.Println("Between 201 and 250")
	}
}
