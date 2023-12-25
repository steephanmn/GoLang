package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s)

	s = []string{"coffee", "tea", "espresso"}
	fmt.Println(s)

	fmt.Println(s[1])

	s2 := s
	fmt.Println(s2)
	s2[2] = "juice"

	fmt.Println(s2, s)

	s = append(s, "cold coffee", "cold drinks")
	fmt.Println(s)

}

//the changes in slices will automatically changed but in array there is no automatic change
