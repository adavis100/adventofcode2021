package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLoadsGraph(t *testing.T) {
	in := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
	g := LoadGraph(strings.NewReader(in))
	assert.Equal(t, []string{"A", "b"}, g["start"])
	assert.Equal(t, []string{"start", "A", "d", "end"}, g["b"])
}

func TestSolves1(t *testing.T) {
	in := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
	assert.Equal(t, 10, Solve1(strings.NewReader(in)))

	in = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`
	assert.Equal(t, 19, Solve1(strings.NewReader(in)))
}
