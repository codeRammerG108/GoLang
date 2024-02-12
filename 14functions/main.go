package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")
	fmt.Println(adder(6, 8))
	fmt.Println(proAdder(6, 8, 5, 9, 1, 5))
}

//Anonymous Function
func proAdder(k ...int) (int, string) {
	total := 0

	for _, value := range k {
		total += value
	}
	return total, "Hi"
}

func adder(i int, j int) int {
	return i + j
}
