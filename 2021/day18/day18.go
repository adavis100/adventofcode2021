package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type snailnum struct {
	prev, next *snailnum
	val        int
	nest       int
}

func main() {
	file, err := os.Open("./2021/day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println(Solve1(file))
	file.Seek(0, io.SeekStart)
	fmt.Println(Solve2(file))
}

func Solve1(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	sum := loadSnailNum(scanner.Text())

	for scanner.Scan() {
		sum = add(sum, loadSnailNum(scanner.Text()))
	}

	return getMagnitude(sum)
}

func loadSnailNum(s string) *snailnum {
	nest := 0
	buf := []byte(s)
	head := &snailnum{}
	cur := head
	var prev *snailnum
	for i := 0; i < len(buf); i++ {
		if buf[i] == '[' {
			nest++
		} else if buf[i] == ']' {
			nest--
		} else if buf[i] >= '0' && buf[i] <= '9' {
			if buf[i+1] >= '0' && buf[i+1] <= '9' {
				cur.val = int(10 + buf[i+1] - '0')
				i++
			} else {
				cur.val = int(buf[i] - '0')
			}
			cur.nest = nest
			cur.prev = prev
			cur.next = &snailnum{}
			prev = cur
			cur = cur.next
		}
	}
	prev.next = nil
	return head
}

func add(s1 *snailnum, s2 *snailnum) *snailnum {
	sn := combine(s1, s2)
	reduce(sn)
	return sn
}

func combine(s1 *snailnum, s2 *snailnum) *snailnum {
	sn := s1
	var prev *snailnum
	for cur := sn; cur != nil; cur = cur.next {
		prev = cur
		cur.nest++
	}
	prev.next = s2
	s2.prev = prev
	for cur := s2; cur != nil; cur = cur.next {
		cur.nest++
	}
	return sn
}

func reduce(sn *snailnum) *snailnum {
	updated := true
	for updated {
		updated = explode(sn)
		if updated {
			continue
		}
		updated = split(sn)
	}
	return sn
}

func split(sn *snailnum) bool {
	var changed bool
	for cur := sn; cur != nil; cur = cur.next {
		if cur.val > 9 {
			rem := cur.val % 2
			half := cur.val / 2
			cur.val = half
			cur.nest++
			next := cur.next
			cur.next = &snailnum{prev: cur, next: next, nest: cur.nest, val: half + rem}
			if next != nil {
				next.prev = cur.next
			}
			changed = true
			break
		}
	}
	return changed
}

func explode(sn *snailnum) bool {
	var changed bool
	for cur := sn; cur != nil; cur = cur.next {
		if cur.nest == 5 {
			if cur.prev != nil {
				cur.prev.val += cur.val
			}
			cur.val = 0
			cur.nest--
			if cur.next.next != nil {
				cur.next.next.val += cur.next.val
				cur.next.next.prev = cur
			}
			cur.next = cur.next.next
			changed = true
			break
		}
	}
	return changed
}

func getMagnitude(sn *snailnum) int {
	head := sn
	for head.next != nil {
		for cur := head; cur != nil && cur.next != nil; cur = cur.next {
			if cur.nest == cur.next.nest {
				cur.val = 3*cur.val + 2*cur.next.val
				cur.nest--
				cur.next = cur.next.next
				if cur.next != nil {
					cur.next.prev = cur
				}
			}
		}
	}
	return sn.val
}

func Solve2(r io.Reader) int {
	return 0
}
