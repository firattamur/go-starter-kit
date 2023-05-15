package main

import "fmt"

const (
	a = iota + 2 // iota = 0 + 2 = 2
	b = iota + 2 // iota = 1 + 2 = 3
	c = iota + 3 // iota = 2 + 3 = 5
	d = iota + 4 // iota = 3 + 4 = 7
	e = iota + 7 // iota = 4 + 7 = 11
)

func main() {

	fmt.Println(a, b, c, d, e)

}
