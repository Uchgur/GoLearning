package main

import (
	"fmt"
	"math/rand"
)

func Ex24() {
	x := rand.Intn(251)
	fmt.Printf("The value of %T is %v\n", x, x)

	if x < 100 {
		fmt.Println("Between 0 and 100")
	} else if x > 100 && x < 200 {
		fmt.Println("Between 101 and 200")
	} else {
		fmt.Println("Between 201 and 250")
	}

}
