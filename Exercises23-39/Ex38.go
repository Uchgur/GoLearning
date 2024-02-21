package main

import (
	"fmt"
	"math/rand"
)

func Ex38() {
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("The x is %v\n", x)
		}
	}
}
