package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var ex = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func TestInsertions(t *testing.T) {
	template, rules := loadInput(strings.NewReader(ex))

	tests := []struct {
		name string
		s    string
		want string
	}{
		{"1", template, "NCNBCHB"},
		{"2", "NCNBCHB", "NBCCNBBBCBHCB"},
		{"3", "NBCCNBBBCBHCB", "NBBBCNCCNBBNBNBBCHBHHBCHB"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, doInsertion(test.s, rules))
		})
	}
}

func TestExample1(t *testing.T) {
	assert.Equal(t, 1588, Solve1(strings.NewReader(ex)))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, 2188189693529, Solve2(strings.NewReader(ex)))
}
