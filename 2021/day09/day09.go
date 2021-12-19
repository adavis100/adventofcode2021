package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func getNeighbors(r, c int, grid [][]int) []int {
	neighbors := make([]int, 0)
	if r > 0 {
		neighbors = append(neighbors, grid[r-1][c])
	}
	if c > 0 {
		neighbors = append(neighbors, grid[r][c-1])
	}
	if r < len(grid)-1 {
		neighbors = append(neighbors, grid[r+1][c])
	}
	if c < len(grid[r])-1 {
		neighbors = append(neighbors, grid[r][c+1])
	}
	return neighbors
}

func buildGrid(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)
	grid := make([][]int, 0)
	for scanner.Scan() {
		l := scanner.Text()
		arr := strings.Split(l, "")
		row := make([]int, 0)
		for _, numStr := range arr {
			n, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, n)
		}
		grid = append(grid, row)
	}
	return grid
}

func Solve1(r io.Reader) int {
	grid := buildGrid(r)
	risk := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			neighbors := getNeighbors(i, j, grid)
			if isLow(grid[i][j], neighbors) {
				risk += grid[i][j] + 1
			}
		}
	}
	return risk
}

func isLow(n int, neighbors []int) bool {
	for _, neighbor := range neighbors {
		if n >= neighbor {
			return false
		}
	}
	return true
}

func Solve2(r io.Reader) string {
	return "TODO: implement part 2"
}

func main() {
	file, err := os.Open("./2021/day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(Solve1(file))
	file.Seek(0, io.SeekStart)
	fmt.Println(Solve2(file))
}
