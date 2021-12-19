package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCountsNumsInLine(t *testing.T) {

	n := countUniqueNumsInLine("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe")
	assert.Equal(t, 2, n)
}

var in = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

func TestSolvesExample1(t *testing.T) {
	assert.Equal(t, 26, countNums(strings.NewReader(in), true))
}

func TestGetsNumsInLine(t *testing.T) {
	tests := []struct {
		name string
		line string
		want int
	}{
		{"1", "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | ab", 1},
		{"7", "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | dab", 7},
		{"4", "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | eafb", 4},
		{"8", "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | acedgfb", 8},
		{"3", "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | fbcad", 3},
		{"5", "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfbe", 5},
		{"2", "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | gcdfa", 2},
		{"6", "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfgeb", 6},
		{"9", "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cefabd", 9},
		{"0", "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cagedb", 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := lineToNum(test.line)
			assert.Equal(t, test.want, n)
		})
	}
}

func TestSolvesExample2(t *testing.T) {

	assert.Equal(t, 61229, countNums(strings.NewReader(in), false))
}
