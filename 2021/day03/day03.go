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

func main() {
	file, err := os.Open("2021/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := loadLines(file)
	gamma, epsilon := part1MostLeast(lines)
	fmt.Println(gamma, epsilon)
	g, err := strconv.ParseInt(gamma, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	e, err := strconv.ParseInt(epsilon, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g * e)

	most, least := part2MostLeast(lines)
	oxy, err := strconv.ParseInt(most, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	co2, err := strconv.ParseInt(least, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oxy * co2)
}

func loadLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func part1MostLeast(lines []string) (string, string) {
	mostSb := strings.Builder{}
	leastSb := strings.Builder{}
	zeros, ones := countZerosOnes(lines)

	for pos := 0; pos < len(lines[0]); pos++ {
		if ones[pos] > zeros[pos] {
			mostSb.WriteString("1")
			leastSb.WriteString("0")
		} else {
			mostSb.WriteString("0")
			leastSb.WriteString("1")
		}
	}
	return mostSb.String(), leastSb.String()
}

func countZerosOnes(lines []string) ([]int, []int) {
	ones := make([]int, len(lines[0]))
	zeros := make([]int, len(lines[0]))
	for pos := 0; pos < len(lines[0]); pos++ {
		for i := 0; i < len(lines); i++ {
			s := []rune(lines[i])
			if string(s[pos]) == "0" {
				zeros[pos]++
			} else {
				ones[pos]++
			}
		}
	}
	return zeros, ones
}

func part2MostLeast(lines []string) (string, string) {
	mostNums := make([]string, len(lines))
	leastNums := make([]string, len(lines))
	for i, s := range lines {
		mostNums[i] = s
		leastNums[i] = s
	}

	most := findMost(mostNums)
	least := findLeast(leastNums)

	return most, least
}

func findMost(nums []string) string {
	zeros, ones := countZerosOnes(nums)
	for len(nums) > 1 {
		for pos := 0; pos < len(zeros) && len(nums) > 1; pos++ {
			filtered := make([]string, 0)
			for _, s := range nums {
				buf := []rune(s)
				if (zeros[pos] > ones[pos] && string(buf[pos]) == "0") ||
					(ones[pos] > zeros[pos] && string(buf[pos]) == "1") ||
					(ones[pos] == zeros[pos] && string(buf[pos]) == "1") {
					filtered = append(filtered, s)
				}
			}
			nums = filtered
			zeros, ones = countZerosOnes(nums)
		}
	}

	return nums[0]
}

func findLeast(nums []string) string {
	zeros, ones := countZerosOnes(nums)
	for len(nums) > 1 {
		for pos := 0; pos < len(zeros) && len(nums) > 1; pos++ {
			filtered := make([]string, 0)
			for _, s := range nums {
				buf := []rune(s)
				if (zeros[pos] > ones[pos] && string(buf[pos]) == "1") ||
					(ones[pos] > zeros[pos] && string(buf[pos]) == "0") ||
					(ones[pos] == zeros[pos] && string(buf[pos]) == "0") {
					filtered = append(filtered, s)
				}
			}
			nums = filtered
			zeros, ones = countZerosOnes(nums)
		}
	}
	return nums[0]
}
