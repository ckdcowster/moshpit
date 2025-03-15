/*
Hands-on exercise 39
	bool logic test
*/

package main

import (
	"fmt"
)

func main() {

	fmt.Printf("true && true is %v\n", (true && true))
	fmt.Printf("true && false is  %v\n", (true && false))
	fmt.Printf("true || true is %v\n", (true || true))
	fmt.Printf("true || false is %v\n", (true || false))
	fmt.Printf("!true is %v\n", (!true))
}
