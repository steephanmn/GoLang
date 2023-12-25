package main

import "fmt"

func main() {
	fmt.Println("structs")

	stpn := User{"stpn", "stpn@gmail.com", true, 20}

	fmt.Println("Details are", stpn)
	fmt.Printf("name= %v, Email=%v", stpn.Name, stpn.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
