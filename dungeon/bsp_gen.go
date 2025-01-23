package dungeon

import (
	"fmt"
	"math/rand"
)

type Tile string

const (
	Wall            Tile = "Wall"
	Floor                = "Floor"
	SplitHorizontal      = "SplitHorizontal"
	SplitVertical        = "SplitVertical"
	Door                 = "Door"
)

type Point struct {
	X, Y int
}

var N = Point{X: 0, Y: -1}
var NE = Point{X: 1, Y: -1}
var E = Point{X: 1, Y: 0}
var SE = Point{X: 1, Y: 1}
var S = Point{X: 0, Y: 1}
var SW = Point{X: -1, Y: 1}
var W = Point{X: -1, Y: 0}
var NW = Point{X: -1, Y: -1}

var deadEnds = map[Point][]Point{
	S: {W, NW, N, NE, E},
	E: {S, SW, W, NW, N},
	N: {E, SE, S, SW, W},
	W: {N, NE, E, SE, S},
}

func (p Point) String() string {
	return fmt.Sprintf("(%2d,%2d)", p.X, p.Y)
}

type BSPRoom struct {
	Point                         // Top-left corner
	Width, Height, ShiftX, ShiftY int
	ExpandedX, ExpandedY          int
}

type Split struct {
	Type         Tile
	Points       [2]Point
	SplitsBefore int
	SplitsAfter  int
}

func (s Split) String() string {
	splitType := "Horizontal"
	if s.Type == SplitVertical {
		splitType = "Vertical"
	}
	return fmt.Sprintf("%10s split %s-%s", splitType, s.Points[0], s.Points[1])
}

type BSPDungeon struct {
	Width, Height             int
	ActualWidth, ActualHeight int
	ShiftX, ShiftY            int
	InitialGrid               [][]Tile
	ExpandedGrid              [][]Tile
	Rooms                     []BSPRoom
	Splits                    []BSPNode
	Expanded                  bool
}

// Initialize dungeon grid with walls
func NewDungeon(width, height int) *BSPDungeon {
	grid := make([][]Tile, height)
	for i := range grid {
		grid[i] = make([]Tile, width)
		for j := range grid[i] {
			grid[i][j] = Wall
		}
	}
	return &BSPDungeon{Width: width, Height: height, InitialGrid: grid}
}

func (d *BSPDungeon) CreateActualGrid() {
	grid := make([][]Tile, d.ActualHeight)
	for i := range grid {
		grid[i] = make([]Tile, d.ActualWidth)
		for j := range grid[i] {
			grid[i][j] = Wall
		}
	}
	d.ExpandedGrid = grid
}

// Carve a room into the grid
func (d *BSPDungeon) CarveRoom(room *BSPRoom) {
	for y := room.Y; y < room.Y+room.Height; y++ {
		for x := room.X; x < room.X+room.Width; x++ {
			d.InitialGrid[y][x] = Floor
		}
	}
	for y := room.ExpandedY; y < room.ExpandedY+room.Height; y++ {
		for x := room.ExpandedX; x < room.ExpandedX+room.Width; x++ {
			d.ExpandedGrid[y][x] = Floor
		}
	}
}

func (d *BSPDungeon) CarveCorridor(node *BSPNode) {
	if node.Split.Type == SplitVertical {
		x := node.Split.Points[0].X + 2*node.ShiftX + 2
		for y := node.Split.Points[0].Y + 2*node.ShiftY; y < d.ActualHeight; y++ {
			if d.ExpandedGrid[y][x] != Wall {
				break
			}
			if d.ExpandedGrid[y][x] == Wall && y != 0 && y != d.ActualHeight-1 {
				d.ExpandedGrid[y][x] = SplitVertical
			}
		}
	} else if node.Split.Type == SplitHorizontal {
		y := node.Split.Points[0].Y + 2*node.ShiftY + 2
		for x := node.Split.Points[0].X + 2*node.ShiftX; x < d.ActualWidth; x++ {
			if d.ExpandedGrid[y][x] != Wall {
				break
			}
			if d.ExpandedGrid[y][x] == Wall && x != 0 && x != d.ActualWidth-1 {
				d.ExpandedGrid[y][x] = SplitHorizontal
			}
		}
	}
}

