// Short declaration in a switch statement

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("space adventures")
	case 1:
		fmt.Println("Space X")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random Spaceline #", num)
	}
}