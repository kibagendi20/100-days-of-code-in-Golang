package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	hotelName := "The Gopher Hotel"
	fmt.Println("Hotel " + hotelName)
	var roomsAvailable int
	rand.Seed(time.Now().UTC().UnixNano())
	var rooms, roomsOccupied = 100, rand.Intn(100)

	roomsAvailable = rooms - roomsOccupied
	fmt.Println(roomsAvailable, "rooms available")
	// ageFunctionPrint1()
	// ageFunctionPrint2()
	// switchFunc()
	forFunc()
	forFunc1()
}

func ageFunctionPrint1() {
	rand.Seed(time.Now().UTC().UnixNano())

	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)
	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else {
		fmt.Println("Paul is younger than John,or both have the same age")
	}
}
func ageFunctionPrint2() {
	rand.Seed(time.Now().UTC().UnixNano())

	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)
	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else {
		if agePaul == ageJohn {
			fmt.Println("Paul and Jphn have the same age")
		} else {

			fmt.Println("Paul is younger than John")

		}
	}
}

func switchFunc() {
	rand.Seed(time.Now().UTC().UnixNano())

	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)
	//switch expression
	switch ageJohn {
	case 10:
		fmt.Println("John is 10 years old")
	case 20:
		fmt.Println("John is 20 years old")
	case 100:
		fmt.Println("John is 100 years old")
	default:
		fmt.Println("John has neither 10, 20 nor 100 years old")

	}

	//switch statement; expression

	switch ageSum := ageJohn + agePaul; ageSum {
	case 10:
		fmt.Println("ageJohn + agePaul = 10")
	case 20, 30, 40:
		fmt.Println("ageJohn + agePaul = 10 or 30 or 40")
	case 2 * agePaul:
		fmt.Println("ageJohn + agePaul = 2 times agePaul")
	}

	//switch (without statement, without expression)

	switch {
	case agePaul > ageJohn:
		fmt.Println("agePaul > ageJohn")
	case agePaul == ageJohn:
		fmt.Println("agePaul == ageJohn")
	case agePaul < ageJohn:
		fmt.Println("agePaul < ageJohn")
	}
}

func forFunc() {
	const emailToSend = 3
	emailSent := 0

	for emailSent < emailToSend {
		fmt.Println("sending email..")
		emailSent++
	}
	fmt.Println("End of Program")
}

func forFunc1() {
	rand.Seed(time.Now().UTC().UnixNano())
	var ageJohn int = rand.Intn(110)
	fmt.Println("John is", ageJohn, "years old")

	for i := 0; 1 < ageJohn; i++ {
		fmt.Println("interation N", i)
	}
}
