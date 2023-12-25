package main

import "fmt"

func main() {
	a := "hello world"
	b := &a

	fmt.Println(b)
	fmt.Println(*b)
	*b = "hiiiiiiii"
	fmt.Println(a)

}
