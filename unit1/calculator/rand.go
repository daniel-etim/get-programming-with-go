// generates a number from 1 -10
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var num = rand.N(10) + 1
	fmt.Printf("%v was chosen\n", num)

	num = rand.N(10) + 1
	fmt.Printf("%v was also chosen", num)

}