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

type fold struct {
	dir string
	pos int
}

func main() {
	file, err := os.Open("./2021/day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(Solve1(file))
	file.Seek(0, io.SeekStart)
	Solve2(file)
}

func Solve1(r io.Reader) int {
	grid, folds := loadGrid(r)

	if folds[0].dir == "y" {
		grid = foldY(grid, folds[0].pos)
	} else if folds[0].dir == "x" {
		grid = foldX(grid, folds[0].pos)
	}

	return count(grid)
}

func count(grid [][]bool) int {
	n := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				n++
			}
		}
	}
	return n
}

func loadGrid(r io.Reader) (grid [][]bool, folds []fold) {
	grid = make([][]bool, 2000)
	for i := 0; i < 2000; i++ {
		grid[i] = make([]bool, 2000)
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	l := scanner.Text()
	for l != "" {
		coord := strings.Split(l, ",")
		y, err := strconv.Atoi(coord[0])
		if err != nil {
			log.Fatal(err)
		}
		x, err := strconv.Atoi(coord[1])
		if err != nil {
			log.Fatal(err)
		}
		grid[x][y] = true
		scanner.Scan()
		l = scanner.Text()
	}

	folds = make([]fold, 0)
	for scanner.Scan() {
		l = scanner.Text()
		l = strings.ReplaceAll(l, "fold along ", "")
		parts := strings.Split(l, "=")
		pos, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		folds = append(folds, fold{parts[0], pos})
	}
	return grid, folds
}

func foldX(grid [][]bool, pos int) [][]bool {
	newGrid := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		newGrid[i] = make([]bool, pos)
		rpos := 2 * pos
		for j := 0; j < pos; j++ {
			newGrid[i][j] = grid[i][j] || grid[i][rpos]
			rpos--
		}
	}
	return newGrid
}

func foldY(grid [][]bool, pos int) [][]bool {
	newGrid := make([][]bool, pos)
	for i := 0; i < pos; i++ {
		newGrid[i] = make([]bool, len(grid[i]))
	}
	bpos := 2 * pos
	for r := 0; r < pos; r++ {
		for c := 0; c < len(grid[r]); c++ {
			newGrid[r][c] = grid[r][c] || grid[bpos][c]
		}
		bpos--
	}
	return newGrid
}

func Solve2(r io.Reader) {
	grid, folds := loadGrid(r)

	for i := 0; i < len(folds); i++ {
		if folds[i].dir == "y" {
			grid = foldY(grid, folds[i].pos)
		} else if folds[i].dir == "x" {
			grid = foldX(grid, folds[i].pos)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
