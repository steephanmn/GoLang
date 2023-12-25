package main

import "fmt"

func main() {

	userData := User{"stpn", "stpn@gmail.com", true, 20}

	fmt.Println(userData)

	userData.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("user status", u.Status)

}
