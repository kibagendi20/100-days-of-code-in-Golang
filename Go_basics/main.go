package main

import (
	"fmt"
	_ "time"
)

func main() {

	const hotelName string = "Gopher Hotel"
	const longitude = 24.806078
	const latitude = -78.243027

	var occupancy int = 12
	totalAmount := 1000

	fmt.Println(hotelName)
	fmt.Println(longitude)
	fmt.Println(latitude)
	fmt.Println(occupancy)
	fmt.Println(totalAmount)

}
