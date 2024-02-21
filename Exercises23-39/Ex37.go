package main

import "fmt"

func Ex37() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	age := m["James"]
	fmt.Printf("Age of James is %v\n", age)

	if v, ok := m["Q"]; !ok {
		fmt.Printf("There is no Q in the map and the age is %v\n", v)
	} else {
		ageQ := m["Q"]
		fmt.Printf("Age of Q is %v\n", ageQ)
	}
}
