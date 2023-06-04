package main

import "fmt"

func main() {
	fmt.Println("Demo for Array & slice")
	arr := [5]int{1, 5, 20, 40, 50}
	slice := []int{1, 5}
	_ = arr
	var str = "Sandeep Parimi"
	// arr[3] = 9
	fmt.Printf("Type of str is %v \n", str[:4])
	fmt.Println("Array elements :", arr)
	if len(slice) >= 2 {
		slice = append(slice, 100, 200)
		fmt.Println("Slice elements :", slice)
	} else {
		slice = append(slice, 200, 100)
		fmt.Println("Slice elements :", slice)
	}

	indextoDelete := 2
	copy(arr[indextoDelete:], arr[indextoDelete+1:])
	arr[len(arr)-1] = 0
	fmt.Println("Array final elements :", arr)

	slice = append(slice[:indextoDelete], slice[indextoDelete+1:]...)
	fmt.Println("Slice final elements :", slice)
	// if slice
	// for _, val := range arr {
	// 	fmt.Println("The Element:", val)
	// }
}
