package main

import "fmt"

func Ex50() {
	mp := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}

	mp[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range mp {
		fmt.Printf("The key is %v ande the value is %v\n", k, v)
	}
}
