package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading Data")

	fmt.Println("Enter the input")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("the entered input is ", input)
	fmt.Printf("type of entered input is %T", input)

}
