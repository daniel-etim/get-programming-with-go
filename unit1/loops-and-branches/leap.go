// checks if a year is leap while it exercises the or and and operator: || and && respectively

package main

import "fmt"

func main() {
	var (
		year = 2000
		leap = year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)
	)

	fmt.Printf("This is year %v, should you leap?\n", year)

	if leap {
		fmt.Println("Yes, look before you leap!")
	} else {
		fmt.Println("NO! Keep your feet on the ground")
	}
}