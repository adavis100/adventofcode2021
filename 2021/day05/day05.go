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

type Coord struct {
	X int
	Y int
}

type Line struct {
	From Coord
	To   Coord
}

func parseLine(s string) Line {
	parts := strings.Split(s, " -> ")
	part1 := strings.Split(parts[0], ",")
	part2 := strings.Split(parts[1], ",")
	x, _ := strconv.Atoi(part1[0])
	y, _ := strconv.Atoi(part1[1])
	from := Coord{x, y}
	x, _ = strconv.Atoi(part2[0])
	y, _ = strconv.Atoi(part2[1])
	to := Coord{x, y}
	return Line{from, to}
}

func getPointsInLine(l Line, includeDiags bool) []Coord {
	var coords []Coord
	if l.From.X != l.To.X && l.From.Y != l.To.Y &&
		(!includeDiags || !isDiagonal(l.From, l.To)) {
		return coords
	}
	if (l.To.X < l.From.X && l.From.Y == l.To.Y) || (l.To.Y < l.From.Y && l.From.X == l.To.X) || l.To.X < l.From.X {
		l.From, l.To = l.To, l.From
	}
	if isHoriontal(l) {
		coords = make([]Coord, l.To.X-l.From.X+1)
		j := 0
		for i := l.From.X; i <= l.To.X; i++ {
			coords[j] = Coord{i, l.From.Y}
			j++
		}
	} else if isVertical(l) {
		coords = make([]Coord, l.To.Y-l.From.Y+1)
		j := 0
		for i := l.From.Y; i <= l.To.Y; i++ {
			coords[j] = Coord{l.From.X, i}
			j++
		}
	} else if isTranspose(l) {
		coords = make([]Coord, l.To.X-l.From.X+1)
		j := 0
		for i := 0; i < l.To.X-l.From.X+1; i++ {
			coords[j] = Coord{l.From.X + i, l.From.Y - i}
			j++
		}
	} else { // if diagonal
		coords = make([]Coord, l.To.X-l.From.X+1)
		j := 0
		for i := 0; i < l.To.X-l.From.X+1; i++ {
			coords[j] = Coord{l.From.X + i, l.From.Y + i}
			j++
		}
	}

	return coords
}

func isTranspose(l Line) bool {
	return l.From.Y > l.To.Y
}

func isDiagonal(from Coord, to Coord) bool {
	xDiff := from.X - to.X
	yDiff := from.Y - to.Y
	return xDiff == yDiff || xDiff == -yDiff
}

func isVertical(l Line) bool {
	return l.From.X == l.To.X
}

func isHoriontal(l Line) bool {
	return l.From.Y == l.To.Y
}

func readLines(r io.Reader) []Line {
	scanner := bufio.NewScanner(r)
	lines := make([]Line, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, parseLine(line))
	}
	return lines
}

func Solve(r io.Reader, includeDiag bool) int {
	grid := makeGrid()
	for _, l := range readLines(r) {
		for _, coord := range getPointsInLine(l, includeDiag) {
			grid[coord.X][coord.Y]++
		}
	}

	return countOverlaps(grid)
}

func makeGrid() [][]int {
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}
	return grid
}

func countOverlaps(grid [][]int) int {
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("2021/day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(Solve(file, false))
	file.Seek(0, io.SeekStart)
	fmt.Println(Solve(file, true))
}
