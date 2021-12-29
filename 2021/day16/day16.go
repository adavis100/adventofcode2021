package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type packet struct {
	version  int
	pktType  int
	num      int
	children []packet
}

func main() {
	file, err := os.Open("./2021/day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(Solve1(file))
	file.Seek(0, io.SeekStart)
	fmt.Println(Solve2(file))
}

func Solve1(r io.Reader) int {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	s := string(bytes)
	p, _ := parsePacket(hexToBits(s))
	return countVersions(p)
}

func parsePacket(bits []bool) (*packet, int) {
	var size int
	v := toInt(bits[0:3])
	t := toInt(bits[3:6])
	p := &packet{version: v, pktType: t}
	switch t {
	case 4:
		size = 6
		more := true
		num := make([]bool, 0)
		for more {
			more = bits[size]
			num = append(num, bits[size+1:size+5]...)
			size += 5
		}
		p.num = toInt(num)
	default:
		lengthType := bits[6]
		switch lengthType {
		case true:
			subPktCt := toInt(bits[7:18])
			size = 18
			p.children = make([]packet, subPktCt)
			for i := 0; i < subPktCt; i++ {
				child, ct := parsePacket(bits[size:])
				size += ct
				p.children[i] = *child
			}
		case false:
			subPktBits := toInt(bits[7:22])
			size = 22
			p.children = make([]packet, 0)
			for subPktBits > 0 {
				child, ct := parsePacket(bits[size:])
				subPktBits -= ct
				size += ct
				p.children = append(p.children, *child)
			}
		}
	}
	return p, size
}

func toInt(bits []bool) int {
	n := 0
	for _, bit := range bits {
		if bit {
			n = 2*n + 1
		} else {
			n = 2 * n
		}
	}
	return n
}

func countVersions(p *packet) int {
	if p == nil {
		return 0
	}
	v := p.version
	for _, packet := range p.children {
		v += countVersions(&packet)
	}
	return v
}

func hexToBits(s string) []bool {
	bits := make([]bool, 0)
	for _, c := range s {
		switch c {
		case '0':
			bits = append(bits, false, false, false, false)
		case '1':
			bits = append(bits, false, false, false, true)
		case '2':
			bits = append(bits, false, false, true, false)
		case '3':
			bits = append(bits, false, false, true, true)
		case '4':
			bits = append(bits, false, true, false, false)
		case '5':
			bits = append(bits, false, true, false, true)
		case '6':
			bits = append(bits, false, true, true, false)
		case '7':
			bits = append(bits, false, true, true, true)
		case '8':
			bits = append(bits, true, false, false, false)
		case '9':
			bits = append(bits, true, false, false, true)
		case 'A':
			bits = append(bits, true, false, true, false)
		case 'B':
			bits = append(bits, true, false, true, true)
		case 'C':
			bits = append(bits, true, true, false, false)
		case 'D':
			bits = append(bits, true, true, false, true)
		case 'E':
			bits = append(bits, true, true, true, false)
		case 'F':
			bits = append(bits, true, true, true, true)

		}
	}
	return bits
}

func Solve2(r io.Reader) int {
	fmt.Println("Part 2 unimplemented")
	return 0
}
