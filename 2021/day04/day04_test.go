package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIsWinner(t *testing.T) {
	grid := make([][]int, 5)
	for i := 0; i < 5; i++ {
		grid[i] = []int{1, 2, 3, 4, 5}
	}
	b := NewBoard(grid)
	b.marked[3][0] = true
	b.marked[3][1] = true
	b.marked[3][2] = true
	b.marked[3][3] = true
	b.marked[3][4] = true

	if !IsWinner(b) {
		t.Errorf("Expected %v to be a winning Board", b)
	}
}

var in string = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

func TestExamplePart1(t *testing.T) {

	r := strings.NewReader(in)
	nums, b := ReadInput(r)

	if len(nums) != 27 {
		t.Errorf("Expected nums length to be 27 got %d", len(nums))
	}
	if nums[0] != 7 || nums[1] != 4 {
		t.Errorf("Expected nums to start with 7, 4 got %d, %d", nums[0], nums[1])
	}

	if len(b) != 3 {
		t.Errorf("Want len(b) = 3 got %d", len(b))
	}

	score := Run(nums, b)
	if score != 4512 {
		t.Errorf("got %d want %d", score, 4512)
	}
}

func TestGetScore(t *testing.T) {
	grid := [][]int{{1, 2, 3, 4, 5}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}
	b := NewBoard(grid)
	b.marked[0] = []bool{true, true, true, true, true}

	score := getScore(b, 5)
	if score != 100 {
		t.Errorf("got %d want %d", score, 100)
	}
}

func TestExamplePart2(t *testing.T) {
	r := strings.NewReader(in)
	nums, b := ReadInput(r)

	score := RunPart2(nums, b)

	assert.Equalf(t, 1924, score, "")
}
