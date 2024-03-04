package main

import "fmt"

func WordsInBook() {
	a := 0
	fmt.Println(a)
	a++
	fmt.Println(a)

	fmt.Println("-------")

	m := make(map[string]int)
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])
}
