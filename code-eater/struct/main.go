package main

import "fmt"

func main() {
	fmt.Println("Demo for structures")
	type Employee struct {
		name string
		id   uint
		dept string
	}

	myEmp := Employee{"Sam", 3456, "development"}
	fmt.Println("Employee details", myEmp)
	// fmt.Println("Employee address details", &myEmp)
	fmt.Println("Employee Name : ", myEmp.name)
}
