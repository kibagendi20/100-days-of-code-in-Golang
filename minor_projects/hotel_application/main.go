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
	//var occupancyLevel string
	roomsOccupied := rand.Intn(134)
	roomsAvailable = totalRooms - roomsOccupied
	//occupancyRate := (float64(roomsOccupied) / float64(totalRooms)) * 100

	fmt.Print("Hotel: ", hotel)
	fmt.Println("      Occupancy level: ", computesOccupancyLevel(computesOccupancyRates(float32(roomsOccupied), float32(totalRooms))))
	fmt.Printf("                              Occupancy rate: %0.2f %%\n", computesOccupancyRates(float32(roomsOccupied), float32(totalRooms)))
	//fmt.Println("                             Occupancy rate: ", occupancyRate, "%")
	fmt.Println("Number of rooms: ", totalRooms)
	fmt.Println("Rooms available: ", roomsAvailable)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	if roomsAvailable != 0 {
		fmt.Println("Rooms:")
		for roomsAvailable != 0 {
			noPeople := rand.Intn(6) + 1
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

func computesOccupancyLevel(rates float32) string {
	var level string
	if rates > 60 {
		level = "High"
	} else if rates > 30 {
		level = "Medium"
	} else {
		level = "Low"
	}

	return level
}
func computesOccupancyRates(rooms, allRooms float32) float32 {
	return (rooms / allRooms) * 100
}
