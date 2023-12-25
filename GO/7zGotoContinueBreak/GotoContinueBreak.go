package main

import "fmt"

func main() {
	fmt.Println("goto break and continue in go")

	for num := 1; num < 10; num++ {
		fmt.Println(num)
		if num == 2 {
			goto gotolabel
		}
		if num == 5 {
			break
		} else if num == 3 {
			num++
			continue
		}
	}
gotolabel:
	fmt.Println("goto executed")
}
