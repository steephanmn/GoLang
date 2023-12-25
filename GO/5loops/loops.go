package main

import "fmt"

func main() {
	//infinite loop
	i := 1
	for {
		fmt.Println(i)
		break
	}

	//loop till condition

	i = 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//counter based loop

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
