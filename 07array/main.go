package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array!")

	var fruitList [4]string
	// var fruitList = [4]string{"Apple", "Peach"}

	fruitList[0] = "Apple"
	fruitList[1] = "Peach"

	fmt.Println("FruitList : ", fruitList)
	//len(fruitList) - Gives Length of List

}
