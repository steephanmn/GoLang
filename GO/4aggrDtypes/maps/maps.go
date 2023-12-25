package main

import "fmt"

func main() {
	var m map[string]int
	m = map[string]int{"coffee": 1, "tea": 2}
	fmt.Println(m)

	fmt.Println(m["coffee"])

	m["juice"] = 3
	fmt.Println(m)

	delete(m, "juice")
	fmt.Println(m)

}
