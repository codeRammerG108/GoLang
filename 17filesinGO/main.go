package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("FIles in Go")

	content := "My name is Abhishek."

	file, err := os.Create("./demoFile.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Print(file)
		defer file.Close()
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	} else {
		fmt.Print(length)
	}

	readFile("./demoFile.txt")

}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nThe data is ", string(data))
}
