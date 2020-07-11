package main

import "fmt"

func main() {
	mymap := make(map[string]int)

	mymap["刘炫"] = 20172519
	mymap["刘火"] = 20172117
	for i, v := range mymap {
		fmt.Println(i, v)
	}
}
