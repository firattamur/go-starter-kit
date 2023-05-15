package main

import "fmt"

func main() {

	var a = make([]int, 5)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 6)
	fmt.Println(a)

	c := make([]int, 5)
	copy(c, a)
	fmt.Println(c)

	for i, v := range a {
		fmt.Println(i, v)
	}

}
