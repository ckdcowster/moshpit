/*
Hands-on exercise #34
	● create a loop that runs 5 times
	● nest a loop within the first loop that also prints 5 times
	● print something in each loop to illustrate what is occurring
*/

package main

import "fmt"

func main() {

	for x := 0; x < 5; x++ {
		fmt.Printf("outer loop is %v\t", x)
		for y := 0; y < 5; y++ {
			fmt.Printf("inner loop is %v\t", y)
		}
		fmt.Println()
	}
}
