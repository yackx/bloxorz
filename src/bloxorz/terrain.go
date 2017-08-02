package bloxorz

// Interface for terrains.
type Terrain interface {
	IsLegal(*Block) bool
	Start() *Block
	End() *Block
}

/*** ARRAY TERRAIN */

// A terrain represented by an array of bytes.
type ArrayTerrain struct {
	start, end *Block
	arr        []byte
	width      int
}

// Create a new ArrayTerrain.
func NewArrayTerrain(start, end *Block, arr []byte, width int) *ArrayTerrain {
	return &ArrayTerrain{start: start, end: end, arr: arr, width: width}
}

// Test if the given Block is legal on this terrain.
//
// The Block is considered legal if the square is not
// represented by a dot ('.').
func (at *ArrayTerrain) IsLegal(b *Block) bool {
	if b.ax < 0 || b.ay < 0 || b.bx >= at.width || b.bx+b.by*at.width >= len(at.arr) {
		return false
	}

	squareA := at.arr[b.ax+at.width*b.ay]
	squareB := at.arr[b.bx+at.width*b.by]

	return squareA != '.' && squareB != '.'
}

// Starting Block on this terrain.
func (at *ArrayTerrain) Start() *Block {
	return at.start
}

// End or exit Block on this terrain.
func (at *ArrayTerrain) End() *Block {
	return at.end
}

/*** INFINITE TERRAIN ***/

// An infinite terrain, where all squares legal.
type InifiniteTerrain struct {
	start, end *Block
}

// Create a new infinite terrain.
func NewInfiniteTerrain(start, end *Block) *InifiniteTerrain {
	return &InifiniteTerrain{start: start, end: end}
}

// Test if the given Block is legal on this terrain.
// Always returns true.
func (it *InifiniteTerrain) IsLegal(b *Block) bool {
	return true
}

// Starting Block on this terrain.
func (it *InifiniteTerrain) Start() *Block {
	return it.start
}

// End or exit Block on this terrain.
func (it *InifiniteTerrain) End() *Block {
	return it.end
}
