package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func Solve1(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	score := 0
	for scanner.Scan() {
		score += getErrorScore(scanner.Text())
	}
	return score
}

func getErrorScore(l string) int {
	stack := make([]rune, 0)

	for _, c := range l {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '<':
			stack = append(stack, '>')
		default:
			expect := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c != expect {
				return getScore(c)
			}
		}
	}
	return 0
}

func getScore(r rune) int {
	m := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	return m[r]
}

func Solve2(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scores := make([]int, 0)
	for scanner.Scan() {
		l := scanner.Text()
		if getErrorScore(l) == 0 {
			scores = append(scores, getScoreForIncompleteLine(l))
		}
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	return scores[len(scores)/2]
}

func getScoreForIncompleteLine(l string) int {
	m := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	stack := buildRemainderStack(l)
	score := 0
	for len(stack) > 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		score *= 5
		score += m[c]
	}
	return score
}

func buildRemainderStack(l string) []rune {
	stack := make([]rune, 0)

	for _, c := range l {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '<':
			stack = append(stack, '>')
		default:
			stack = stack[:len(stack)-1]
		}
	}
	return stack
}

func main() {
	file, err := os.Open("./2021/day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(Solve1(file))
	file.Seek(0, io.SeekStart)
	fmt.Println(Solve2(file))
}
