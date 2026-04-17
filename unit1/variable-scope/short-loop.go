// short declaration in a for loop

package main

import "fmt"

func main() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}

	fmt.Println(count) // <-------- this is results in error: undefined
}

// For the best readability, declare variables near where they are used.