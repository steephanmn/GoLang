package main

import "fmt"

func main() {
	a, b := 10, 2
	fmt.Printf("%v devided by %v is %v", a, b, divide(a, b))

}

func divide(a, b int) int {

	c := a / b
	return c
}
