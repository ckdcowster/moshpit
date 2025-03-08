/*
Hands-on exercise #32
	● create a for loop that uses break statement
	● increment or decrement the variable being checked as a condition in the loop
*/

package main

import "fmt"

func main() {

	x := 11
	for {
		if x < 0 {
			break
		}
		fmt.Printf("x is %v\n", x)
		x--
	}
	fmt.Println("Done!")
}
