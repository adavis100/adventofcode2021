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
		n    int
		want []int64
	}{
		{"1", template, 1, countChars("NCNBCHB")},
		{"2", template, 2, countChars("NBCCNBBBCBHCB")},
		{"3", template, 3, countChars("NBBBCNCCNBBNBNBBCHBHHBCHB")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := doIterations(test.s, test.n, rules)
			assert.Equal(t, test.want, got)
		})
	}
}

func countChars(s string) []int64 {
	counts := make([]int64, 26)
	for _, c := range s {
		counts[c-'A']++
	}
	return counts
}

func TestExample1(t *testing.T) {
	assert.Equal(t, int64(1588), Solve(strings.NewReader(ex), 10))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, int64(2188189693529), Solve(strings.NewReader(ex), 40))
}
