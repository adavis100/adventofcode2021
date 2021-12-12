package main

import (
	"strings"
	"testing"
)

func TestRun2(t *testing.T) {
	r := strings.NewReader(`forward 5
down 5
forward 8
up 3
down 8
forward 2`)
	hpos, depth := run2(loadMoves(r))
	if hpos != 15 {
		t.Errorf("wanted hpos %d got %d", 15, hpos)
	}
	if depth != 60 {
		t.Errorf("wanted depth %d got %d", 60, depth)
	}
}
