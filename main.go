/*
Hands on Exercise 25

● create a program that uses both SEQUENTIAL and CONDITIONAL control flow. Your program should do the following
	○ create a random int between 0 and 250
	○ store the value of that int in a variable with the identifier of x
		■ func Intn(n int) int
			● rand.Intn()
	○ print out the name and value of the variable
	○ use a switch statement to do the following
		■ the value is between 0 and 100
			● print between 0 and 100
		■ the value is between 101 and 200
			● print between 101 and 200
		■ the value is between 201 and 250
			● print between 201 and 250
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(400)
	fmt.Printf("x = %v\n", x)

	switch {
	case x <= 100:
		fmt.Println("x is between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("x is between 101 and 200")
	case x >= 101 && x <= 250:
		fmt.Println("x is between 201 and 250")
	default:
		fmt.Println("x is over 250")
	}

}
