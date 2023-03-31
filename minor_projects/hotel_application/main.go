package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const totalRooms = 134
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println()
	fmt.Println()
	var hotel string = "Gopher Paris Inn"
	var roomsAvailable int
	var roomLabeling int = 110
	var occupancyLevel string
	roomsAvailable = rand.Intn(134)
	roomsOccupied := totalRooms - roomsAvailable
	occupancyRate := roomsOccupied / totalRooms * 100
	if occupancyRate > 60 {
		occupancyLevel = "High"
	} else if occupancyRate > 30 {
		occupancyLevel = "Medium"
	} else {
		occupancyLevel = "Low"
	}

	fmt.Print("Hotel: ", hotel)
	fmt.Println("      Occupancy level: ", occupancyLevel)
	fmt.Println("                             Occupancy rate: ", occupancyRate, "%")
	fmt.Println("Number of rooms: ", totalRooms)
	fmt.Println("Rooms available: ", roomsAvailable)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	if roomsAvailable != 0 {
		fmt.Println("Rooms:")
		for roomsAvailable != 0 {
			noPeople := rand.Intn(10) + 1
			nightsAvailable := rand.Intn(10) + 1
			if nightsAvailable == 1 {
				fmt.Printf("- %d : %d people / %d night \n", roomLabeling, noPeople, nightsAvailable)
			} else {
				fmt.Printf("- %d : %d people / %d nights \n", roomLabeling, noPeople, nightsAvailable)
			}

			roomLabeling++
			roomsAvailable--
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}
}
