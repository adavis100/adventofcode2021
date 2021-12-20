package main

import (
	"fmt"
	"github.com/adavis100/aoc/utils"
	"io"
	"log"
	"math"
	"os"
)

func main() {
	file, err := os.Open("2021/day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(GetMinCost(file, false))
	file.Seek(0, io.SeekStart)
	fmt.Println(GetMinCost(file, true))
}

func GetMinCost(r io.Reader, adjust bool) int {
	nums := utils.LoadIntList(r)

	max := findMax(nums)
	minCost := math.MaxInt32

	for i := 0; i < max; i++ {
		cost := getCost(nums, i, adjust)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func findMax(nums []int) int {
	max := 0
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func getCost(nums []int, pos int, adjust bool) int {
	cost := 0
	for _, n := range nums {
		if adjust {
			cost += sum(absDif(n, pos))
		} else {
			cost += absDif(n, pos)
		}
	}
	return cost
}

func sum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

func absDif(x int, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	}
	return diff
}
