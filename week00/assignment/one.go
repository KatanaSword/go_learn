package main

import "fmt"

func main() {
	numMessageFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessageFromDoris)

	fmt.Printf("Doris spent %.2f on text message today\n", totalCost)
}