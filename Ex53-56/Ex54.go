package main

import "fmt"

func Ex54() {
	p1 := person1{
		firstName: "James",
		lastName:  "Bond",
		flavors:   []string{"Chocolate", "Banana"},
	}

	p2 := person1{
		firstName: "Jenny",
		lastName:  "Moneypenny",
		flavors:   []string{"Strawberry", "Chocolate"},
	}

	mp := map[string]person1{
		"Bond":       p1,
		"Moneypenny": p2,
	}

	for _, v := range mp {
		fmt.Println(v)
		for _, d := range v.flavors {
			fmt.Println(d)
		}
	}
}
