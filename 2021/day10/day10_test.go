package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetErrorScore(t *testing.T) {
	tests := []struct {
		l    string
		want int
	}{
		{"{([(<{}[<>[]}>{[]{[(<()>", 1197},
		{"[[<[([]))<([[{}[[()]]]", 3},
		{"[{[{({}]{}}([{[{{{}}([]", 57},
		{"<{([([[(<>()){}]>(<<{{", 25137},
	}
	for _, test := range tests {
		t.Run(test.l, func(t *testing.T) {
			assert.Equal(t, test.want, getErrorScore(test.l))
		})
	}
}

var in = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestSolve1(t *testing.T) {
	assert.Equal(t, 26397, Solve1(strings.NewReader(in)))
}

func TestGetRemainderScore(t *testing.T) {
	tests := []struct {
		l    string
		want int
	}{
		{"[({(<(())[]>[[{[]{<()<>>", 288957},
		{"[(()[<>])]({[<{<<[]>>(", 5566},
	}
	for _, test := range tests {
		t.Run(test.l, func(t *testing.T) {
			assert.Equal(t, test.want, getScoreForIncompleteLine(test.l))
		})
	}
}

func TestSolve2(t *testing.T) {
	assert.Equal(t, 288957, Solve2(strings.NewReader(in)))
}
