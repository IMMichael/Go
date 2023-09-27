package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://microsoftedge.github.io/Demos/json-dummy-data/64KB-min.json"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("Response type: %T\n", resp) // The type will be a pointer named *http.Response
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes) //Converts the bytes received to a string using the string function

	//fmt.Print(content) //This will display the entire unformatted JSON

	persons := namesFromJson(content)

	for _, tour := range persons {
		fmt.Println(tour.Name)
	}

}
func namesFromJson(content string) []Tour {
	persons := make([]Tour, 0, 20) //Creates a slice from...???
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		persons = append(persons, tour)
	}
	return persons
}

type Tour struct {
	Name, Price string
}
