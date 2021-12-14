package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetMinCost(t *testing.T) {
	in := "16,1,2,0,4,2,7,1,2,14"

	min := GetMinCost(strings.NewReader(in), false)
	assert.Equal(t, 37, min)
}

func TestGetMinCost2(t *testing.T) {
	in := "16,1,2,0,4,2,7,1,2,14"

	min := GetMinCost(strings.NewReader(in), true)
	assert.Equal(t, 168, min)
}
