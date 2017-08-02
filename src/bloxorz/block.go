package bloxorz

import (
	"fmt"
)

// A bloxorz Block is represented by 2 (x, y) pairs.
type Block struct {
	ax, ay, bx, by int
}

// Create a new Block lying down.
//
// The 'a' part must always be the left-most or the upper-most part
// of the block.
//
// There are no internal consistency checks
// e.g. left-most/upper-most rule or (ax, ay) adjacent to (bx, by)
func NewBlockDown(ax, ay, bx, by int) *Block {
	return &Block{ax: ax, ay: ay, bx: bx, by: by}
}

// Create a new Block standing up.
//
// If the Block is up, bx and by are set to ax and ay
// values respectively.
func NewBlockUp(x, y int) *Block {
	return &Block{ax: x, ay: y, bx: x, by: y}
}

// Is this Block up?
func (block Block) isUp() bool {
	return block.ax == block.bx && block.ay == block.by
}

// Create a new Block moved relative to this Block
// according to the given vectors.
//
// dx and dy are expected to contain either 0, 1 or -1,
// with exactly one of them equal to either 1 or -1.
func (block Block) move(dx, dy int) *Block {
	if block.isUp() {
		if dx == 1 || dy == 1 {
			// Falling east or south: 'b' part takes double delta
			return NewBlockDown(
				block.ax+dx, block.ay+dy,
				block.bx+2*dx, block.by+2*dy)
		}
		// Falling west or north: 'a' part takes double delta
		return NewBlockDown(
			block.ax+2*dx, block.ay+2*dy, block.bx+dx, block.by+dy)
	}
	// Block is down

	isNorthSouth := block.ax == block.bx

	if (isNorthSouth && dy == 0) || (!isNorthSouth && dx == 0) {
		// Roll: add delta to both block parts
		return NewBlockDown(
			block.ax+dx, block.ay+dy, block.bx+dx, block.by+dy)
	}

	if (isNorthSouth && dy == -1) || (!isNorthSouth && dx == -1) {
		// Rising north or west: use 'a' part
		return NewBlockUp(block.ax+dx, block.ay+dy)
	}

	if (isNorthSouth && dy == 1) || (!isNorthSouth && dx == 1) {
		// Rising east or south: use 'b' part
		return NewBlockUp(block.bx+dx, block.by+dy)
	}

	panic(fmt.Sprintf("Block.move: unreachable", block, dx, dy))
}

// Is this Block equal to that other Block?
// Blocks are considered equal (on the same square or squares)
// if all parts have the same values.
func (block Block) Equals(other *Block) bool {
	return block.ax == other.ax && block.ay == other.ay &&
		block.bx == other.bx && block.by == other.by
}

// Human-readable representation of this block.
func (block Block) String() string {
	if block.isUp() {
		return fmt.Sprintf("(%d, %d)", block.ax, block.ay)
	}
	return fmt.Sprintf("(%d, %d)-(%d, %d)", block.ax, block.ay, block.bx, block.by)
}
