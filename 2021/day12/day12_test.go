package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var ex1 = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var ex2 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

var ex3 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

func TestLoadsGraph(t *testing.T) {

	g := LoadGraph(strings.NewReader(ex1))
	assert.Equal(t, []string{"A", "b"}, g["start"])
	assert.Equal(t, []string{"start", "A", "d", "end"}, g["b"])
}

func TestSolves1(t *testing.T) {

	assert.Equal(t, 10, Solve1(strings.NewReader(ex1)))

	assert.Equal(t, 19, Solve1(strings.NewReader(ex2)))
}

func TestSolves2(t *testing.T) {

	assert.Equal(t, 36, Solve2(strings.NewReader(ex1)))
	assert.Equal(t, 103, Solve2(strings.NewReader(ex2)))
	assert.Equal(t, 3509, Solve2(strings.NewReader(ex3)))
}
