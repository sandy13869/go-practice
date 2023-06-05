package main

import "fmt"

func main() {
	fmt.Println("Hello Sandy!")

	name, age := "Sandy", 30
	rank := 3
	_, _ = age, rank
	fmt.Printf("Details \n Name: %v\n Rank: %T\n Age: %v\n", name, rank, age)
	name, age_ := "sam", 25
	_ = age_ // to silent unused variable in the code
	fmt.Printf("The New Name: %v \n", name)

	var (
		school string = name
		roll   int
		gender bool
	)

	fmt.Printf("The Group School: %v\n Roll: %d\n Gender is %v ", school, roll, gender)
}
