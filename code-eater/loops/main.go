package main

import "fmt"

func main() {
	fmt.Println("Loops Demo")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 5 {
			break // Exit the loop when i reaches 5
		}
	}

}
