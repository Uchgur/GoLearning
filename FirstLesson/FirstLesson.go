package main

import "fmt"

func main() {
	//Base output syntax
	//fmt.Println("Hello world!😒")

	//Testing formating
	//const name, age = "Kim", 22
	//fmt.Printf("%s is %d years old and the type is %T and %T.\n", name, age, name, age)

	//Testing raw string literals
	// fmt.Println(`
	// UTF-8 saves space.
	// In UTF-8, common characters like “C” take 8 bits,
	// while rare characters like “💕” take 32 bits.
	// Other characters take 16 or 24 bits.
	// A blog post like this one takes about
	// four times less space in UTF-8
	// than it would in UTF-32.
	// So it loads four times faster.
	// `)

	//Exercise #1
	fmt.Println("Some text with some emojis 👍👌")
	fmt.Println(`Some
	raw
	string
	literal`)
}
