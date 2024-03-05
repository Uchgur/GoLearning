package main

import "fmt"

type person1 struct {
	firstName string
	lastName  string
	flavors   []string
}

func Ex53_1() {
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

	fmt.Println(p1)
	fmt.Println(p2)

	for _, v := range p1.flavors {
		fmt.Println(v)
	}

	for _, v := range p2.flavors {
		fmt.Println(v)
	}

}
