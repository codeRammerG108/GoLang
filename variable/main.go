package main

import "fmt"

// JWTToken := "bzcjkdcdkj" // This is not allowed
// var JWTToken = "bzcjkdcdkj" // but this is allowed

func main() {
	var username string = "codeRammer108"
	fmt.Println("Username ::", username)
	fmt.Printf("Variable username is off type :: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println("Logged In ::", isLoggedIn)
	fmt.Printf("Variable username is off type :: %T\n", isLoggedIn)

	var intValue int64 = 255
	fmt.Println("Variable ::", intValue)
	fmt.Printf("Variable username is off type :: %T\n", intValue)

	var floatValue float64 = 255.956
	fmt.Println("Variable ::", floatValue)
	fmt.Printf("Variable username is off type :: %T\n", floatValue)

	var unIntializedInteger int64
	fmt.Println("Variable ::", unIntializedInteger)
	fmt.Printf("Variable username is off type :: %T\n", unIntializedInteger)

	//implicit Way
	var website = "codeRammer07.com"
	fmt.Println("Variable ::", website)
	fmt.Printf("Variable username is off type :: %T\n", website)
	// website = 5; // It will not allow you to re-intialize the variable

	url := "dummy.com"
	fmt.Printf("Variavble:: ", url)
	// url = 5; // It will not allow you to re-intialize the variable
}
