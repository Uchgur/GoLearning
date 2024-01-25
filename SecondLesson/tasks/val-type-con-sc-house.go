package exerc

import "fmt"

func TypeCon() {
	y := 42
	z := 42.0
	fmt.Printf("%v of type %T\n", y, y)
	fmt.Printf("%v of type %T\n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		// this does not work!
		// in go you can't take a VALUE that is float32 and store it
		// in a variable that is declared to hold a VALUE of float64
		z = m
		fmt.Printf("%v of type %T \n", z, z)
	*/

	// this does work
	z = float64(m) //Here we will get some trash numbers in our extended part of a number
	fmt.Printf("%v of type %T \n", z, z)
	//So, this lines of code show, that if we want to conver value like this, we should do it explixitly
}
