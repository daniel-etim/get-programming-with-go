// Short declaration in a if statement

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	if num := rand.N(3); num == 0 {
		fmt.Println("space adventures")
	} else if num == 1 {
		fmt.Println("Space X")
	} else {
		fmt.Print("Virgin Galactic")
	}
}