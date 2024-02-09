package main

import "fmt"

func main() {
	fmt.Println("Pointer Section")

	var ptr *int
	number := 5
	*ptr = number

	// fmt.Println("Value of Pointer is ", ptr)            //Its store address So this print this
	fmt.Println("Value of Pointer's Address is ", *ptr) //Value of PTR
	// fmt.Println("Address of Pointer is ", &ptr) //This print the memory address of ptr

	//Okay

	// *ptr = *ptr + 1
	// fmt.Println("Value of Number is ", number) //Will the value of number changes?

}
