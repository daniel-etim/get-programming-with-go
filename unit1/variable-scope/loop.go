// a condensed countdown

package main

import (
	"fmt"
)

func main() {
	var count = 0

	for count = 10; count > 5; count-- {
		fmt.Println(count)
	}

	fmt.Println("check")
	fmt.Println(count) // count remains in scope
}