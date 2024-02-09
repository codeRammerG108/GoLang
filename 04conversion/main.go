package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome!")
	fmt.Println("Please enter your number : ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Printf("Hello, your number is type: %T, \n Let's Change its number", input)

	numConv, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nNow number is of type: %T", numConv)
		fmt.Println("\n", numConv)
	}

}
