// branching with switch. comparing one value to others

package main

import "fmt"

func main() {
	fmt.Println("There's a cavern entrance and path to the east")

	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You head further up the mountain")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern")
	case "read sign":
		fmt.Println("The sign reads: 'No minors")
	default:
		fmt.Println("didnt quite get that")
	}
}