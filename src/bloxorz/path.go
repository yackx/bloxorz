package bloxorz

import "bytes"

// A Path is an array of Blocks.
type Path []Block

// Create a new empty Path.
func NewPath() Path {
	var path []Block
	return path
}

// Does this Path contain the given Block?
func (path Path) Contains(b Block) bool {
	for _, block := range path {
		if block.Equals(&b) {
			return true
		}
	}
	return false
}

// Add a Block at the end of this Path.
func (path *Path) Add(b Block) {
	*path = append(*path, b)
}

// Return the last Block of this Path.
func (path *Path) Tail() *Block {
	p := *path
	if len(p) == 0 {
		return nil
	}
	block := p[len(p)-1]
	return &block
}

// Returns a clone of this Path.
func (path Path) Clone() Path {
	clone := make([]Block, len(path))
	copy(clone, path)
	return clone
}

// Human-readable representation of this Path.
func (path Path) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("[")
	first := true

	for _, block := range path {
		if !first {
			buffer.WriteString("; ")
		}
		first = false
		buffer.WriteString(block.String())
	}
	buffer.WriteString("]")

	return buffer.String()
}
