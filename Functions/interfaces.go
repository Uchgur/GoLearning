package main

import "fmt"

type person2 struct {
	first string
}

type secretAgent struct {
	person2
	ltk bool
}

func (p person2) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func InterAndPol() {
	sa1 := secretAgent{
		person2: person2{
			first: "James",
		},
		ltk: true,
	}

	p2 := person2{
		first: "Jenny",
	}

	sa1.speak()
	p2.speak()

	saySomething(sa1)
	saySomething(p2)
}
