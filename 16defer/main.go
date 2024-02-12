package main

import "fmt"

func main() {
	fmt.Println("Hello Defer!\n")
	greatMe()
	greatUser()
}

func greatMe() {
	defer fmt.Print("Abhishek.\n")
	defer fmt.Print("is ")
	fmt.Print("Hello ")
	fmt.Print("My ")
}

func greatUser() {
	defer fmt.Print("your name.")
	defer fmt.Print("is ")
	fmt.Print("Hello ")
	fmt.Print("What ")
}

//Defer
// Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.

// In easy Language defer line will execute in end when the respective function is about to end and in form of filo
