/*
Hands on Exercise 24

● create a program that uses both SEQUENTIAL and CONDITIONAL control flow. Your program should do the following
	○ create a random int between 0 and 250
	○ store the value of that int in a variable with the identifier of x
		■ func Intn(n int) int
			● rand.Intn()
	○ print out the name and value of the variable
	○ use an if statement to do the following
		■ if the value is between 0 and 100
			● print between 0 and 100
		■ if the value is between 101 and 200
			● print between 101 and 200
		■ if the value is between 201 and 250
			● print between 201 and 250
● re: inclusive, exclusive – does rand.Intn()
	○ include zero in its output? - Yes
	○ include 250 in its output? - yes
	○ show this in code using the numbers 0 - 3

*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(400)
	fmt.Printf("x = %v\n", x)

	if x <= 100 {
		fmt.Println("x is between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("x is between 101 and 200")
	} else if x >= 101 && x <= 250 {
		fmt.Println("x is between 201 and 250")
	} else {
		fmt.Println("x is over 250")
	}

	// [0, 3)
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))

}
