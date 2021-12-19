package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coord struct {
	r int
	c int
}

func getNeighbors(r, c int, grid [][]int) []coord {
	neighbors := make([]coord, 0)
	if r > 0 {
		neighbors = append(neighbors, coord{r - 1, c})
	}
	if c > 0 {
		neighbors = append(neighbors, coord{r, c - 1})
	}
	if r < len(grid)-1 {
		neighbors = append(neighbors, coord{r + 1, c})
	}
	if c < len(grid[r])-1 {
		neighbors = append(neighbors, coord{r, c + 1})
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
			if isLow(grid[i][j], neighbors, grid) {
				risk += grid[i][j] + 1
			}
		}
	}
	return risk
}

func isLow(n int, neighbors []coord, grid [][]int) bool {
	for _, neighbor := range neighbors {
		if n >= grid[neighbor.r][neighbor.c] {
			return false
		}
	}
	return true
}

func Solve2(r io.Reader) int {
	grid := buildGrid(r)
	basins := make([]int, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			neighbors := getNeighbors(i, j, grid)
			if isLow(grid[i][j], neighbors, grid) {
				basins = append(basins, countBasin(i, j, grid))
			}
		}
	}
	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})
	return basins[0] * basins[1] * basins[2]
}

func countBasin(r int, c int, grid [][]int) int {
	queue := make([]coord, 0)
	queue = append(queue, coord{r, c})
	count := 0

	for len(queue) > 0 {
		coord := queue[0]
		queue = queue[1:]
		if grid[coord.r][coord.c] != 9 {
			count++
			grid[coord.r][coord.c] = 9
			for _, neighbor := range getNeighbors(coord.r, coord.c, grid) {
				queue = append(queue, neighbor)
			}
		}
	}
	return count
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
