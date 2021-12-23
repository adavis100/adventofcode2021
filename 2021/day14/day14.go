package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("./2021/day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(Solve1(file))
	file.Seek(0, io.SeekStart)
	//fmt.Println(Solve2(file))
}

func Solve1(r io.Reader) int {
	return solveIterations(r, 10)
}

func solveIterations(r io.Reader, iterations int) int {
	template, rules := loadInput(r)
	cur := template
	for i := 0; i < iterations; i++ {
		cur = doInsertion(cur, rules)
	}
	most, least := findMostAndLeastCommon(cur)
	return most - least
}

func loadInput(r io.Reader) (template string, rules map[string]string) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	template = scanner.Text()
	scanner.Scan()

	rules = make(map[string]string)
	for scanner.Scan() {
		l := scanner.Text()
		parts := strings.Split(l, " -> ")
		rules[parts[0]] = parts[1]
	}
	return template, rules
}

func doInsertion(cur string, rules map[string]string) string {
	sb := strings.Builder{}
	for i := 1; i < len(cur); i++ {
		pair := string(cur[i-1]) + string(cur[i])
		insert := rules[pair]
		if i == 1 {
			sb.WriteByte(cur[i-1])
		}
		sb.WriteString(insert)
		sb.WriteByte(cur[i])
	}
	return sb.String()
}

func findMostAndLeastCommon(cur string) (most int, least int) {
	elements := make([]int, 26)
	for i := 0; i < len(cur); i++ {
		elements[cur[i]-'A']++
	}
	sort.Slice(elements, func(i, j int) bool {
		return elements[i] < elements[j]
	})
	for i := len(elements) - 1; i >= 0 && elements[i] != 0; i-- {
		least = elements[i]
	}
	return elements[len(elements)-1], least
}

func Solve2(r io.Reader) int {
	return solveIterations(r, 40)
}
