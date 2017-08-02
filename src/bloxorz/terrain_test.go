package bloxorz

import (
	"bytes"
	"testing"
)

func TestArrayTerrainIsLegal(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString(`*.*E*`)
	buffer.WriteString(`*****`)
	buffer.WriteString(`*****`)
	buffer.WriteString(`S***.`)
	buffer.WriteString(`*.***`)
	arr := buffer.Bytes()

	// start and end irrelevant
	terrain := &ArrayTerrain{arr: arr, width: 5, start: nil, end: nil}

	var tests = []struct {
		block *Block
		legal bool
	}{
		{NewBlockUp(0, 0), true},
		{NewBlockUp(1, 0), false},
		{NewBlockUp(2, 0), true},
		{NewBlockUp(3, 0), true},
		{NewBlockUp(4, 0), true},
		{NewBlockUp(5, 0), false},
		{NewBlockUp(0, 1), true},
		{NewBlockUp(4, 1), true},
		{NewBlockUp(5, 1), false},
		{NewBlockUp(0, 3), true},
		{NewBlockUp(4, 3), false}, // dot on right border
		{NewBlockUp(0, 4), true},
		{NewBlockUp(0, 5), false},
		{NewBlockUp(4, 4), true},
		{NewBlockUp(4, 5), false}, // below last row
		{NewBlockUp(5, 4), false},
		{NewBlockDown(0, 0, 0, 1), true},
		{NewBlockDown(0, 0, 1, 0), false},
		{NewBlockDown(3, 4, 4, 4), true},
		{NewBlockDown(4, 3, 4, 4), false},
	}

	for _, c := range tests {
		block := c.block
		got := terrain.IsLegal(block)
		want := c.legal
		if got != want {
			t.Errorf("isLegal %s: got %t, wanted %t\n", block, got, want)
		}
	}
}
