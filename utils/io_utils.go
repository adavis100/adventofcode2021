package utils

import (
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
