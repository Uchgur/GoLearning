package exerc

import "fmt"

const (
	c0 = iota + 1
	c1
	c2
	c3
	c4
	c5
)

func Ex8() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<c0, 1<<c0)
	fmt.Printf("%d \t %b\n", 1<<c1, 1<<c1)
	fmt.Printf("%d \t %b\n", 1<<c2, 1<<c2)
	fmt.Printf("%d \t %b\n", 1<<c3, 1<<c3)
	fmt.Printf("%d \t %b\n", 1<<c4, 1<<c4)
	fmt.Printf("%d \t %b\n", 1<<c5, 1<<c5)
}
