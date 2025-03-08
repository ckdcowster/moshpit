/*
Hands-on exercise #29
	● there are two parts ot this hands on exercise
		○ create a program that has a loop that prints every number from 0 to 99
		○ modify the program from the previous hands on exercise to run 100 times
*/

package main

import (
	"fmt"

	"math/rand"
)

func main() {

	// loop from 0 to 99
	for i := 0; i < 100; i++ {
		fmt.Printf("iteration is %v\t", i)
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("x = %v\ty = %v\t", x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("x and y are both less than 4")
		case x > 6 && y > 6:
			fmt.Println("x and y are both greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("none of the previous cases were met")
		}
	}
}
