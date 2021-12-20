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

type command string

const (
	forward command = "forward"
	up              = "up"
	down            = "down"
)

type move struct {
	cmd  command
	dist int
}

func main() {
	file, err := os.Open("2021/day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	moves := loadMoves(file)

	hpos, depth := run1(moves)
	fmt.Println(hpos * depth)

	hpos, depth = run2(moves)
	fmt.Println(hpos * depth)
}

func run1(moves []move) (int, int) {
	hpos := 0
	depth := 0
	for _, move := range moves {
		switch move.cmd {
		case forward:
			hpos += move.dist
		case up:
			depth -= move.dist
		case down:
			depth += move.dist
		}
	}
	return hpos, depth
}

func run2(moves []move) (int, int) {
	hpos := 0
	depth := 0
	aim := 0
	for _, move := range moves {
		switch move.cmd {
		case forward:
			hpos += move.dist
			depth += aim * move.dist
		case up:
			aim -= move.dist
		case down:
			aim += move.dist
		}
	}
	return hpos, depth
}

func loadMoves(r io.Reader) []move {
	scanner := bufio.NewScanner(r)
	moves := make([]move, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		cmd := command(line[0])
		dist, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		moves = append(moves, move{cmd, dist})
	}
	return moves
}
