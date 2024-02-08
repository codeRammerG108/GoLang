package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey! Please Enter your phone number: ")
	reader := bufio.NewReader(os.Stdin)
	phoneNumber, _ := reader.ReadString('\n')
	fmt.Println("Your phone number : ", phoneNumber)
}
