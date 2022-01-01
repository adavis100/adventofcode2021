package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var ex = "target area: x=20..30, y=-10..-5"

func TestSolvesExample(t *testing.T) {
	assert.Equal(t, 45, Solve1(strings.NewReader(ex)))
}

func TestSolvesExample2(t *testing.T) {
	assert.Equal(t, 112, Solve2(strings.NewReader(ex)))
}