type BSPNode struct {
	X, Y, Width, Height  int
	Left, Right          *BSPNode
	ID                   int
	Split                Split
	ShiftX, ShiftY       int
	EndShiftX, EndShiftY int
}

func (n BSPNode) String() string {
	return fmt.Sprintf("%s | ID: %-2d", n.Split, n.ID)
}

// Recursively split the grid
func SplitSpace(x, y, width, height, minSize int, d *BSPDungeon, counter *int) *BSPNode {
	if width <= minSize*2 || height <= minSize*2 {
		*counter++
		return &BSPNode{X: x, Y: y, Width: width, Height: height, ID: *counter}
	}

	horizontalSplit := rand.Intn(2) == 0
	if width > height {
		horizontalSplit = false
	} else if height > width {
		horizontalSplit = true
	}

	if horizontalSplit {
		split := rand.Intn(height-minSize*2) + minSize
		*counter++
		d.InitialGrid[y+split-1][x] = Tile(fmt.Sprintf("%d", *counter))
		for i := x + 1; i < x+width; i++ {
			if d.InitialGrid[y+split-1][i] == Wall {
				d.InitialGrid[y+split-1][i] = SplitHorizontal
			}
		}
		node := &BSPNode{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
			ID:     *counter,
			Split:  Split{Type: SplitHorizontal, Points: [2]Point{{X: x, Y: y + split - 1}, {X: x + width - 1, Y: y + split - 1}}},
		}
		node.Left = SplitSpace(x, y, width, split, minSize, d, counter)
		node.Right = SplitSpace(x, y+split, width, height-split, minSize, d, counter)
		d.Splits = append(d.Splits, *node)
		return node
	} else {
		split := rand.Intn(width-minSize*2) + minSize
		*counter++
		d.InitialGrid[y][x+split-1] = Tile(fmt.Sprintf("%d", *counter))
		for i := y + 1; i < y+height; i++ {
			if d.InitialGrid[i][x+split-1] == Wall {
				d.InitialGrid[i][x+split-1] = SplitVertical
			}
		}
		node := &BSPNode{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
			ID:     *counter,
			Split:  Split{Type: SplitVertical, Points: [2]Point{{X: x + split - 1, Y: y}, {X: x + split - 1, Y: y + height - 1}}},
		}
		node.Left = SplitSpace(x, y, split, height, minSize, d, counter)
		node.Right = SplitSpace(x+split, y, width-split, height, minSize, d, counter)
		d.Splits = append(d.Splits, *node)
		return node

	}
}

// Place rooms within BSP regions
func CarveBSP(node *BSPNode, dungeon *BSPDungeon, count *int, numberOfRooms int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil && *count <= numberOfRooms && rand.Intn(3) != 0 {
		roomWidth := rand.Intn(node.Width/2) + node.Width/2
		if roomWidth < 2 {
			roomWidth = 2
		}
		roomHeight := rand.Intn(node.Height/2) + node.Height/2
		if roomHeight < 2 {
			roomHeight = 2
		}
		x := node.X + rand.Intn(node.Width-roomWidth)
		y := node.Y + rand.Intn(node.Height-roomHeight)

		room := BSPRoom{
			Point: Point{X: x, Y: y}, Width: roomWidth, Height: roomHeight,
			ShiftX: node.ShiftX, ShiftY: node.ShiftY,
			ExpandedY: 1 + y + 2*node.ShiftY, ExpandedX: 1 + x + 2*node.ShiftX,
		}
		dungeon.Rooms = append(dungeon.Rooms, room)
		dungeon.CarveRoom(&room)
		*count++
		return
	}
	dungeon.CarveCorridor(node)
	CarveBSP(node.Left, dungeon, count, numberOfRooms)
	CarveBSP(node.Right, dungeon, count, numberOfRooms)
}

func (d *BSPDungeon) PlaceDoors() {
	for _, room := range d.Rooms {
		door, direction := d.randomDoorDirection(room)
		// door, _ := d.randomDoorDirection(room)
		d.ExpandedGrid[door.Y][door.X] = Door
		d.ExpandCorridor(door, direction)
	}
}

