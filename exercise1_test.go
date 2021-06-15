package main

import (
	"fmt"
	"testing"
)

func TestCamelCase(t *testing.T) {
	var tests = []struct {
		s string
		want int
	}{
		{"saveChangesInTheEditor", 5},
		{"save", 1},
		{"", 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := camelCase(tt.s)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}