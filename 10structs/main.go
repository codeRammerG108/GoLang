package main

import "fmt"

func main() {
	fmt.Println("Structs")

	Abhi := User{"Abhi", "abhi@gmai.com", true, 18}
	fmt.Printf("Abhi: %+v", Abhi)
	fmt.Printf("Abhishek name is %v", Abhi.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
