/*
Hands-on exercise #36
	● below is the code to create a data structure called a map
	● put this code into a program
		m := map[string]int{
				"James": 42,
				"Moneypenny": 32,
			}
	● use a for range loop to print each value and the key associated with each value
*/

package main

import "fmt"

func main() {

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for key, value := range m {
		fmt.Printf("key %v \t value %v\n", key, value)
	}
}
