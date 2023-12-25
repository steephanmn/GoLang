package main

import "fmt"

func main() {
	i := 5
	switch i {
	case 1:
		fmt.Println("one")
	case 5:
		fmt.Println("correct")
	default:
		fmt.Println("enter valid opt")

	}
}
