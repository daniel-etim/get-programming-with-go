// scoping rules

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var count = 0

	for count < 0 { // a new scope begins.
		var num = rand.N(10) + 1
		fmt.Println(num)
		
		count++
	} // the scope ends
}