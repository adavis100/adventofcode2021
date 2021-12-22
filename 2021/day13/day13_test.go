package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var ex = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func TestLoadsInput(t *testing.T) {
	grid, folds := loadGrid(strings.NewReader(ex))
	assert.Equal(t, true, grid[14][0])
	assert.Equal(t, 2, len(folds))
	assert.Equal(t, "y", folds[0].dir)
	assert.Equal(t, 7, folds[0].pos)
}

func TestExample1(t *testing.T) {
	assert.Equal(t, 17, Solve1(strings.NewReader(ex)))
}
