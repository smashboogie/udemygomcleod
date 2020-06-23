package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// fmt.Println(m)

	m["job_odd"] = []string{`evilness`, `fat`, `henchmen`}
	fmt.Println(m)

	delete(m, `bond_james`)

	for k, v := range m {
		fmt.Println("This is the record for key", k)
		for i, v := range v {
			fmt.Println("\t", i, v)
		}
	}
}
