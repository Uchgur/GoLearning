package main

import "fmt"

func Ex56() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "Ilya",
		friends:   map[string]int{"Nikita": 20, "Vanya": 19},
		favDrinks: []string{"Vse, chto gorit"},
	}

	fmt.Println(p1)
}
