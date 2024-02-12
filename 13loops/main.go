package main

import "fmt"

func main() {
	fmt.Println("Loops in go Lang\n")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}

	fmt.Println(days)

	//type 1
	for d := 0; d < len(days); d++ {
		fmt.Println("Days ", days[d])
	}

	//type 2
	for i := range days { //range will return index in go unlike other language
		fmt.Println("Days ", days[i])
	}

	//type 3
	for index, day := range days {
		fmt.Println("Days ", day, index)
	}

	//While loop type - type 4
	roughNumber := 1
	for roughNumber < 10 {

		if roughNumber == 2 {
			goto lco //this is goto statement will jump to specific part of code
		}

		if roughNumber == 5 {
			break //Break := Break the loop while continue skip one phase of loops
		}
		fmt.Println("please add 1")
		roughNumber++
	}

lco:
	fmt.Println("Hello")
}
