// how long does it take to get to mars

package main

import "fmt"

func main() {
	const lightspeed = 299792 // km/s
	var (
		distance = 56000000 // km
		time = float64(distance) / float64(lightspeed)
)
	fmt.Printf("It takes %vseconds to reach the Mars when distance is %v \n", time, distance)

	distance = 401000000
	time = float64(distance) / float64(lightspeed)
	fmt.Printf("It takes %vseconds to reach the Mars. The distance: %v \n", time, distance)
}