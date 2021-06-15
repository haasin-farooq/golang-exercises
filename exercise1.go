/*
Exercise 1: CamelCase
See problem statement here: https://gist.github.com/meer-online/e933ebf41da99a7c32530d5db839be45
*/

package main

import (
	"fmt"
)

func camelCase(s string) int {
	if s == "" {
		return 0
	}
	n := 1
	for _, v := range s {
		if v >= 65 && v <=90 {
			n += 1
		}
	}
	return n
}

func main() {
	fmt.Printf("Number of letter in camelCase string: %v\n", camelCase("saveChangesInTheEditor"))
}