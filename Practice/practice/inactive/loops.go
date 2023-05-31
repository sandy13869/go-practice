// Slices for dynamic array
package main

import (
	"fmt"
)

func main() {
	// Traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Normal %v \n", i)
	}

	// Create an arrya of range 1 to n
	n := 6 // Number of elements

	_numbers := make([]int, 0) // Declare an empty slice
	_strings := make([]string, 0)
	i := 1
	for i <= n {
		_numbers = append(_numbers, i)                    // Append each number to the slice
		_strings = append(_strings, fmt.Sprintf("%d", i)) // Append each number as a string to the slice
		i++
	}

	fmt.Println(_numbers) // Print the slice
	fmt.Println(_strings) // Print the string slice

	// For loop with a single condition
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// For loop as a while loop
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}

	// For loop with range
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
