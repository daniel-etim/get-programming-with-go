// Generate a random year instead of always using 2018.
//  For February, assign daysInMonth to 29 for leap years and 28 for other years.
// Hint: you can put an if statement inside of a case block.
// Use a for loop to generate and display 10 random dates.


package main

import (
	"fmt"
	"math/rand/v2"
)

var era = "AD"

func main() {
	days := 31
	
	for range(10) {
		year := rand.N(21) + 2005
		leap := year%4 == 0 && year%100 != 0 || year%400 == 0
		month := rand.N(12) + 1

		switch month {
		case 2:
			if leap {
				days = 29
			} else {
				days = 28
			}
		case 4, 6, 9, 11:
			days = 30
		}
		
		day := rand.N(days) + 1
		fmt.Printf("era: %-5v year: %-5v month: %-5v day: %-5v\n", era, year, month, day)
	}
}