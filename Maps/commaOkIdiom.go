package main

import "fmt"

func CommaOkIdiom() {
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)

	delete(an, "George")

	// v, ok := an["Lucas"]
	// if ok {
	// 	fmt.Println("the value prints",v)
	// } else {
	// 	fmt.Println("Key didn't exist")
	// }

	if v, ok := an["Georgey"]; !ok {
		fmt.Println("Key didn't exist")
	} else {
		fmt.Println("the value prints", v)
	}

}
