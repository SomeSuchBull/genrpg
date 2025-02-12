package dungeon

import "math/rand"

type CADungeon struct {
	Grid [][]int
}

func NewCADungeon(width, height int) *CADungeon {
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
		for j := range grid[i] {
			if rand.Intn(100) < 45 {
				grid[i][j] = 1
			} else {
				grid[i][j] = 0
			}
			// grid[i][j] = rand.Intn(2)
		}
	}
	return &CADungeon{Grid: grid}
}

func (d *CADungeon) Print() {
	for i := range d.Grid {
		for j := range d.Grid[i] {
			if d.Grid[i][j] == 0 {
				print(" █")
			} else {
				print(" .")
			}
		}
		println()
	}
}

func CreateCADungeon(width, height int) *CADungeon {
	d := NewCADungeon(width, height)
	d.Print()
	for i := 0; i < 7; i++ {
		println()
		d.Step()
		d.Print()
	}
	return d
}

func (d *CADungeon) Step() {
	newGrid := make([][]int, len(d.Grid))
	for i := range d.Grid {
		newGrid[i] = make([]int, len(d.Grid[i]))
		for j := range d.Grid[i] {
			// if i == 0 || i == len(d.Grid)-1 || j == 0 || j == len(d.Grid[i])-1 {
			// 	newGrid[i][j] = 1
			// 	continue
			// }
			newGrid[i][j] = d.Grid[i][j]
		}
	}
	// for i := range d.Grid {
	for i := 1; i < len(d.Grid)-1; i++ {
		// for j := range d.Grid[i] {
		for j := 1; j < len(d.Grid[i])-1; j++ {
			neighbors := 0
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}
					ni := i + x
					nj := j + y
					if ni < 0 || ni >= len(d.Grid) || nj < 0 || nj >= len(d.Grid[i]) {
						continue
					}
					neighbors += d.Grid[ni][nj]
				}
			}
			if d.Grid[i][j] == 0 && neighbors > 4 {
				newGrid[i][j] = 1
			} else if d.Grid[i][j] == 0 && neighbors == 0 {
				newGrid[i][j] = 1
			} else if d.Grid[i][j] == 1 && neighbors < 4 {
				newGrid[i][j] = 0
			} else {
				newGrid[i][j] = d.Grid[i][j]
			}
		}
	}
	d.Grid = newGrid
}
