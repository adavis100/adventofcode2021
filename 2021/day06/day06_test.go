package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDay6(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		days     int
		expected uint64
	}{
		{"0 days", "3,4,3,1,2", 0, 5},
		{"1 day", "3,4,3,1,2", 1, 5},
		{"2 days", "3,4,3,1,2", 2, 6},
		{"3 days", "3,4,3,1,2", 3, 7},
		{"18 days", "3,4,3,1,2", 18, 26},
		{"80 days", "3,4,3,1,2", 80, 5934},
		{"256 days", "3,4,3,1,2", 256, 26984457539},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := Solve(strings.NewReader(test.in), test.days)
			assert.Equal(t, test.expected, n)
		})
	}
}
