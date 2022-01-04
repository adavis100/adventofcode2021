package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLoadsNum(t *testing.T) {
	s := "[[1,2],3]"
	snailnum := loadSnailNum(s)
	assert.Equal(t, 1, snailnum.val)
	assert.Nil(t, snailnum.prev)
	assert.Equal(t, 2, snailnum.nest)
	assert.Equal(t, 2, snailnum.next.val)
	assert.Equal(t, 2, snailnum.next.nest)
	assert.Equal(t, 1, snailnum.next.prev.val)
	assert.Equal(t, 3, snailnum.next.next.val)
	assert.Equal(t, 1, snailnum.next.next.nest)
	assert.Equal(t, 2, snailnum.next.next.prev.val)
	assert.Nil(t, snailnum.next.next.next)
}

func TestExplodes(t *testing.T) {
	tests := []struct {
		s    string
		want *snailnum
	}{
		{"[[[[[9,8],1],2],3],4]", loadSnailNum("[[[[0,9],2],3],4]")},
		{"[7,[6,[5,[4,[3,2]]]]]", loadSnailNum("[7,[6,[5,[7,0]]]]")},
		{"[[6,[5,[4,[3,2]]]],1]", loadSnailNum("[[6,[5,[7,0]]],3]")},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			sn := loadSnailNum(test.s)
			explode(sn)
			assert.Equal(t, test.want, sn)
		})
	}
}

func TestSplits(t *testing.T) {
	tests := []struct {
		s    string
		want *snailnum
	}{
		{"[[[[0,7],4],[15,[0,13]]],[1,1]]", loadSnailNum("[[[[0,7],4],[[7,8],[0,13]]],[1,1]]")},
		{"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", loadSnailNum("[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]")},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			sn := loadSnailNum(test.s)
			split(sn)
			assert.Equal(t, test.want, sn)
		})
	}
}

func TestAddition(t *testing.T) {

	tests := []struct {
		name string
		s1   *snailnum
		s2   *snailnum
		want *snailnum
	}{
		{`"[[[[4,3],4],4],[7,[[8,4],9]]]" + "[1,1]`, loadSnailNum("[[[[4,3],4],4],[7,[[8,4],9]]]"), loadSnailNum("[1,1]"), loadSnailNum("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sn := add(test.s1, test.s2)
			assert.Equal(t, test.want, sn)
		})
	}
}

var ex = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

func TestSolve1(t *testing.T) {
	assert.Equal(t, 4140, Solve1(strings.NewReader(ex)))
}

func TestGetMagnitude(t *testing.T) {
	tests := []struct {
		name string
		sn   *snailnum
		want int
	}{
		{"[[1,2],[[3,4],5]]", loadSnailNum("[[1,2],[[3,4],5]]"), 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", loadSnailNum("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"), 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", loadSnailNum("[[[[1,1],[2,2]],[3,3]],[4,4]]"), 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", loadSnailNum("[[[[3,0],[5,3]],[4,4]],[5,5]]"), 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", loadSnailNum("[[[[5,0],[7,4]],[5,5]],[6,6]]"), 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", loadSnailNum("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"), 3488},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getMagnitude(tt.sn, 0, 0), "getMagnitude(%v)", tt.sn)
		})
	}
}

func TestSolve2(t *testing.T) {
	assert.Equal(t, 3993, Solve2(strings.NewReader(ex)))
}
