package main

import "fmt"

func main() {
	johnPrice := computePrices(20.00, 4)
	wesPrice := computePrices(160.0, 9)

	total := johnPrice + wesPrice

	fmt.Println(total)

}

// Func definition with 2 parameters with their datatype, and a named return type with its data type
func computePrices(rate float32, night int) (price float32) {
	price = rate * float32(night)
	return
}
