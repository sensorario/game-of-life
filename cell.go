package main

type Cell struct {
	x int
	y int
}

type Board struct {
	cells []Cell
}

func (b *Board) Contains(c Cell) bool {
	for _, cell := range b.cells {
		if cell.x == c.x && cell.y == c.y {
			return true
		}
	}

	return false
}

func (b Board) Add(c Cell) Board {
	nextBoard := b
	nextBoard.cells = append(nextBoard.cells, c)
	return nextBoard
}

func NeighborsCoords(c Cell) Board {
	return Board{
		[]Cell{
			Cell{c.x - 1, c.y - 1},
			Cell{c.x - 1, c.y},
			Cell{c.x - 1, c.y + 1},
			Cell{c.x, c.y - 1},
			Cell{c.x, c.y + 1},
			Cell{c.x + 1, c.y - 1},
			Cell{c.x + 1, c.y},
			Cell{c.x + 1, c.y + 1},
		},
	}
}

func (b *Board) NeighborsAlive(c Cell) int {
	num := 0
	for _, cell := range NeighborsCoords(c).cells {
		if b.Contains(cell) {
			num++
		}
	}
	return num
}

func (b *Board) NewGen() Board {
	newBoard := Board{}

	for _, cell := range b.cells {
		n := b.NeighborsAlive(cell)
		if n == 2 || n == 3 {
			newBoard = newBoard.Add(cell)
		}

		for _, neighbor := range NeighborsCoords(cell).cells {
			if b.NeighborsAlive(neighbor) == 3 {
				if !newBoard.Contains(neighbor) {
					newBoard = newBoard.Add(neighbor)
				}
			}
		}
	}

	return newBoard
}

func (a *Board) EqualsTo(b Board) bool {
	if a.cells == nil {
		return b.cells == nil
	}

	if len(a.cells) != len(b.cells) {
		return false
	}

	for i := range a.cells {
		if !b.Contains(a.cells[i]) {
			return false
		}
	}

	return true
}
