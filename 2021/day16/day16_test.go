package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConvertsFirstExampleToBits(t *testing.T) {
	s := "D2FE28"
	bits := hexToBits(s)
	// 110100101111111000101000
	want := []bool{true, true, false, true, false, false, true, false, true, true, true, true, true, true, true, false, false, false, true, false, true, false, false, false}
	assert.Equal(t, want, bits)
}

func TestBitsToInt(t *testing.T) {
	bits := []bool{true, true, false, true}
	assert.Equal(t, 13, toInt(bits))
}

func TestParseExample1(t *testing.T) {
	s := "D2FE28"
	p, size := parsePacket(hexToBits(s))
	assert.Equal(t, 6, p.version, "version")
	assert.Equal(t, 4, p.pktType, "type")
	assert.Equal(t, 2021, p.num, "num")
	assert.Equal(t, 21, size)
}

func TestParseExample2(t *testing.T) {
	s := "38006F45291200"
	p, size := parsePacket(hexToBits(s))
	assert.Equal(t, 1, p.version, "version")
	assert.Equal(t, 6, p.pktType, "type")
	assert.Equal(t, 2, len(p.children), "children")
	assert.Equal(t, 49, size)
}

func TestParseExample3(t *testing.T) {
	s := "EE00D40C823060"
	p, size := parsePacket(hexToBits(s))
	assert.Equal(t, 7, p.version, "version")
	assert.Equal(t, 3, p.pktType, "type")
	assert.Equal(t, 3, len(p.children), "children")
	assert.Equal(t, 51, size)
}

func TestGetsVersions(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"D2FE28", 6},
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			assert.Equal(t, test.want, Solve1(strings.NewReader(test.s)))
		})
	}
}

func TestEvalsPackets(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			assert.Equal(t, test.want, Solve2(strings.NewReader(test.s)))
		})
	}
}
