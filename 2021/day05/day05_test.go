package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParsesLine(t *testing.T) {
	line := parseLine("230,308 -> 429,308")
	assert.Equal(t, Coord{230, 308}, line.From, "from")
	assert.Equal(t, Coord{429, 308}, line.To, "to")
}

func TestGetPointsBetween(t *testing.T) {
	t.Run("Horizontal line", func(t *testing.T) {
		coords := getPointsInLine(Line{Coord{0, 9}, Coord{5, 9}}, false)
		assert.Equal(t, len(coords), 6)
		assert.Equal(t, Coord{0, 9}, coords[0])
		assert.Equal(t, Coord{1, 9}, coords[1])
	})
	t.Run("Vertical line", func(t *testing.T) {
		coords := getPointsInLine(Line{Coord{0, 1}, Coord{0, 9}}, false)
		assert.Equal(t, 9, len(coords))
		assert.Equal(t, Coord{0, 1}, coords[0])
		assert.Equal(t, Coord{0, 2}, coords[1])
	})
	t.Run("Horizontal line, reversed", func(t *testing.T) {
		coords := getPointsInLine(Line{Coord{5, 9}, Coord{0, 9}}, false)
		assert.Equal(t, 6, len(coords))
		assert.Equal(t, Coord{0, 9}, coords[0])
		assert.Equal(t, Coord{1, 9}, coords[1])
	})
	t.Run("Vertical line, reversed", func(t *testing.T) {
		coords := getPointsInLine(Line{Coord{0, 9}, Coord{0, 1}}, false)
		assert.Equal(t, len(coords), 9)
		assert.Equal(t, Coord{0, 1}, coords[0])
		assert.Equal(t, Coord{0, 2}, coords[1])
	})
	t.Run("Ignores diagonal lines", func(t *testing.T) {
		coords := getPointsInLine(Line{Coord{0, 9}, Coord{9, 0}}, false)
		assert.Equal(t, 0, len(coords))
	})
	t.Run("Handles diagonal line", func(t *testing.T) {
		coords := getPointsInLine(Line{Coord{0, 3}, Coord{3, 0}}, true)
		assert.Equal(t, 4, len(coords))
		assert.Equal(t, Coord{0, 3}, coords[0])
		assert.Equal(t, Coord{1, 2}, coords[1])
		assert.Equal(t, Coord{2, 1}, coords[2])
		assert.Equal(t, Coord{3, 0}, coords[3])
	})
	t.Run("Handles diagonal line, reversed", func(t *testing.T) {
		coords := getPointsInLine(Line{Coord{3, 0}, Coord{0, 3}}, true)
		assert.Equal(t, 4, len(coords))
		assert.Equal(t, Coord{0, 3}, coords[0])
		assert.Equal(t, Coord{1, 2}, coords[1])
		assert.Equal(t, Coord{2, 1}, coords[2])
		assert.Equal(t, Coord{3, 0}, coords[3])
	})
	t.Run("Handles another diagonal line", func(t *testing.T) {
		coords := getPointsInLine(Line{Coord{0, 0}, Coord{2, 2}}, true)
		assert.Equal(t, 3, len(coords))
		assert.Equal(t, Coord{0, 0}, coords[0])
		assert.Equal(t, Coord{1, 1}, coords[1])
		assert.Equal(t, Coord{2, 2}, coords[2])
	})
}

var in string = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestSolvesExample1(t *testing.T) {

	assert.Equal(t, 5, Solve(strings.NewReader(in), false))
}

func TestSolvesExample2(t *testing.T) {
	assert.Equal(t, 12, Solve(strings.NewReader(in), true))
}
