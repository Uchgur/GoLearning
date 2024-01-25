package main

import (
	"fmt"

	exerc "github.com/Uchgur/GoLearning/SecondLesson/tasks"
)

func main() {
	a := 42
	fmt.Printf("42 as binary, %b \n", a)      //This print a binary value
	fmt.Printf("42 as hexadecimal, %x \n", a) //This print a hex value

	exerc.Ex8()
	exerc.Ex9()
	exerc.Ex10()
	exerc.Ex11()
	exerc.Ex12()
	exerc.Ex13()
	exerc.TypeCon()
	exerc.ZeroVal()
}
