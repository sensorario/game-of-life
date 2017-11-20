package main

import "testing"
import "strconv"

func TestBoardCanContainAliveCell(t *testing.T) {
	board := Board{
		[]Cell{
			Cell{2, 2},
		},
	}

	if board.Contains(Cell{42, 42}) {
		t.Error("Cell should be dead")
	}

	if !board.Contains(Cell{2, 2}) {
		t.Error("Cell should be alive")
	}
}

func TestProvideNeightborsCoords(t *testing.T) {
	board := Board{
		[]Cell{
			Cell{0, 0},
			Cell{0, 1},
			Cell{0, 2},
			Cell{1, 0},
			Cell{1, 2},
			Cell{2, 0},
			Cell{2, 1},
			Cell{2, 2},
		},
	}

	coords := NeighborsCoords(Cell{1, 1})

	if false == board.EqualsTo(coords) {
		t.Error("Wrong neightbor found")
	}
}

func TestCountNeighborsAlive(t *testing.T) {
	board := Board{
		[]Cell{
			Cell{2, 2},
			Cell{0, 1},
			Cell{0, 42},
		},
	}

	neighborsAlive := board.NeighborsAlive(Cell{1, 1})
	if neighborsAlive != 2 {
		t.Error("Found " + strconv.Itoa(neighborsAlive) + " instead of 2")
	}

	neighborsAlive = board.NeighborsAlive(Cell{0, 43})
	if neighborsAlive != 1 {
		t.Error("Found " + strconv.Itoa(neighborsAlive) + " instead of 1")
	}
}

func TestExampleOfNewGeneration(t *testing.T) {
	before := Board{
		[]Cell{
			Cell{1, 0},
			Cell{1, 1},
			Cell{1, 2},
		},
	}

	after := Board{
		[]Cell{
			Cell{0, 1},
			Cell{1, 1},
			Cell{2, 1},
		},
	}

	nextG := before.NewGen()

	if !after.EqualsTo(nextG) {
		t.Error("Next gen should be different")
	}

	nextG = nextG.NewGen()
	before = Board{
		[]Cell{
			Cell{1, 0},
			Cell{1, 1},
			Cell{1, 2},
		},
	}

	if !before.EqualsTo(nextG) {
		t.Error("Next gen should be different")
	}
}
