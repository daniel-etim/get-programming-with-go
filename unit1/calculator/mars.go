// My weight loss program

package main

import "fmt"

// main is the function where it all begins
func main() {
	fmt.Println("My weight on the surface of  Mars is ")
	fmt.Println(78 * 0.3783)
	fmt.Println(" kg, and i would be ")
	fmt.Println(19 * 365 / 687)
	fmt.Println(" years old")

	fmt.Printf("My weight on the surface of Mars is %v kg, ", 78*0.3783)
	fmt.Printf("and i would be %v years old.\n", 19*365/687)

	fmt.Printf("My weight on Earth is %15vkg \n", 70)
	fmt.Printf("And my age is %15v \n", 25)
}
