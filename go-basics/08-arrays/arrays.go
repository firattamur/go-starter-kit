package main

import "fmt"

func main() {

	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(len(a))

	var b [2][2]int = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(b)
	fmt.Println(len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for row := 0; row < len(b); row++ {
		for col := 0; col < len(b[0]); col++ {
			fmt.Println(b[row][col])
		}
	}

}
