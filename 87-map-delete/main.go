//delete(<map name>, "key")
package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      32,
		"Miss Money": 27,
	}
	fmt.Println(m)

	//deleting from a map
	delete(m, "James")
	fmt.Println(m)
	//deleting a map value that doesnt exists
	delete(m, "Dr. No")
	fmt.Println(m)

	fmt.Println(m["Miss Money"])
	fmt.Println(m["Dr. No"])

	//ensure you deleted it - us an IF statment and then comma, ok
	if v, ok := m["Miss Money"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Money")
	}
	fmt.Println(m)

}
