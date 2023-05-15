package main

import "fmt"

func main() {

	var m map[string]int = make(map[string]int)
	m["foo"] = 42
	m["bar"] = 27

	fmt.Println(m)
	fmt.Println(len(m))

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "foo")
	fmt.Println(m)
	fmt.Println(len(m))

	fmt.Println(m["foo"])

	if v, ok := m["bar"]; ok {
		fmt.Println("bar,", v)
	}

}
