package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func countNums(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count += countNumsInLine(line)
	}
	return count
}

func countNumsInLine(s string) int {
	parts := strings.Split(s, " | ")
	digits := strings.Split(parts[1], " ")
	count := 0
	for _, s := range digits {
		if isUnique(len(s)) {
			count++
		}
	}
	return count
}

func isUnique(length int) bool {
	return length == 2 || length == 3 || length == 4 || length == 7
}

func main() {
	file, err := os.Open("./2021/day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(countNums(file))
}
