package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "MoneyPenny", "Hazelnut", "Bubblegum"}
	fmt.Println(mp)
	//multidimenstional slice below - a slice of a slice of string
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
