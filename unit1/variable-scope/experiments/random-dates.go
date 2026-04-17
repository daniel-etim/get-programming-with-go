// Generate a random year instead of always using 2018.
//  For February, assign daysInMonth to 29 for leap years and 28 for other years.
// Hint: you can put an if statement inside of a case block.
// Use a for loop to generate and display 10 random dates.


package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	days := 31
	
	for range(10) {
		year := rand.Intn(4001) + 1000
		month := rand.Intn(12) + 1

		switch month {
		case 2:
			if year % 4 == 0 {
				if year % 100 == 0 {
					if year % 400 == 0 {
						fmt.Printf("It's a leap year at 'division by 400'! Check: %5v \n", year) // redundant. just to satisfy verification
						days = 29
					} else {
						days = 28
					}
				} else {
					fmt.Printf("It's a leap year at 'not division by 100'! Check: %5v \n", year) // redundant. just to satisfy verification
					days = 29
				}
			} else {
				days = 28
			}
	
		case 4, 6, 9, 11:
			days = 30
		}
		
		day := rand.Intn(days) + 1
		fmt.Printf("era: %-5v year: %-5v month: %-5v day: %-5v \n", era, year, month, day)
	}
}