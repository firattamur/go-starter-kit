package main

import "fmt"

func main() {

	var customerType string = "regular"

	switch customerType {

	case "regular":
		fmt.Println("You have earned 10 points!")

	case "premium":
		fmt.Println("You have earned 20 points!")

	case "vip":
		fmt.Println("You have earned 30 points!")

	default:
		fmt.Println("You have earned 0 points!")

	}

}
