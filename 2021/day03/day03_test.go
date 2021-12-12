package main

import (
	"strings"
	"testing"
)

var in string = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`

func TestExamplePart1(t *testing.T) {

	gamma, epsilon := part1MostLeast(loadLines(strings.NewReader(in)))
	if gamma != "10110" {
		t.Errorf("wanted gamma %s got %s", "10110", gamma)
	}
	if epsilon != "01001" {
		t.Errorf("wanted epsilon %s got %s", "01001", epsilon)
	}
}

func TestExamplePart2(t *testing.T) {
	most, least := part2MostLeast(loadLines(strings.NewReader(in)))
	if most != "10111" {
		t.Errorf("wanted most %s got %s", "10111", most)
	}
	if least != "01010" {
		t.Errorf("wanted least %s got %s", "01010", least)
	}
}
