// Variable scoping rules

package main

import (
	"fmt"
	"math/rand/v2"
)

var era = "AD" // era is available throughout the package

func main() {
	year := 2018 // era and year are in scope

	switch month := rand.N(12) + 1; month { // era, year, and month are in scope
	case 2:
		day := rand.N(28) + 1 // era, year, month, and day are in scope
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11: // its a new day
		day := rand.N(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.N(31) + 1
		fmt.Println(era, year, month, day)
	} // month and day are out of scope
} // year is no longer 