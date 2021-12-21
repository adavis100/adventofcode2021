package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./2021/day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(Solve1(file))
	file.Seek(0, io.SeekStart)
	fmt.Println(Solve2(file))
}

func Solve1(r io.Reader) int {
	graph := LoadGraph(r)
	return dfs(graph, "start", []string{"start"})
}

func LoadGraph(r io.Reader) map[string][]string {
	graph := make(map[string][]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, "-")

		if dests, ok := graph[points[0]]; ok {
			dests = append(dests, points[1])
			graph[points[0]] = dests
		} else {
			graph[points[0]] = []string{points[1]}
		}

		if dests, ok := graph[points[1]]; ok {
			dests = append(dests, points[0])
			graph[points[1]] = dests
		} else {
			graph[points[1]] = []string{points[0]}
		}
	}
	return graph
}

func dfs(graph map[string][]string, from string, visited []string) int {
	count := 0
	for _, s := range graph[from] {
		if s == "end" {
			count++
		} else if isUpper(s) {
			count += dfs(graph, s, visited)
		} else if isLower(s) && !contains(visited, s) {
			count += dfs(graph, s, append(visited, s))
		}
	}
	return count
}

func isUpper(s string) bool {
	return strings.ToLower(s) != s
}

func isLower(s string) bool {
	return !isUpper(s)
}

func contains(visited []string, s string) bool {
	for _, cave := range visited {
		if s == cave {
			return true
		}
	}
	return false
}

func Solve2(r io.Reader) int {
	graph := LoadGraph(r)
	return dfs2(graph, "start", map[string]int{})
}

func dfs2(graph map[string][]string, from string, visited map[string]int) int {
	count := 0
	for _, s := range graph[from] {
		if s == "end" {
			count++
		} else if isUpper(s) {
			count += dfs2(graph, s, visited)
		} else if isLower(s) && canVisit(visited, s) {
			visited[s]++
			count += dfs2(graph, s, visited)
			visited[s]--
		}
	}
	return count
}

func canVisit(visited map[string]int, s string) bool {
	if s == "start" {
		return false
	}
	if visited[s] == 0 {
		return true
	}
	if visited[s] > 1 {
		return false
	}
	for _, v := range visited {
		if v == 2 {
			return false
		}
	}
	return true
}
