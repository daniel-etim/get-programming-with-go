// branching with switch. using the "fallthrough" keyword

package main

import "fmt"

func main() {
	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern")
	case room == "lake":
		fmt.Println("The ice seem cold enough")
		fallthrough
	case room == "underwater":
		fmt.Println("The water is freezing cold")
	}
}