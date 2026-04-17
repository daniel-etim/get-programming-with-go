package main

import (
	"fmt"
	"math/rand/v2"
)

// testing the modern math/rand/v2 package for simple and effecient use

func main() {
	for range(30) {
		num := rand.N(1000)
		fmt.Println(num)
	}
}