package bloxorz

import (
	"testing"
)

func TestAddAndContains(t *testing.T) {
	path := NewPath()
	start := NewBlockUp(0, 1)

	if path.Contains(*start) {
		t.Error("path should not contain start")
	}

	path.Add(*start)
	if !path.Contains(*start) {
		t.Error("path should contain start", path)
	}

}

func TestTail(t *testing.T) {
	path := NewPath()
	if path.Tail() != nil {
		t.Error("empty list tail should be nil")
	}

	block := NewBlockUp(0, 1)
	path.Add(*block)
	tail := path.Tail()
	if !tail.Equals(block) {
		t.Error("list should return tail", block, tail)
	}
}

func TestClone(t *testing.T) {
	path := NewPath()
	block := NewBlockUp(2, 3)
	path.Add(*block)

	clone := path.Clone()
	if !clone.Contains(*block) {
		t.Error("Cloned path should contain the block")
	}

	otherBlock := NewBlockUp(4, 4)
	path.Add(*otherBlock)
	if clone.Contains(*otherBlock) {
		t.Error("Cloned path should not contain the otherBlock")
	}

	thirdBlock := NewBlockUp(1, 1)
	clone.Add(*thirdBlock)
	if path.Contains(*thirdBlock) {
		t.Error("Original path should not contain thirdBlock")
	}
}
