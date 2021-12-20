package main

import (
	"fmt"
	"github.com/adavis100/aoc/utils"
	"io"
	"log"
	"math"
	"os"
)

func doOneStep(grid [][]int) {
	flashed := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		flashed[i] = make([]bool, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			inc(i, j, grid, flashed)
		}
	}
}

func inc(r int, c int, grid [][]int, flashed [][]bool) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || flashed[r][c] {
		return
	}

	grid[r][c] = (grid[r][c] + 1) % 10
	if grid[r][c] == 0 {
		flashed[r][c] = true
		inc(r-1, c, grid, flashed)
		inc(r-1, c-1, grid, flashed)
		inc(r-1, c+1, grid, flashed)
		inc(r, c-1, grid, flashed)
		inc(r, c+1, grid, flashed)
		inc(r+1, c, grid, flashed)
		inc(r+1, c-1, grid, flashed)
		inc(r+1, c+1, grid, flashed)
	}
}

func Solve1(r io.Reader) int {
	grid := utils.LoadGrid(r)
	count := 0
	for i := 0; i < 100; i++ {
		doOneStep(grid)
		count += getZeros(grid)
	}
	return count
}

func getZeros(grid [][]int) int {
	n := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				n++
			}
		}
	}
	return n
}

func Solve2(r io.Reader) int {
	grid := utils.LoadGrid(r)
	for i := 1; i < math.MaxInt32; i++ {
		doOneStep(grid)
		if getZeros(grid) == 100 {
			return i
		}
	}
	return -1
}

func main() {
	file, err := os.Open("./2021/day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(Solve1(file))
	file.Seek(0, io.SeekStart)
	fmt.Println(Solve2(file))
}
