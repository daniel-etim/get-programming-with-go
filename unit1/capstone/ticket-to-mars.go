// Write a ticket generator in the Go Playground that makes use of variables, constants, switch, if, and for. 
// It should also draw on the fmt and math/rand packages to display and align text and to generate random numbers.

package main

import (
	"fmt"
	"math/rand"
)

const seconds_per_day = 86400

func main() {
	distance := 62000000
	company := ""
	trip := ""

	fmt.Println("Spaceline          Days      TripType     Price")
	fmt.Println("=================================================")

	for range(10) {
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventure"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}

		speed := rand.Intn(15) + 16 // km/s
		duration := distance / speed / seconds_per_day // in days
		price := 20.0 + speed

		if rand.Intn(2) == 1 {
			trip = "Round Trip"
			price = price * 2
		} else {
			trip = "one-way"
		}

		fmt.Printf("%-16v %4v     %10v    $ %-10v \n", company, duration, trip, price)
	}
}