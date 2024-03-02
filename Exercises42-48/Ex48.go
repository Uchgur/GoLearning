package main

import "fmt"

func Ex48() {
	sl := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "I'm 008"}}
	for _, v := range sl {
		for _, z := range v {
			fmt.Println(z)
		}
	}
}
