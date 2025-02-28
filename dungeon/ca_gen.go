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
				print(" .")
			} else {
				print(" █")
			}
		}
		println()
	}
}

// func CreateCADungeon(width, height int) *CADungeon {
// 	d := NewCADungeon(width, height)
// 	d.Print()
// 	for i := 0; i < 7; i++ {
// 		println()
// 		d.Step()
// 		d.Print()
// 	}
// 	return d
// }

func (d *CADungeon) Step() {
	newGrid := make([][]int, len(d.Grid))
	for i := range d.Grid {
		newGrid[i] = make([]int, len(d.Grid[i]))
		for j := range d.Grid[i] {
			if i == 0 || i == len(d.Grid)-1 || j == 0 || j == len(d.Grid[i])-1 {
				newGrid[i][j] = 1
				continue
			}
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
				// } else if d.Grid[i][j] == 0 && neighbors == 0 {
				// 	newGrid[i][j] = 1
			} else if d.Grid[i][j] == 1 && neighbors < 4 {
				newGrid[i][j] = 0
			} else {
				newGrid[i][j] = d.Grid[i][j]
			}
		}
	}
	d.Grid = newGrid
}

// ...existing code...

// ConnectCaves connects all disconnected cave areas with natural-looking tunnels
func (d *CADungeon) ConnectCaves() {
	// First, identify all disconnected regions
	regions := d.findRegions()
	if len(regions) <= 1 {
		return // All caves are already connected
	}

	// Connect each region to the next one
	for i := 0; i < len(regions)-1; i++ {
		// Find the closest points between regions
		p1, p2 := d.findClosestPoints(regions[i], regions[i+1])
		// Connect these points with a winding path
		d.createTunnel(p1, p2)
	}
}

// findRegions uses flood fill to identify separate cave areas
func (d *CADungeon) findRegions() [][]Point {
	height := len(d.Grid)
	width := len(d.Grid[0])

	// Create a visited grid
	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	var regions [][]Point

	// Traverse the grid to find unvisited cave cells
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// If this is a cave cell and not visited yet
			if d.Grid[i][j] == 0 && !visited[i][j] {
				// Start a flood fill from this cell
				region := d.floodFill(i, j, visited)
				if len(region) > 0 { // Only consider regions of some size
					regions = append(regions, region)
				}
			}
		}
	}

	return regions
}

// floodFill returns all connected cells in a region
func (d *CADungeon) floodFill(startI, startJ int, visited [][]bool) []Point {
	var region []Point
	var queue []Point

	// Start flood fill from the given position
	queue = append(queue, Point{startI, startJ})
	visited[startI][startJ] = true

	// Define the four cardinal directions
	dirs := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		// Pop the first element
		curr := queue[0]
		queue = queue[1:]
		region = append(region, curr)

		// Check all adjacent cells
		for _, dir := range dirs {
			ni, nj := curr.X+dir.X, curr.Y+dir.Y

			// Check bounds
			if ni < 0 || ni >= len(d.Grid) || nj < 0 || nj >= len(d.Grid[0]) {
				continue
			}

			// If it's a cave cell and not visited
			if d.Grid[ni][nj] == 0 && !visited[ni][nj] {
				visited[ni][nj] = true
				queue = append(queue, Point{ni, nj})
			}
		}
	}

	return region
}

// findClosestPoints finds the closest points between two regions
func (d *CADungeon) findClosestPoints(region1, region2 []Point) (Point, Point) {
	minDist := float64(len(d.Grid) * len(d.Grid[0])) // Start with max possible distance
	var closest1, closest2 Point

	for _, p1 := range region1 {
		for _, p2 := range region2 {
			dist := distance(p1, p2)
			if dist < minDist {
				minDist = dist
				closest1, closest2 = p1, p2
			}
		}
	}

	return closest1, closest2
}

// distance calculates the Euclidean distance between two points
func distance(p1, p2 Point) float64 {
	// Using direct subtraction to avoid potential integer overflow
	di := p1.X - p2.X
	dj := p1.Y - p2.Y
	return float64(di*di + dj*dj)
}

// createTunnel creates a winding tunnel between two points
func (d *CADungeon) createTunnel(p1, p2 Point) {
	// Current position
	curr := p1

	// Create tunnel until we reach the destination
	for curr != p2 {
		// Mark current position as a cave
		d.Grid[curr.X][curr.Y] = 0

		// Determine the direction to move
		di := sign(p2.X - curr.X)
		dj := sign(p2.Y - curr.Y)

		// Randomly decide whether to move in i or j direction
		// This creates more winding, cave-like tunnels
		if rand.Intn(100) < 50 {
			// Try to move in i direction if possible
			if di != 0 {
				curr.X += di
			} else {
				curr.Y += dj
			}
		} else {
			// Try to move in j direction if possible
			if dj != 0 {
				curr.Y += dj
			} else {
				curr.X += di
			}
		}

		// Sometimes add a random variation to make the tunnel more cave-like
		if rand.Intn(100) < 20 {
			// Create a small cave-like bulge
			bulgeI := curr.X + rand.Intn(3) - 1
			bulgeJ := curr.Y + rand.Intn(3) - 1

			// Ensure the bulge is within bounds
			if bulgeI >= 0 && bulgeI < len(d.Grid) && bulgeJ >= 0 && bulgeJ < len(d.Grid[0]) {
				d.Grid[bulgeI][bulgeJ] = 0
			}
		}
	}
}

// sign returns the sign of an integer: -1, 0, or 1
func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

// Update CreateCADungeon to connect caves
func CreateCADungeon(width, height int) *CADungeon {
	d := NewCADungeon(width, height)
	d.Print()
	for i := 0; i < 10; i++ {
		d.Step()
	}
	d.Print()

	// Connect any disconnected caves
	println("\nConnecting disconnected caves...")
	d.ConnectCaves()
	d.Print()

	return d
}
