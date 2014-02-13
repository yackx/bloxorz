package bloxorz

import (
	"bytes"
	"fmt"
	"testing"
)

/****** TERRAIN BUILDERS ******/

// Build a super simple terrain
func buildTerrain1() *ArrayTerrain {
	arr := []byte(`S**E`)

	start := NewBlockUp(0, 0)
	end := NewBlockUp(3, 0)

	return NewArrayTerrain(start, end, arr, 4)
}

// Build a simple sample terrain
func buildTerrain2() *ArrayTerrain {

	var buffer bytes.Buffer
	buffer.WriteString(`***E*`)
	buffer.WriteString(`*****`)
	buffer.WriteString(`*****`)
	buffer.WriteString(`S****`)
	arr := buffer.Bytes()

	start := NewBlockUp(0, 3)
	end := NewBlockUp(3, 0)

	return NewArrayTerrain(start, end, arr, 5)
}

// Build an unsolvable terrain
func buildUnsolvableTerrain() *ArrayTerrain {

	arr := []byte(`..S..E..`)

	start := NewBlockUp(2, 0)
	end := NewBlockUp(5, 0)

	return NewArrayTerrain(start, end, arr, 8)
}

// Build the level 3 terrain
func buildLevel3() *ArrayTerrain {
	var buffer bytes.Buffer
	//                  012345678901234
	buffer.WriteString(`......*******..`) // 0
	buffer.WriteString(`****..***..**..`) // 1
	buffer.WriteString(`*********..****`) // 2
	buffer.WriteString(`****.......**E*`) // 3
	buffer.WriteString(`*S**.......****`) // 4
	buffer.WriteString(`****........***`) // 5
	arr := buffer.Bytes()

	start := NewBlockUp(1, 4)
	end := NewBlockUp(13, 3)

	return NewArrayTerrain(start, end, arr, 15)
}

// Build the level 6 terrain
func buildLevel6() *ArrayTerrain {
	var buffer bytes.Buffer
	//                  012345678901234
	buffer.WriteString(`.....******....`) // 0
	buffer.WriteString(`.....*..***....`) // 1
	buffer.WriteString(`.....*..*****..`) // 2
	buffer.WriteString(`S*****.....****`) // 3
	buffer.WriteString(`....***....**E*`) // 4
	buffer.WriteString(`....***.....***`) // 5
	buffer.WriteString(`......*..**....`) // 6
	buffer.WriteString(`......*****....`) // 7
	buffer.WriteString(`......*****....`) // 8
	buffer.WriteString(`.......***.....`) // 9
	arr := buffer.Bytes()

	start := NewBlockUp(0, 3)
	end := NewBlockUp(13, 4)

	return NewArrayTerrain(start, end, arr, 15)
}

/****** SOLVER TESTS ******/

func testSolvableTerrain(t *testing.T, terrain Terrain, name string) {
	solutionChannel := make(chan Path)
	go Solve(terrain, solutionChannel)

	solution := <-solutionChannel
	fmt.Printf("%s first solution in %d moves: %s\n",
		name, len(solution)-1, solution.String())
	/*
		var i int
		for i = 0; solution != nil; i++ {
			solution = <-solutionChannel
		}
		fmt.Printf("terrain2 has %d other solutions\n", i)
	*/
}

func TestTerrain1(t *testing.T) {
	testSolvableTerrain(t, buildTerrain1(), "terrain1")
}

func TestTerrain2(t *testing.T) {
	testSolvableTerrain(t, buildTerrain2(), "terrain2")
}

func TestLevel3(t *testing.T) {
	testSolvableTerrain(t, buildLevel3(), "level3")
}

func TestLevel6(t *testing.T) {
	testSolvableTerrain(t, buildLevel6(), "level6")
}

func TestInfiniteTerrain(t *testing.T) {
	start := NewBlockUp(0, 0)
	end := NewBlockUp(5, 1)
	infinite := NewInfiniteTerrain(start, end)

	testSolvableTerrain(t, infinite, "infinite")
}

func TestTerrain1ActualSolution(t *testing.T) {
	var terrain = buildTerrain1()

	solutionChannel := make(chan Path)
	go Solve(terrain, solutionChannel)

	solution := <-solutionChannel

	expectedPath := NewPath()
	expectedPath.Add(*NewBlockUp(0, 0))
	expectedPath.Add(*NewBlockDown(1, 0, 2, 0))
	expectedPath.Add(*NewBlockUp(3, 0))
	if len(expectedPath) != len(solution) {
		t.Errorf("should have the same size. Got %d, wanted %d\n", len(expectedPath), len(solution))
		return
	}

	for i := 0; i < len(expectedPath); i++ {
		if !expectedPath[i].Equals(&solution[i]) {
			t.Errorf(
				"incorrect solution at step %d. Got %s, expected %s",
				i, solution[i].String(), expectedPath[i].String())
			return
		}
	}
	fmt.Println("terrain1 detailed solution ok")

	end := <-solutionChannel
	if end != nil {
		t.Error("should not have found another solution", end)
	}
	fmt.Println("terrain1 only one solution ok")

}

func TestUnsolvableTerrain(t *testing.T) {
	var unsolvable = buildUnsolvableTerrain()

	solutionChannel := make(chan Path)
	go Solve(unsolvable, solutionChannel)

	noSolution := <-solutionChannel
	close(solutionChannel)

	if noSolution != nil {
		t.Error("should not have found a solution", noSolution)
		return
	}

	fmt.Println("unsolvable terrain has no solution")
}
