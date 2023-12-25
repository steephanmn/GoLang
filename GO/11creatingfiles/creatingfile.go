package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("working with files in go")

	content := "this should be displayed inside file"

	file, err := os.Create("./newFile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is:", length)

}
