// exercising the if else conditions, Note: Both else if and else are optional. When there are several paths to consider, you can repeat else if as many times as needed.

package main

import "fmt"

func main() {
	var command = "go left"
	if command == "go east" {
		fmt.Println("You head further up the mountain")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live most of your life!")
	} else {
		fmt.Println("Didnt quite get that.")
	}
}