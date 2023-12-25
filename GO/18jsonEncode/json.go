package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //these are used to change the name which should be displayed
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("creating JSON file in go")
	EncodeJson()
}

func EncodeJson() {
	lcoCources := []course{
		{"reactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"webdev", "js"}},
		{"mern Bootcamp", 199, "learncodeonline.in", "bcd123", []string{"fullstack", "js"}},
		{"angular Bootcamp", 299, "learncodeonline.in", "hit123", nil},
	}

	//packaging this datas into json data

	finalJson, err := json.MarshalIndent(lcoCources, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(finalJson))
}
