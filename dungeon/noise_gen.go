package dungeon

import (
	"fmt"
	"time"

	"github.com/KEINOS/go-noise"
)

type NoiseDungeon struct {
	Grid [][]float32
}

func NewNoiseDungeon(width, height int) *NoiseDungeon {
	grid := make([][]float32, height)
	for i := range grid {
		grid[i] = make([]float32, width)
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}
	return &NoiseDungeon{Grid: grid}
}

func (d *NoiseDungeon) Print() {
	for i := range d.Grid {
		for j := range d.Grid[i] {
			switch {
			case d.Grid[i][j] < -(2.0 / 3):
				fmt.Print("  ")
			case d.Grid[i][j] < -(1.0 / 3):
				fmt.Print(" .")
			case d.Grid[i][j] < 0:
				fmt.Print(" ░")
			case d.Grid[i][j] < 1.0/3:
				fmt.Print(" ▒")
			case d.Grid[i][j] < 2.0/3:
				fmt.Print(" ▓")
			default:
				fmt.Print(" █")
			}
		}
		fmt.Println()
	}
}

func CreateNoiseDungeon(width, height int) *NoiseDungeon {
	d := NewNoiseDungeon(width, height)
	err := d.Generate()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	d.Print()
	return d
}

func (d *NoiseDungeon) Generate() error {
	// Parameters to play with
	gen, err := noise.New(noise.OpenSimplex, time.Now().UnixNano())
	if err != nil {
		return err
	}
	for y := range d.Grid {
		for x := range d.Grid[y] {
			// Parameters to play with
			val := gen.Eval32(float32(x)/60, float32(y)/60)
			d.Grid[y][x] = val
		}
	}
	return nil

}
