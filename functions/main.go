package main

import "fmt"

func main() {
	johnPrice := computePrice(20.00, 4)
	wesPrice := computePrice(40.0, 9)

	total := johnPrice + wesPrice

	fmt.Println(total)
}

func computePrice(rate float32, night int) float32 {
	return rate * float32(night)
}
