package main

import "fmt"

func Delete() {
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	delete(an, "George")

	fmt.Println("--- accessing keys that don't exist")
	delete(an, "Georage")      // won't panic
	fmt.Println(an["Georgey"]) // won't panice
	fmt.Println("------------------------")
}