func (d *BSPDungeon) ExpandCorridor(previousPoint Point, direction Point) {
	currentPoint := Point{X: previousPoint.X + direction.X, Y: previousPoint.Y + direction.Y}
	if d.ExpandedGrid[currentPoint.Y][currentPoint.X] == Wall {
		switch direction {
		case N:
			d.ExpandedGrid[currentPoint.Y][currentPoint.X] = SplitVertical
		case S:
			d.ExpandedGrid[currentPoint.Y][currentPoint.X] = SplitVertical
		case E:
			d.ExpandedGrid[currentPoint.Y][currentPoint.X] = SplitHorizontal
		case W:
			d.ExpandedGrid[currentPoint.Y][currentPoint.X] = SplitHorizontal
		}
		d.ExpandCorridor(currentPoint, direction)
	}
}

func (d *BSPDungeon) randomDoorDirection(room BSPRoom) (Point, Point) {
	possibleDoors := [][2]Point{}
	if room.ShiftX > 0 {
		possibleDoors = append(possibleDoors, [2]Point{Point{X: room.ExpandedX - 1, Y: room.ExpandedY + rand.Intn(room.Height)}, W})
	}
	if room.ShiftY > 0 {
		possibleDoors = append(possibleDoors, [2]Point{Point{X: room.ExpandedX + rand.Intn(room.Width), Y: room.ExpandedY - 1}, N})
	}
	if room.ShiftX < d.ShiftX {
		possibleDoors = append(possibleDoors, [2]Point{Point{X: room.ExpandedX + room.Width, Y: room.ExpandedY + rand.Intn(room.Height)}, E})
	}
	if room.ShiftY < d.ShiftY {
		possibleDoors = append(possibleDoors, [2]Point{Point{X: room.ExpandedX + rand.Intn(room.Width), Y: room.ExpandedY + room.Height}, S})
	}
	roll := rand.Intn(len(possibleDoors))
	randomDoor := possibleDoors[roll]
	for {
		breaking := false
		point := Point{X: randomDoor[0].X, Y: randomDoor[0].Y}
		for {
			point = Point{X: point.X + randomDoor[1].X, Y: point.Y + randomDoor[1].Y}
			if point.X == d.ActualWidth || point.Y == d.ActualHeight {
				break
			}
			if d.ExpandedGrid[point.Y][point.X] != Wall {
				breaking = true
				break
			}
		}
		if breaking {
			break
		}
		possibleDoors = append(possibleDoors[:roll], possibleDoors[roll+1:]...)
		if len(possibleDoors) == 0 {
			break
		}
		roll = rand.Intn(len(possibleDoors))
		randomDoor = possibleDoors[roll]
	}
	return randomDoor[0], randomDoor[1]
}

func NewBSPDungeon(width, height, minSize, numberOfRooms int) *BSPDungeon {
	d := NewDungeon(width, height)
	counter := new(int)
	root := SplitSpace(0, 0, width, height, minSize, d, counter)
	d.ShiftX = CalculateShiftX(root, 0)
	d.ShiftY = CalculateShiftY(root, 0)
	d.ActualWidth = d.Width + 2
	d.ActualHeight = d.Height + 2
	// if d.Expanded {
	d.ActualWidth += 2 * d.ShiftX
	d.ActualHeight += 2 * d.ShiftY
	// }
	d.CreateActualGrid()
	count := new(int)
	CarveBSP(root, d, count, numberOfRooms)
	d.PlaceDoors()
	d.RemoveDeadEnds()
	d.Print()
	// printTreeSpatially(root, 0, 2)
	// GetShiftY(root)
	// fmt.Println("Width:", d.Width)
	// fmt.Println("ActualWidth:", d.ActualWidth)
	// fmt.Println("ShiftX:", d.ShiftX)
	// fmt.Println("Height:", d.Height)
	// fmt.Println("ActualHeight:", d.ActualHeight)
	// fmt.Println("ShiftY:", d.ShiftY)
	// } else {
	// 	println(d.MaxLeftHorizontalSplits(root, 0))
	// }
	return d
}

