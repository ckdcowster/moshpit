/*
Hands-on exercise #33
	● use modulus and the continue statement in a loop to print all ODD numbers
*/

package main

import "fmt"

func main() {

	for i := 0; i < 15; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("x is %v\n", i)
		i++
	}
	fmt.Println("Done!")
}
