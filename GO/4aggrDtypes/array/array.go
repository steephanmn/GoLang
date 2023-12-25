package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"coffee", "tea", "espresso"}
	fmt.Println(arr)

	fmt.Println(arr[1])

	arr2 := arr
	fmt.Println(arr2)
	arr2[2] = "juice"

	fmt.Println(arr2)

}
