package main

import "fmt"

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