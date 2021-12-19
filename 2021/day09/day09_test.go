package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var in string = `2199943210
3987894921
9856789892
8767896789
9899965678
`

func TestBuildsGrid(t *testing.T) {
	grid := buildGrid(strings.NewReader(in))
	assert.Equal(t, 5, len(grid))
	assert.Equal(t, 10, len(grid[0]))
	assert.Equal(t, []int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0}, grid[0])
}

func TestGetsNeighbors(t *testing.T) {
	g := `123
456
789`
	grid := buildGrid(strings.NewReader(g))
	neighbors := getNeighbors(0, 0, grid)
	assert.Equal(t, []int{4, 2}, neighbors)

	neighbors = getNeighbors(0, 1, grid)
	assert.Equal(t, []int{1, 5, 3}, neighbors)

	neighbors = getNeighbors(1, 1, grid)
	assert.Equal(t, []int{2, 4, 8, 6}, neighbors)

	neighbors = getNeighbors(2, 2, grid)
	assert.Equal(t, []int{6, 8}, neighbors)
}

func TestSolve1(t *testing.T) {
	count := Solve1(strings.NewReader(in))
	assert.Equal(t, 15, count)
}

func TestCountsBasin(t *testing.T) {
	grid := buildGrid(strings.NewReader(in))
	count := countBasin(0, 9, grid)
	assert.Equal(t, 9, count)
}

func TestSolve2(t *testing.T) {
	count := Solve2(strings.NewReader(in))
	assert.Equal(t, 1134, count)
}
