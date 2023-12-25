package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("web request")

	PerformGetRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code:", response.StatusCode)
	fmt.Println("content length:", response.ContentLength)

	//content, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(content))

	//instead of writing line 28 and 29 you can write the code below

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("byteCount is:", byteCount)
	fmt.Println(responseString.String())

}
