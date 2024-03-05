package main

import "fmt"

type secretAgent struct {
	person
	ltk   bool
	first string
}

func EmbeddedStructs() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "ETHAN HAWK",
		ltk:   true,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", sa1, sa1)

	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk, sa1.person)
	fmt.Println(sa1.first, sa1.person.first)
}
