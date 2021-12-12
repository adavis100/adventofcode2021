package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	nums   [][]int
	marked [][]bool
}

func NewBoard(grid [][]int) Board {
	nums := make([][]int, 5)
	marked := make([][]bool, 5)

	for i := 0; i < 5; i++ {
		marked[i] = make([]bool, 5)
		nums[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			nums[i][j] = grid[i][j]
		}
	}
	return Board{nums, marked}
}

func IsWinner(b Board) bool {
	for row := 0; row < 5; row++ {
		win := true
		for col := 0; col < 5; col++ {
			if !b.marked[row][col] {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}
	for col := 0; col < 5; col++ {
		win := true
		for row := 0; row < 5; row++ {
			if !b.marked[row][col] {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}
	return false
}

func ReadInput(r io.Reader) ([]int, []Board) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	s := scanner.Text()
	numsTxt := strings.Split(s, ",")
	nums := convertToNums(numsTxt)
	boards := make([]Board, 0)
	re := regexp.MustCompile("[ ]+")
	for scanner.Scan() != false {
		grid := make([][]int, 5)
		for i := 0; i < 5; i++ {
			scanner.Scan()
			s := strings.TrimSpace(scanner.Text())
			numsTxt := re.Split(s, -1)
			nums := convertToNums(numsTxt)
			grid[i] = nums
		}
		boards = append(boards, NewBoard(grid))
	}
	return nums, boards
}

func convertToNums(numsTxt []string) []int {
	nums := make([]int, len(numsTxt))
	for i := 0; i < len(nums); i++ {
		var err error
		nums[i], err = strconv.Atoi(numsTxt[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	return nums
}

func getScore(b Board, n int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				sum += b.nums[i][j]
			}
		}
	}
	return sum * n
}

func Run(nums []int, boards []Board) int {
	winner := -1
	for i := 0; i < len(nums) && winner < 0; i++ {
		mark(boards, nums[i])
		winner = getWinner(boards)
		if winner >= 0 {
			return getScore(boards[winner], nums[i])
		}
	}
	return -1
}

func getWinner(boards []Board) int {
	for i := 0; i < len(boards); i++ {
		if IsWinner(boards[i]) {
			return i
		}
	}
	return -1
}

func mark(boards []Board, n int) {
	for _, b := range boards {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if b.nums[i][j] == n {
					b.marked[i][j] = true
				}
			}
		}
	}
}

func RunPart2(nums []int, boards []Board) int {

	for i := 0; i < len(nums) && len(boards) > 0; i++ {
		mark(boards, nums[i])
		for getWinner(boards) >= 0 {
			winner := getWinner(boards)
			if len(boards) == 1 || i == len(nums)-1 {
				return getScore(boards[winner], nums[i])
			}
			boards[winner] = boards[len(boards)-1]
			boards = boards[:len(boards)-1]
		}
	}
	return -1
}

func main() {
	file, err := os.Open("2021/day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	nums, boards := ReadInput(file)

	fmt.Println(Run(nums, boards))

	file.Seek(0, io.SeekStart)
	nums, boards = ReadInput(file)
	fmt.Println(RunPart2(nums, boards))
}
