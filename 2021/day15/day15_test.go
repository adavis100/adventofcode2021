package main

import (
	"github.com/adavis100/aoc/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var ex = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

func TestExample1(t *testing.T) {
	assert.Equal(t, 40, Solve1(strings.NewReader(ex)))
}

func TestBuildsBigGrid(t *testing.T) {
	ex := `5`
	big := `56789
67891
78912
89123
91234`
	grid := utils.LoadGrid(strings.NewReader(ex))
	grid = expandGrid(grid)
	assert.Equal(t, utils.LoadGrid(strings.NewReader(big)), grid)
}

func TestExample2(t *testing.T) {
	assert.Equal(t, 315, Solve2(strings.NewReader(ex)))
}
