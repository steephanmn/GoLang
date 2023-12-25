package main

import "fmt"

func main() {
	var s struct {
		name string
		id   int
	}
	s.name = "stpn"
	s.id = 1
	fmt.Println(s)
}
