package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello Sandy!")
	const owner = "Sandeep"
	name := "Sandy"
	roll := 18.5
	height := 5.10

	fmt.Printf("The name type is %T\n Rolle type %T\n height type %T\n", name, roll, height)

	getType(roll)
	compareType(roll)

}

func compareType(v interface{}) {
	if reflect.TypeOf(v).Kind() == reflect.Int {
		fmt.Println("Variable is of type Integer")
	} else {
		fmt.Println("Variable is not of type Integer")
	}
}

func getType(v interface{}) {
	res := fmt.Sprintf("%T", v)
	if res == "int" {
		fmt.Println("It's an Integer")
	} else {
		fmt.Println("It's not an Integer")
	}
}
