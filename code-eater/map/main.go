package main

import "fmt"

func main() {
	fmt.Println("demo for map")

	people := map[string]int{}
	people["height"] = 6
	people["weight"] = 65

	students := map[int]string{
		1: "Sam",
		2: "Sandy",
		3: "Sha",
	}
	students[9] = "Shandy"

	fmt.Println("The People", people)

	fmt.Printf("Students details %v \n", students)
	for key, value := range students {
		fmt.Println(key, value)
	}
}
