package main

import "fmt"

func main() {

	messagesFromDoris := []string{
		"Message 1",
		"Message 2",
		"Message 3",
		"Message 4",
	}

	numMessages := float64(len(messagesFromDoris))
	costPerMessage := .02

	totalCost := costPerMessage * numMessages

	fmt.Printf("Total messages sent by Doris %.f & Bill %.2f \n", numMessages, totalCost)
}
