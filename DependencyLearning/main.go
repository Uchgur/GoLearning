package main

import (
	"fmt"

	"github.com/Uchgur/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()

	fmt.Println(s3)
}
