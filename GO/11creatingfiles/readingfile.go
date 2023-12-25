package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("reading the file from directory")
	readFile("./newFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("data inside the file:", string(databyte))
}
