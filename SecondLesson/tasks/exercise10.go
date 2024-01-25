package exerc

import "fmt"

func Ex10() {
	var a int
	var b int = 2
	c := 3
	var d, e = 4, 5

	_ = a
	fmt.Println(Sum(b, c, d, e))
}

func Sum(b, c, d, e int) int {
	sum := b + c + d + e
	return sum
}
