package main

import "fmt"

func Ex36() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("The key is %v and the value is %v\n", v, k)
	}
}
