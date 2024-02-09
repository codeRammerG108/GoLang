package main

import "fmt"

func main() {
	fmt.Println("Maps")
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["TS"] = "Typescript"
	languages["PY"] = "Python"

	fmt.Println("Maps : ", languages)
	fmt.Println("JS : ", languages["JS"])

	delete(languages, "PY") // to delete PY

	for key, value := range languages {
		fmt.Println("For ", key, " value is : ", value)
	}
}
