package main

import "fmt"

func Multidimensional() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)

	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
}
