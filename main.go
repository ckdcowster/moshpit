/*
Hands-on exercise #37
	● use the code from the previous exercise
	● add this code to print a single value stored in the map
		age := m["James"]
		fmt.Println(age)
	● now use similar code to use the lookup of "Q" and print that value
	● now use the "comma ok" idiom to test whether "Q" is actually stored in the map, then
		print out a statement if it is not stored in the map
		○ hint: check effective go for the "comma ok"*/

package main

import "fmt"

func main() {

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	age := m["James"]
	fmt.Println(age)

	if age, ok := m["Q"]; !ok {
		fmt.Println("cannot find Q in map, and the zero value of int is ", age)
	}

	if age, ok := m["James"]; ok {
		fmt.Println("James Bond age is ", age)
	}
}
