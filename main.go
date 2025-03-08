/*
Hands-on exercise #31
	● create a for loop using only a condition
	● increment or decrement the variable being checked in the condition
*/

package main

import "fmt"

func main() {

	x := 37
	for x < 42 {
		fmt.Printf("x is %v\n", x)
		x++
	}

	fmt.Println("Done!")
}
