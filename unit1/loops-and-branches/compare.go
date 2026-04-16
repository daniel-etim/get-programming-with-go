// boolean: using comparison operators(e.g ==, <, >, <=, etc)
package main

import (
	"fmt"
)

func main() {
	fmt.Println("There's a sign near the entrance that reads: 'No minors'")

	var (
		age = 41
		minor = age < 41
	)

	fmt.Printf("At age %v, am i a minor? %v\n", age, minor)


	fmt.Println("Which is greater? Apple or Banana")
	fmt.Println("Lapple" > "Banana")
}