func (d *BSPDungeon) Print() {
	for _, row := range d.InitialGrid {
		for _, cell := range row {
			if cell == Wall {
				print(" .")
			} else if cell == Floor {
				print(" █")
			} else if cell == SplitHorizontal {
				print("——")
			} else if cell == SplitVertical {
				print(" |")
			} else if cell == Door {
				print(" D")
			} else {
				print(fmt.Sprintf("%2s", cell))
			}
		}
		println()
	}
	println()
	for _, row := range d.ExpandedGrid {
		for _, cell := range row {
			if cell == Wall {
				print(" .")
			} else if cell == Floor {
				print(" █")
			} else if cell == SplitHorizontal {
				print("——")
			} else if cell == SplitVertical {
				print(" |")
			} else if cell == Door {
				print(" D")
			} else {
				print(fmt.Sprintf("%2s", cell))
			}
		}
		println()
	}
	// for i, split := range d.Splits {
	// 	for _, split2 := range d.Splits[i+1:] {
	// 		if split2.Type == split.Type {
	// 			switch split.Type {
	// 			case SplitVertical:
	// 				if split.Points[0].X > split2.Points[0].X {
	// 					d.Splits[i].SplitsAfter++
	// 				}
	// 			default:
	// 				if split.Points[0].Y > split2.Points[0].Y {
	// 					d.Splits[i].SplitsAfter++
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// for _, split := range d.Splits {
	// 	fmt.Printf("%s\n", split)
	// }
}

// This is the way
// It can now be used for shifting rooms
// Probably best to init these CalculateShift funcs before carving rooms, in order to carve them into the "actual grid"
// at the time as well
func GetShiftX(node *BSPNode) {
	startingShift := 0
	node.ShiftX = CalculateShiftX(node.Left, startingShift)
}

func CalculateShiftX(node *BSPNode, previousShift int) int {
	if node.Split.Type == SplitVertical {
		node.ShiftX = CalculateShiftX(node.Left, previousShift)
		rightNodeShift := CalculateShiftX(node.Right, node.ShiftX+1)
		return rightNodeShift
	} else if node.Split.Type == SplitHorizontal {
		node.ShiftX = previousShift
		node.EndShiftX = max(CalculateShiftX(node.Left, node.ShiftX), CalculateShiftX(node.Right, node.ShiftX))
		return node.EndShiftX
	} else {
		node.ShiftX = previousShift
		return node.ShiftX
	}
}

func GetShiftY(node *BSPNode) {
	startingShift := 0
	node.ShiftY = CalculateShiftY(node.Left, startingShift)
}

func CalculateShiftY(node *BSPNode, previousShift int) int {
	if node.Split.Type == SplitHorizontal {
		node.ShiftY = CalculateShiftY(node.Left, previousShift)
		return CalculateShiftY(node.Right, node.ShiftY+1)
	} else if node.Split.Type == SplitVertical {
		node.ShiftY = previousShift
		node.EndShiftY = max(CalculateShiftY(node.Left, node.ShiftY), CalculateShiftY(node.Right, node.ShiftY))
		return node.EndShiftY
	} else {
		node.ShiftY = previousShift
		return node.ShiftY
	}
}

func (d *BSPDungeon) RemoveDeadEnds() {
	for y := 1; y < d.ActualHeight-1; y++ {
		for x := 1; x < d.ActualWidth-1; x++ {
			if d.ExpandedGrid[y][x] == SplitHorizontal || d.ExpandedGrid[y][x] == SplitVertical {
				d.collapseDeadEnd(Point{X: x, Y: y})
			}
		}
	}
}

func (d *BSPDungeon) collapseDeadEnd(point Point) {
	if d.ExpandedGrid[point.Y][point.X] == Wall {
		return
	}
	for direction, deadEnd := range deadEnds {
		if d.checkDeadEnd(point, deadEnd) {
			d.ExpandedGrid[point.Y][point.X] = Wall
			d.collapseDeadEnd(Point{X: point.X + direction.X, Y: point.Y + direction.Y})
		}
	}
}

func (d *BSPDungeon) checkDeadEnd(point Point, deadEnd []Point) bool {
	for _, direction := range deadEnd {
		if d.ExpandedGrid[point.Y+direction.Y][point.X+direction.X] != Wall {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
