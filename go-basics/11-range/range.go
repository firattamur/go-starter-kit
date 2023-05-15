package main

import "fmt"

func main() {

	var sentence string = "Go is awesome!"

	for i, c := range sentence {
		fmt.Println(i, string(c))
	}

	var m map[string]int = make(map[string]int)

	for i, c := range sentence {
		m[string(c)] = i
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
