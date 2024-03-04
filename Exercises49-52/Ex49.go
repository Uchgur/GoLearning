package main

import "fmt"

func Ex49() {
	mp := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}

	for _, v := range mp {
		for i, d := range v {
			fmt.Printf("The index is %v and the value is %v\n", i, d)
		}
	}
}
