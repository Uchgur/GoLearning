package main

import "fmt"

func Ex46() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	sl = append(sl[0:3], sl[6:]...)

	fmt.Println(sl)
}
