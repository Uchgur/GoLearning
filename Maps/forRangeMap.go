package main

import "fmt"

func forRangeMap() {
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	// for range over a MAP
	for k, v := range an {
		fmt.Println(k, v)
	}

	for _, v := range an {
		fmt.Println(v)
	}

	for k := range an {
		fmt.Println(k)
	}
}
