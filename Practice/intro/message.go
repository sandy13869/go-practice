package main

import (
	"fmt"
)

func main() {
	learner := "Sandy"
	age := 30
	fmt.Printf("Welcome to Practice Area %v @ %v ..!\n", learner, age)

	type names []string
	friends := names{"sam", "sha"}
	fmt.Printf("My bae is %v \n", friends[1])

	nums := make([]int, 3)
	fmt.Println(nums)

	numbers := []int{1, 32, 3}
	for index, value := range numbers {
		fmt.Printf("Index : %v and value : %v \n", index, value)
	}
}
