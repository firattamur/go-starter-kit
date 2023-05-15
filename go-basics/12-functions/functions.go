package main

import "fmt"

func hello() {
	fmt.Println("Hello, Go!")
}

func sum(a int, b int) int {
	return a + b
}

func sumAndDiff(a, b int) (int, int) {
	return a + b, a - b
}

func sumAndDiffv2(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b

	return sum, diff
}

func sumAndDiffv3(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b

	return
}

func multiSum(nums ...int) int {

	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {

	hello()
	fmt.Println(sum(1, 2))
	fmt.Println(sumAndDiff(1, 2))
	fmt.Println(sumAndDiffv2(1, 2))
	fmt.Println(sumAndDiffv3(1, 2))
	fmt.Println(multiSum(1, 2, 3, 4, 5))

}
