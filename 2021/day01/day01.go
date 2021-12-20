package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("2021/day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nums := make([]int, 0)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, x)
	}

	fmt.Println(CountIncreases(nums))
	fmt.Println(CountIncreases(combine3(nums)))
}

func CountIncreases(nums []int) int {
	prev := math.MaxInt
	count := 0
	for _, n := range nums {
		if n > prev {
			count++
		}
		prev = n
	}
	return count
}

func combine3(nums []int) []int {
	sums := make([]int, len(nums)-2)
	for i := 2; i < len(nums); i++ {
		sums[i-2] = nums[i-2] + nums[i-1] + nums[i]
	}
	return sums
}
