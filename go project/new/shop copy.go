package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

var menu = []menuItem{
	{name: "coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
	{name: "espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
}

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
			printMenu()
		case "2":
			aaddItem()
		case "q":
			break loop
		default:
			fmt.Println("unknown option")
		}
	}
}

func aaddItem() {
	fmt.Println("enter the name of the new item")
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})

}

func printMenu() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println("----------")
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}
