package main

import "fmt"

type engine struct {
	electric bool
}

type vechile struct {
	engine
	make  string
	model string
	doors int
	color string
}

func Ex55() {
	v1 := vechile{
		engine: engine{
			electric: true,
		},
		make:  "Opel",
		model: "Vectra",
		doors: 4,
		color: "Red",
	}

	v2 := vechile{
		engine: engine{
			electric: false,
		},
		make:  "Ford",
		model: "Mustang",
		doors: 4,
		color: "Pink",
	}

	fmt.Printf("Car is %v, engine is %v\n", v1, v1.engine)
	fmt.Printf("Car is %v, engine is %v\n", v2, v2.engine)
}
