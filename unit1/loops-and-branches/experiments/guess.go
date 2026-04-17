// Guess-the-number program. 
// Make the computer pick random numbers between 1–100 until it guesses your number, which you declare at the top of the program. 
// Display each guess and whether it was too big or too small.

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var my_num = 78

	for {
		var rand_num = rand.N(100) + 1
		if rand_num < my_num {
			fmt.Printf("%v : Go higher\n", rand_num)
		} else if rand_num > my_num {
			fmt.Printf("%v : Go lower\n", rand_num)
		} else {
			fmt.Printf("%v : Gotcha!\n", rand_num)
			break
		}
	}
}