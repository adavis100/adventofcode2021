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

	fmt.Println(Solve(file, 10))
	file.Seek(0, io.SeekStart)
	fmt.Println(Solve(file, 40))
}

func Solve(r io.Reader, iterations int) int64 {
	template, rules := loadInput(r)

	counts := doIterations(template, iterations, rules)
	most, least := findMostAndLeastCommon(counts)
	return most - least
}

func doIterations(s string, n int, rules map[string]string) []int64 {
	polyMap := buildPolyMap(s)
	for i := 0; i < n; i++ {
		newPolyMap := make(map[string]int64)
		for poly, count := range polyMap {
			if c, ok := rules[poly]; ok {
				newPolyMap[string(poly[0])+c] += count
				newPolyMap[c+string(poly[1])] += count
			}
		}
		polyMap = newPolyMap
	}
	return getCounts(polyMap, s[0], s[len(s)-1])
}

func buildPolyMap(s string) map[string]int64 {
	m := make(map[string]int64)
	for i := 1; i < len(s); i++ {
		m[string(s[i-1])+string(s[i])]++
	}
	return m
}

func getCounts(m map[string]int64, first byte, last byte) []int64 {
	counts := make([]int64, 26)
	for k, v := range m {
		counts[k[0]-'A'] += v
		counts[k[1]-'A'] += v
	}

	for i := 0; i < len(counts); i++ {
		counts[i] /= 2
	}

	counts[first-'A']++
	counts[last-'A']++

	return counts
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

func findMostAndLeastCommon(counts []int64) (most int64, least int64) {

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] < counts[j]
	})
	for i := len(counts) - 1; i >= 0 && counts[i] != 0; i-- {
		least = counts[i]
	}
	return counts[len(counts)-1], least
}
