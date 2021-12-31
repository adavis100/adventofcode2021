package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
)

type target struct {
	x1, x2, y1, y2 int
}

func main() {
	file, err := os.Open("./2021/day17/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(Solve1(file))
	file.Seek(0, io.SeekStart)
	fmt.Println(Solve2(file))
}

func Solve1(r io.Reader) int {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	t := target{}
	n, err := fmt.Sscanf(string(buf), "target area: x=%d..%d, y=%d..%d", &t.x1, &t.x2, &t.y1, &t.y2)
	if err != nil {
		log.Fatal(err)
	}
	if n != 4 {
		log.Fatal("Unable to parse input")
	}

	max := 0
	for vx := 0; vx < t.x2; vx++ {
		for vy := 0; vy < -t.y1; vy++ {
			thisMax := findMax(vx, vy, t)
			if thisMax > max {
				max = thisMax
			}
		}
	}

	return max
}

func findMax(vx int, vy int, t target) int {
	max := 0
	x, y := 0, 0
	for i := 0; i < math.MaxInt; i++ {
		if inTarget(x, y, t) {
			return max
		} else if invalid(x, y, t) {
			return 0
		}
		if y > max {
			max = y
		}
		x += vx
		if vx > 0 {
			vx--
		} else if vx < 0 {
			vx++
		}
		y += vy
		vy--
	}
	return 0
}

func inTarget(x int, y int, t target) bool {
	return x >= t.x1 && x <= t.x2 && y >= t.y1 && y <= t.y2
}

func invalid(x int, y int, t target) bool {
	return x > t.x2 || y < t.y2
}

func Solve2(r io.Reader) int {
	fmt.Println("Unimplemented")
	return 0
}
