package bloxorz

import (
	"testing"
)

func TestBlockEquals(t *testing.T) {
	one := NewBlockDown(1, 2, 1, 3)
	oneAgain := NewBlockDown(1, 2, 1, 3)
	two := NewBlockDown(1, 2, 2, 2)

	if !one.Equals(oneAgain) {
		t.Error("Equals: one should equal oneAgain")
	}

	if one.Equals(two) {
		t.Error("Equals: one should not equal two")
	}

	fiveZeroFiveOne := NewBlockDown(5, 0, 5, 1)
	fiveOne := NewBlockUp(5, 1)
	if fiveZeroFiveOne.Equals(fiveOne) {
		t.Errorf("Equals: %s should not equal %s\n", fiveZeroFiveOne, fiveOne)
	}
}

func TestIsUp(t *testing.T) {
	var tests = []struct {
		block *Block
		isUp  bool
	}{
		{NewBlockUp(0, 0), true},
		{NewBlockDown(1, 0, 1, 0), true},
		{NewBlockDown(1, 0, 2, 0), false},
	}

	for _, c := range tests {
		isUp := c.block.isUp()
		if isUp != c.isUp {
			t.Errorf("IsUp() %s should be %t, got %t", c.block, c.isUp, isUp)
		}
	}
}

func TestMove(t *testing.T) {
	var tests = []struct {
		description string
		block       *Block
		dx, dy      int
		want        *Block
	}{
		{"up falls east", NewBlockUp(0, 0), 1, 0, NewBlockDown(1, 0, 2, 0)},
		{"up falls west", NewBlockUp(0, 0), -1, 0, NewBlockDown(-2, 0, -1, 0)},
		{"up falls north", NewBlockUp(0, 0), 0, -1, NewBlockDown(0, -2, 0, -1)},
		{"up falls south", NewBlockUp(0, 0), 0, 1, NewBlockDown(0, 1, 0, 2)},
		{"down NS rises north", NewBlockDown(0, 0, 0, 1), 0, -1, NewBlockUp(0, -1)},
		{"down NS rises south", NewBlockDown(0, 0, 0, 1), 0, 1, NewBlockUp(0, 2)},
		{"down EW rises west", NewBlockDown(0, 0, 1, 0), -1, 0, NewBlockUp(-1, 0)},
		{"down EW rises east", NewBlockDown(0, 0, 1, 0), 1, 0, NewBlockUp(2, 0)},
		{"down NS rolls west", NewBlockDown(0, 0, 0, 1), -1, 0, NewBlockDown(-1, 0, -1, 1)},
		{"down NS rolls east", NewBlockDown(0, 0, 0, 1), 1, 0, NewBlockDown(1, 0, 1, 1)},
		{"down EW rolls north", NewBlockDown(0, 0, 1, 0), 0, -1, NewBlockDown(0, -1, 1, -1)},
		{"down EW rolls south", NewBlockDown(0, 0, 1, 0), 0, 1, NewBlockDown(0, 1, 1, 1)},
	}

	for _, c := range tests {
		block := c.block
		got := block.move(c.dx, c.dy)
		if !got.Equals(c.want) {
			t.Errorf(
				"Invalid move '%s' from %s to %s using vector (x=%d, y=%d). Expected %s",
				c.description, block, got, c.dx, c.dy, c.want)
		}
	}
}
