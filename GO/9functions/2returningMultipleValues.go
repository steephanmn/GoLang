package main

import "fmt"

func main() {
	a, b := 10, 2
	result, ok := divide(a, b)
	if ok {
		fmt.Printf("%v devided by %v is %v", a, b, result)
	}

}

func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	c := a / b
	return c, true
}
