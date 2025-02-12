package dungeon

import (
	"math/rand"
)

type DrunkDungeon struct {
	Grid [][]int
}

func NewDrunkDungeon(width, height int) *DrunkDungeon {
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}
	return &DrunkDungeon{Grid: grid}
}

func (d *DrunkDungeon) Print() {
	for i := range d.Grid {
		for j := range d.Grid[i] {
			if d.Grid[i][j] == 0 {
				print(" .")
			} else {
				print(" █")
			}
		}
		println()
	}
}

func CreateDrunkDungeon(width, height int) *DrunkDungeon {
	d := NewDrunkDungeon(width, height)
	d.Walk()
	d.Print()
	return d
}

func (d *DrunkDungeon) Walk() {
	// x := len(d.Grid[0]) / 2
	// y := len(d.Grid) / 2
	x := rand.Intn(len(d.Grid[0]))
	y := rand.Intn(len(d.Grid))
	d.Grid[y][x] = 1
	var breaking bool
	steps := len(d.Grid) * len(d.Grid[0])
	for i := 0; i < steps; i++ {
		shuffledDirections := rand.Perm(len(directions))
		for tries, j := range shuffledDirections {
			dx, dy := directions[j].X, directions[j].Y
			if x+dx < 0 || x+dx >= len(d.Grid[0]) || y+dy < 0 || y+dy >= len(d.Grid) {
				if tries == 3 {
					x = len(d.Grid[0]) / 2
					y = len(d.Grid) / 2
				}
				continue
			} else if d.Grid[y+dy][x+dx] == 1 {
				if tries == 3 {
					x = len(d.Grid[0]) / 2
					y = len(d.Grid) / 2
				}
				continue
			} else {
				x += dx
				y += dy
			}
			d.Grid[y][x] = 1
			break
		}
		if breaking {
			break
		}
	}
}
