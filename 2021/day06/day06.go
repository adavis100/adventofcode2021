package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(r io.Reader, days int) uint64 {
	fish := parseInput(r)
	for i := 0; i < days; i++ {
		next := make([]uint64, 9)
		for i := 0; i < 9; i++ {
			if i == 0 {
				next[8] = fish[i]
				next[6] = fish[i]
			} else {
				next[i-1] += fish[i]
			}
		}
		fish = next
	}

	var count uint64 = 0
	for i := 0; i < 9; i++ {
		count += fish[i]
	}
	return count
}

func parseInput(r io.Reader) []uint64 {
	s, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	nums := make([]uint64, 9)
	for _, numStr := range strings.Split(string(s), ",") {
		n, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		nums[n]++
	}
	return nums
}

func main() {
	file, err := os.Open("2021/day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(Solve(file, 80))
	file.Seek(0, io.SeekStart)
	fmt.Println(Solve(file, 256))
}
