package main

import (
	"container/heap"
	"fmt"
	"github.com/adavis100/aoc/utils"
	"io"
	"log"
	"os"
)

type Point struct {
	x int
	y int
}

type Coord struct {
	row  int
	col  int
	cost int
}

type CoordHeap []*Coord

func (h CoordHeap) Len() int           { return len(h) }
func (h CoordHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h CoordHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CoordHeap) Push(x interface{}) {
	*h = append(*h, x.(*Coord))
}

func (h *CoordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	file, err := os.Open("./2021/day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(Solve1(file))
	file.Seek(0, io.SeekStart)
	fmt.Println(Solve2(file))
}

func Solve1(r io.Reader) int {
	grid := utils.LoadGrid(r)
	paths := getPaths(grid)

	return paths[len(grid)-1][len(grid[0])-1]
}

func getPaths(grid [][]int) [][]int {
	grid[0][0] = 0
	h := &CoordHeap{}
	heap.Init(h)
	heap.Push(h, &Coord{0, 0, grid[0][0]})
	paths := make([][]int, len(grid))
	for i := 0; i < len(paths); i++ {
		paths[i] = make([]int, len(grid[i]))
	}
	paths[0][0] = grid[0][0]
	visited := make(map[Point]bool)
	visited[Point{0, 0}] = true

	for h.Len() > 0 {
		c := heap.Pop(h).(*Coord)
		paths[c.row][c.col] = c.cost
		neighbors := getNeighbors(c, grid, visited)
		for _, coord := range neighbors {
			heap.Push(h, coord)
		}
	}
	return paths
}

func getNeighbors(c *Coord, grid [][]int, visited map[Point]bool) []*Coord {
	coords := make([]*Coord, 0)
	if c.col > 0 && !visited[Point{c.row, c.col - 1}] {
		visited[Point{c.row, c.col - 1}] = true
		coords = append(coords, &Coord{c.row, c.col - 1, c.cost + grid[c.row][c.col-1]})
	}
	if c.col < len(grid[c.row])-1 && !visited[Point{c.row, c.col + 1}] {
		visited[Point{c.row, c.col + 1}] = true
		coords = append(coords, &Coord{c.row, c.col + 1, c.cost + grid[c.row][c.col+1]})
	}
	if c.row > 0 && !visited[Point{c.row - 1, c.col}] {
		visited[Point{c.row - 1, c.col}] = true
		coords = append(coords, &Coord{c.row - 1, c.col, c.cost + grid[c.row-1][c.col]})
	}
	if c.row < len(grid)-1 && !visited[Point{c.row + 1, c.col}] {
		visited[Point{c.row + 1, c.col}] = true
		coords = append(coords, &Coord{c.row + 1, c.col, c.cost + grid[c.row+1][c.col]})
	}
	return coords
}

func Solve2(r io.Reader) int {
	grid := utils.LoadGrid(r)
	grid = expandGrid(grid)
	paths := getPaths(grid)
	return paths[len(grid)-1][len(grid[0])-1]
}

func expandGrid(grid [][]int) [][]int {
	newGrid := make([][]int, 5*len(grid))
	for i := 0; i < 5*len(grid[0]); i++ {
		newGrid[i] = make([]int, 5*len(grid[0]))
	}

	for multr := 0; multr < 5; multr++ {
		for multc := 0; multc < 5; multc++ {
			for r := 0; r < len(grid); r++ {
				for c := 0; c < len(grid[r]); c++ {
					newRisk := grid[r][c] + multr + multc
					if newRisk > 9 {
						newRisk = newRisk%10 + 1
					}
					newr := r + multr*len(grid)
					newc := c + multc*len(grid[0])
					newGrid[newr][newc] = newRisk
				}
			}
		}
	}
	return newGrid
}
