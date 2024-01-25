package exerc

import "fmt"

func ZeroVal() {
	//Two ways of creating variable. Second is more explicit and readable than the first one
	a := 42
	var b int = 31
	fmt.Println(a, " ", b)

	//Multiple creation is quite ok, but better not to use it, because it difficult to read and easy to make a mistake
	b, c, d, _, f := 0, 1, 2, 3, "happy" // This _ is a blank identifier. I can use it if I, fo example, defined some variable, by don't want to see it somewhere. So, I just "throw it intp space"
	fmt.Println(b, c, d, f)

	//g here will have a default value (for int it's 0)
	var g int
	fmt.Println(g)
}
