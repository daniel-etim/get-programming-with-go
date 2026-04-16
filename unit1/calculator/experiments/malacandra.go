// Malacandra is another name for Mars in The Space Trilogy by C. S. Lewis. 
// Write a program to determine how fast a ship would need to travel (in km/h) in order to reach Malacandra in 28 days. 
// Assume a distance of 56,000,000 km.

package main

import "fmt"

func main() {
	var (
		distance = 56000000 // km
		days = 28 // days
		hours = days * 24 // hours
		speed = distance / hours
	)

	fmt.Printf("Avg speed to reach Malacandra in 28 days: %5vkm/h\n", speed)
}