package main

import "fmt"

func main() {
	mymap := make(map[string]int)

	mymap["炫"] = 2017
	mymap["火"] = 2017
	for i, v := range mymap {
		fmt.Println(i, v)
	}
}
