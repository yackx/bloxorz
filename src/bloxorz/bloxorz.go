package bloxorz

import (
	"github.com/hishboy/gocommons/lang" // Q implementation
)

// Solve a basic Bloxorz terrain.
//
// Does not implement fancy terrains e.g. switches, soft squares, ...
//
// Each time a solution is found, it will be sent to the solution channel.
// The first solution is the (or one of the) optimal solution(s).
//
// Game:
// http://www.coolmath-games.com/0-bloxorz/
//
// Walkthrough:
// http://wiki.answers.com/Q/How_do_you_complete_all_the_levels_on_Bloxorz
func Solve(terrain Terrain, solution chan Path) {
	queue := lang.NewQueue()
	path := NewPath()
	start := terrain.Start()
	path.Add(*start)
	queue.Push(path)
	solveIter(terrain, queue, solution)
}

// Iterative solving using BFS.
// Valid solutions emitted on the solution channel.
func solveIter(terrain Terrain, queue *lang.Queue, solution chan Path) {
	for queue.Len() > 0 {

		// Take the first Path from the queue
		path := queue.Poll().(Path)

		// The last Block from the first path of the queue
		// We'll try to move the block to each of its neighbours
		block := path.Tail()

		// Going upwards, downwards, left and right
		directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

		for _, direction := range directions {
			dx := direction[0]
			dy := direction[1]
			neighbour := block.move(dx, dy)

			// Did we visit this new spot? Check on the current path
			visited := path.Contains(*neighbour)
			if !visited && terrain.IsLegal(neighbour) {
				// Not visited and legal, add it to path and push it to Q
				newPath := path.Clone()
				newPath.Add(*neighbour)
				queue.Push(newPath)
				if neighbour.Equals(terrain.End()) {
					// Found a solution
					solution <- newPath
				}
			}
		}
	}

	// Done
	solution <- nil
}
