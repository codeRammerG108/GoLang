package main

import "fmt"

func main() {
	fmt.Println("Structs")

	Abhi := User{"Abhi", "abhi@gmai.com", true, 18}
	Abhi.NewMail()
	fmt.Printf("Abhi: %+v", Abhi)
	fmt.Printf("Abhishek name is %v\n", Abhi.Name)
	Abhi.GetStatus()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Print("Mail", u.Email)
}
