// repeats a for loop but decides whether to break or start over
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var degrees = 0

	for {
		fmt.Println(degrees)
		degrees++

		if degrees >= 360 {
			degrees = 0
			if rand.N(2) == 0 {
				break
			}
		}
	}
}