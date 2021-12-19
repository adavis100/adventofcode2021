package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countNums(r io.Reader, onlyUnique bool) int {
	scanner := bufio.NewScanner(r)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if onlyUnique {
			count += countUniqueNumsInLine(line)
		} else {
			count += lineToNum(line)
		}
	}
	return count
}

func countUniqueNumsInLine(s string) int {
	_, digits := parseLine(s)
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

func lineToNum(l string) int {
	signals, digits := parseLine(l)
	nums := deduceNums(signals)

	m := make(map[string]string, 10)
	for i := 0; i < 10; i++ {
		m[sortStr(nums[i])] = strconv.Itoa(i)
	}
	numStr := ""
	for _, digit := range digits {
		numStr += m[sortStr(digit)]
	}
	n, err := strconv.Atoi(numStr)
	if err != nil {
		log.Println(err)
	}
	if n == 470 {
		fmt.Println("here")
	}
	fmt.Println(n)
	return n
}

func sortStr(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func deduceNums(signals []string) []string {
	nums := make([]string, 10)
	count := 0
	for count < 10 {
		for _, signal := range signals {
			if len(signal) == 2 && nums[1] == "" {
				nums[1] = signal
				count++
			} else if len(signal) == 3 && nums[7] == "" {
				nums[7] = signal
				count++
			} else if len(signal) == 4 && nums[4] == "" {
				nums[4] = signal
				count++
			} else if len(signal) == 7 && nums[8] == "" {
				nums[8] = signal
				count++
			} else if len(signal) == 5 {
				if nums[1] != "" && diffChars(signal, nums[1]) == 0 && nums[3] == "" {
					nums[3] = signal
					count++
				} else if nums[4] != "" && diffChars(signal, nums[4]) == 1 && nums[5] == "" && nums[3] != "" && signal != nums[3] {
					nums[5] = signal
					count++
				} else if nums[4] != "" && diffChars(signal, nums[4]) == 2 && nums[2] == "" {
					nums[2] = signal
					count++
				}
			} else if len(signal) == 6 {
				if nums[1] != "" && diffChars(signal, nums[1]) == 1 && nums[6] == "" {
					nums[6] = signal
					count++
				} else if nums[4] != "" && diffChars(signal, nums[4]) == 0 && nums[9] == "" {
					nums[9] = signal
					count++
				} else if nums[4] != "" && diffChars(signal, nums[4]) == 1 && nums[0] == "" && nums[6] != "" && signal != nums[6] {
					nums[0] = signal
					count++
				}
			}
		}
	}

	return nums
}

func diffChars(s string, chars string) int {
	n := 0
	for _, c := range chars {
		if !strings.Contains(s, string(c)) {
			n++
		}
	}
	return n
}

func parseLine(l string) (signals []string, digits []string) {
	parts := strings.Split(l, " | ")
	signals = strings.Split(parts[0], " ")
	digits = strings.Split(parts[1], " ")
	return signals, digits
}

func main() {
	file, err := os.Open("./2021/day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(countNums(file, true))
	file.Seek(0, io.SeekStart)
	fmt.Println(countNums(file, false))
}
