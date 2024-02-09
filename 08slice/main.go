package main

import "fmt"

func main() {
	fmt.Println("Slices!")

	var fruitList = []string{}
	//Another way can be declaration
	// fruitLists := make([]int, 4)
	fruitList = append(fruitList, "Mango", "Banana", "Peach", "Apple", "Carrot")

	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("FruitList is", fruitList)
	fmt.Println("Sliced FruitList is", fruitList[:])

}
