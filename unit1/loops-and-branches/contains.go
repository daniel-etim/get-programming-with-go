// check that a character(s) is contained in a string
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Did you enter the cave?")

	var answer = "yes, i did"
	var check = strings.Contains(answer, "yes")

	fmt.Println("Did you enter the cave:", check)

	// ------------ similitude ------------
	fmt.Println("You find yourself in a dimly lit cavern.")

	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)
}