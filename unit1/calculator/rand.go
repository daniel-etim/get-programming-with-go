// generates a number from 1 -10
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Printf("%v was chosen\n", num)

	num = rand.Intn(10) + 1
	fmt.Printf("%v was also chosen", num)

}