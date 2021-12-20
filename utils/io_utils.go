package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func LoadIntList(r io.Reader) []int {
	s, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	nums := make([]int, 0)
	for _, numStr := range strings.Split(string(s), ",") {
		n, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}
	return nums
}

func LoadGrid(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)
	grid := make([][]int, 0)
	for scanner.Scan() {
		l := scanner.Text()
		arr := strings.Split(l, "")
		row := make([]int, 0)
		for _, numStr := range arr {
			n, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, n)
		}
		grid = append(grid, row)
	}
	return grid
}
