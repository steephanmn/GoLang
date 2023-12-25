package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"demo/menu"
)

var in = bufio.NewReader(os.Stdin)

func main() {

loop:
	for {
		fmt.Println("please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			menu.Aadd()
		case "q":
			break loop
		default:
			fmt.Println("unknown option")
		}
	}
}
