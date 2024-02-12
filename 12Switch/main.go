package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("The diceNumber is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 you can move 1")
	case 2:
		fmt.Println("Dice value is 2 you can move 2")
	case 3:
		fmt.Println("Dice value is 3 you can move 3")
	case 4:
		fmt.Println("Dice value is 4 you can move 4")
	case 5:
		fmt.Println("Dice value is 5 you can move 5")
		fallthrough
	case 6:
		fmt.Println("Dice value is 6 you can move 6 or add one more piece in gaming area with aditional one extra chance")
	default:
		fmt.Println("What was this!")

	}

	//Break is included and need to be added in code but in some case you might need to run next case also you can use fallthrough
	//Example : if you run 5 by default it wont run 6 but because fallthrough it will run 6.
}
