package main

import "fmt"

// Preferred Datatypes
// bool
// string
// int
// uint32
// byte
// rune
// float64
// complex128

func main() {
	const premiumPlan = "Premium Plan"
	const basicPlan = "Basic Plan"

	const secondsinMinute = 60
	const minutesinHour = 60
	const secondsinHour = secondsinMinute * minutesinHour

	const ourPlans = premiumPlan + "," + basicPlan

	fmt.Println("Their plans :", ourPlans)

	fmt.Println("Number of Seconds :", secondsinHour)

	fmt.Println("Sandy !")
}
