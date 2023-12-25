package main

import "fmt"

func main() {
	a, b := 10, 2
	fmt.Printf("%v devided by %v is %v", a, b, divide(a, b))

	a, b = 10, 0
	fmt.Printf("%v devided by %v is %v", a, b, divide(a, b))
}

func divide(a, b int) int {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()

	c := a / b
	return c
}
