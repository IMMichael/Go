package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Maps")

	states := make(map[string]string) //Create a map using the make function
	states["WA"] = "Washington"		  // assign values to each label(ie: "CA") and member (ie: "California")
	states["OR"] = "Oregon"
	states["CA"] = "California"

	fmt.Println(states)

	california := states["CA"] //Using the label from the map I assigned the contents of the member to that var
	fmt.Println(california)

	delete(states, "OR") //Delete a key/value using the delete function
	fmt.Println(states)

	states["NY"] = "New York"

	for k, v := range states { //Iterate through the keys and values of the map
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states)) //Create a slice with the contents of the map in order to use sort
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i:= range keys {
		fmt.Println(states[keys[i]])
	}



	
}