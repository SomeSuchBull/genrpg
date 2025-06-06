package dungeon

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

type Tile string

const (
	Wall            Tile = "Wall"
	Perimeter            = "Perimeter"
	Floor                = "Floor"
	SplitHorizontal      = "SplitHorizontal"
	SplitVertical        = "SplitVertical"
	Door                 = "Door"
)

type Cell struct {
	Tile          Tile
	DoorPerimeter DoorPerimeter
	PathValue     int
	ID            int
	CorridorID    int
	Connected     bool
}

type DoorPerimeter struct {
	DoorID int
	RoomID int
}

type Point struct {
	X, Y int
}

type BSPDoor struct {
	Direction Point
	Point
	ID int
}

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
	Doors                         []BSPDoor
	ConnectedRooms                []int
	Point                         // Top-left corner
	ID                            int
	Width, Height, ShiftX, ShiftY int
	ExpandedX, ExpandedY          int
}

type Split struct {
	Points       [2]Point
	Type         Tile
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
	Corridors                 map[int]map[int]bool
	Rooms                     map[int]*BSPRoom
	InitialGrid               [][]Tile
	ExpandedGrid              [][]Cell
	Splits                    []BSPNode
	Width, Height             int
	ActualWidth, ActualHeight int
	ShiftX, ShiftY            int
	Expanded                  bool
	UseSplitsAsCorridors      bool
}

// Initialize dungeon grid with walls
func NewDungeon(width, height int, expanded, useSplitsAsCorridors bool) *BSPDungeon {
	grid := make([][]Tile, height)
	for i := range grid {
		grid[i] = make([]Tile, width)
		for j := range grid[i] {
			grid[i][j] = Wall
		}
	}
	return &BSPDungeon{Width: width, Height: height, InitialGrid: grid, Expanded: expanded, UseSplitsAsCorridors: useSplitsAsCorridors,
		Corridors: map[int]map[int]bool{}, Rooms: map[int]*BSPRoom{}}
}

func (d *BSPDungeon) CreateActualGrid() {
	grid := make([][]Cell, d.ActualHeight)
	for i := range grid {
		grid[i] = make([]Cell, d.ActualWidth)
		for j := range grid[i] {
			cell := Cell{Tile: Wall, PathValue: normalDijkstraWeight}
			grid[i][j] = cell
		}
	}
	d.ExpandedGrid = grid
}

// Carve a room into the grid
func (d *BSPDungeon) CarveRoom(room *BSPRoom) {
	for y := room.Y; y < room.Y+room.Height; y++ {
		for x := room.X; x < room.X+room.Width; x++ {
			if y == room.Y && x == room.X {
				d.InitialGrid[y][x] = Tile(fmt.Sprint(room.ID))
				continue
			}
			d.InitialGrid[y][x] = Floor
		}
	}
	for y := room.ExpandedY - 1; y <= room.ExpandedY+room.Height; y++ {
		for x := room.ExpandedX - 1; x <= room.ExpandedX+room.Width; x++ {
			d.ExpandedGrid[y][x].Tile = Perimeter
			d.ExpandedGrid[y][x].PathValue = blockedDijkstraWeight
		}
	}
	for y := room.ExpandedY; y < room.ExpandedY+room.Height; y++ {
		for x := room.ExpandedX; x < room.ExpandedX+room.Width; x++ {
			if y == room.ExpandedY && x == room.ExpandedX {
				d.ExpandedGrid[y][x].Tile = Tile(fmt.Sprint(room.ID))
				continue
			}
			d.ExpandedGrid[y][x].Tile = Floor
		}
	}
}

func (d *BSPDungeon) CarveSplitCorridor(node *BSPNode) {
	if node.Split.Type == SplitVertical {
		x := node.Split.Points[0].X + 2*node.ShiftX + 2
		for y := node.Split.Points[0].Y + 2*node.ShiftY; y < d.ActualHeight; y++ {
			if d.ExpandedGrid[y][x].Tile != Wall && d.ExpandedGrid[y][x].Tile != Perimeter {
				break
			}
			if (d.ExpandedGrid[y][x].Tile == Wall || d.ExpandedGrid[y][x].Tile == Perimeter) && y != 0 && y != d.ActualHeight-1 {
				d.ExpandedGrid[y][x].Tile = SplitVertical
			}
		}
	} else if node.Split.Type == SplitHorizontal {
		y := node.Split.Points[0].Y + 2*node.ShiftY + 2
		for x := node.Split.Points[0].X + 2*node.ShiftX; x < d.ActualWidth; x++ {
			if d.ExpandedGrid[y][x].Tile != Wall && d.ExpandedGrid[y][x].Tile != Perimeter {
				break
			}
			if (d.ExpandedGrid[y][x].Tile == Wall || d.ExpandedGrid[y][x].Tile == Perimeter) && x != 0 && x != d.ActualWidth-1 {
				d.ExpandedGrid[y][x].Tile = SplitHorizontal
			}
		}
	}
}

type BSPNode struct {
	Split                Split
	Left, Right          *BSPNode
	X, Y, Width, Height  int
	ID                   int
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
			if d.InitialGrid[y+split-1][i] == Wall || d.InitialGrid[y+split-1][i] == Perimeter {
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
			if d.InitialGrid[i][x+split-1] == Wall || d.InitialGrid[i][x+split-1] == Perimeter {
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
func (d *BSPDungeon) CarveBSP(node *BSPNode, count *int, numberOfRooms int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil && *count <= numberOfRooms && rand.Intn(5) != 0 {
		roomWidth := rand.Intn(node.Width/2) + node.Width/2
		if roomWidth < 2 {
			roomWidth = 2
		}
		roomHeight := rand.Intn(node.Height/2) + node.Height/2
		if roomHeight < 2 {
			roomHeight = 2
		}
		if roomHeight > roomWidth && float64(roomHeight)/float64(roomWidth) > 3 {
			roomHeight = 3 * roomWidth
		}
		if roomWidth > roomHeight && float64(roomWidth)/float64(roomHeight) > 3 {
			roomWidth = 3 * roomHeight
		}
		x := node.X + rand.Intn(node.Width-roomWidth)
		y := node.Y + rand.Intn(node.Height-roomHeight)
		expandedY := 1 + y
		expandedX := 1 + x
		if d.Expanded {
			expandedY += 2 * node.ShiftY
			expandedX += 2 * node.ShiftX
		}
		*count++
		room := BSPRoom{
			ID:    *count,
			Point: Point{X: x, Y: y}, Width: roomWidth, Height: roomHeight,
			ShiftX: node.ShiftX, ShiftY: node.ShiftY,
			ExpandedY: expandedY, ExpandedX: expandedX,
			Doors: []BSPDoor{},
		}
		d.Rooms[*count] = &room
		d.CarveRoom(&room)
		return
	}
	if d.UseSplitsAsCorridors {
		d.CarveSplitCorridor(node)
	}
	d.CarveBSP(node.Left, count, numberOfRooms)
	d.CarveBSP(node.Right, count, numberOfRooms)
}

func (d *BSPDungeon) PlaceBSPDoors() {
	for _, room := range d.Rooms {
		usedDirections := []Point{}
		doorID := 1
		for {
			doorPoint, direction := d.randomDoorDirection(*room)
			breaking := false
			for _, usedDirection := range usedDirections {
				if direction == usedDirection {
					breaking = true
					break
				}
			}
			if breaking {
				break
			}
			doorCell := &d.ExpandedGrid[doorPoint.Y][doorPoint.X]
			doorCell.Tile = Door
			doorCell.ID = room.ID
			doorCell.PathValue = room.ID + 1
			extraY, extraX := doorPoint.Y+direction.Y, doorPoint.X+direction.X
			extraCell := &d.ExpandedGrid[extraY][extraX]
			if extraCell.PathValue == 1 {
				extraCell.PathValue = room.ID + 1
				extraCell.Tile = Floor
				extraCell.ID = room.ID
				for _, dir := range deadEnds[oppositeDirections[direction]] {
					d.ExpandedGrid[extraY+dir.Y][extraX+dir.X].DoorPerimeter = DoorPerimeter{DoorID: doorID, RoomID: room.ID}
				}
			}
			room.Doors = append(room.Doors, BSPDoor{Point: doorPoint, Direction: direction, ID: doorID})
			if d.UseSplitsAsCorridors {
				d.ExpandToSplitCorridor(doorPoint, direction)
				break
			} else {
				usedDirections = append(usedDirections, direction)
				if len(usedDirections) == 4 {
					break
				}
			}
			doorID++
		}
		room.Doors = sortDoors(room.Doors)
	}
}

func sortDoors(doors []BSPDoor) []BSPDoor {
	for i := 0; i < len(doors); i++ {
		for j := i + 1; j < len(doors); j++ {
			if dirMap[doors[j].Direction] < dirMap[doors[i].Direction] {
				doors[i], doors[j] = doors[j], doors[i]
			}
		}
	}
	return doors
}

func (d *BSPDungeon) ExpandToSplitCorridor(previousPoint Point, direction Point) {
	currentPoint := Point{X: previousPoint.X + direction.X, Y: previousPoint.Y + direction.Y}
	if d.ExpandedGrid[currentPoint.Y][currentPoint.X].Tile == Wall || d.ExpandedGrid[currentPoint.Y][currentPoint.X].Tile == Perimeter {
		switch direction {
		case N:
			d.ExpandedGrid[currentPoint.Y][currentPoint.X].Tile = SplitVertical
		case S:
			d.ExpandedGrid[currentPoint.Y][currentPoint.X].Tile = SplitVertical
		case E:
			d.ExpandedGrid[currentPoint.Y][currentPoint.X].Tile = SplitHorizontal
		case W:
			d.ExpandedGrid[currentPoint.Y][currentPoint.X].Tile = SplitHorizontal
		}
		d.ExpandToSplitCorridor(currentPoint, direction)
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
			if point.X == d.ActualWidth || point.Y == d.ActualHeight ||
				point.X < 0 || point.Y < 0 {
				break
			}
			if d.ExpandedGrid[point.Y][point.X].Tile != Wall && d.ExpandedGrid[point.Y][point.X].Tile != Perimeter {
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

// This is fixed, but I'm still not satisfied
// I think this is more of a bsp problem, the cells are too far removed
func (d *BSPDungeon) CarveDijkstraCorridors() {
	for _, room := range d.Rooms {
		for _, door := range room.Doors {
			if d.ExpandedGrid[door.Y][door.X].CorridorID != 0 {
				continue
			}
			start := Point{X: door.X + door.Direction.X, Y: door.Y + door.Direction.Y}
			if d.ExpandedGrid[start.Y][start.X].Connected {
				continue
			}
			_, endPoint, path := dijkstraFindNearest(*d, start, room.ID, door.ID)
			if endPoint.X == -1 || endPoint.Y == -1 {
				continue
			}
			corridorID := d.ExpandedGrid[endPoint.Y][endPoint.X].CorridorID
			if corridorID == 0 {
				corridorID = len(d.Corridors) + 1
				d.Corridors[corridorID] = map[int]bool{}
			}
			d.Corridors[corridorID][room.ID] = true
			d.Corridors[corridorID][d.ExpandedGrid[endPoint.Y][endPoint.X].ID] = true
			for k, _ := range d.Corridors[corridorID] {
				room.ConnectedRooms = append(room.ConnectedRooms, k)
			}
			d.ExpandedGrid[door.Y][door.X].CorridorID = corridorID
			d.ExpandedGrid[door.Y][door.X].Connected = true
			d.ExpandedGrid[start.Y][start.X].Connected = true
			d.ExpandedGrid[endPoint.Y][endPoint.X].Connected = true
			for _, point := range path {
				d.ExpandedGrid[point.Y][point.X].CorridorID = corridorID
				// if d.ExpandedGrid[point.Y][point.X].Tile == Wall || d.ExpandedGrid[point.Y][point.X].Tile == Perimeter || d.ExpandedGrid[point.Y][point.X].Tile == Floor {
				if d.ExpandedGrid[point.Y][point.X].Tile != Door {
					d.ExpandedGrid[point.Y][point.X].Tile = Floor
					d.ExpandedGrid[point.Y][point.X].PathValue = room.ID + 1
					// grid[point.Y][point.X] = d.ExpandedGrid[point.Y][point.X].PathValue
				}
			}
		}
	}
}

func NewBSPDungeon(width, height, minSize, numberOfRooms int) *BSPDungeon {
	d := NewDungeon(width, height, true, false)
	// rand.Seed(422134)
	seed := time.Now().UnixNano()
	// seed = 1738669161687714725
	rand.Seed(seed)
	counter := new(int)
	root := SplitSpace(0, 0, width, height, minSize, d, counter)
	d.ShiftX = CalculateShiftX(root, 0)
	d.ShiftY = CalculateShiftY(root, 0)
	d.ActualWidth = d.Width + 2
	d.ActualHeight = d.Height + 2
	if d.Expanded {
		d.ActualWidth += 2 * d.ShiftX
		d.ActualHeight += 2 * d.ShiftY
	}
	d.CreateActualGrid()
	count := new(int)
	d.CarveBSP(root, count, numberOfRooms)
	d.PlaceBSPDoors()
	// d.Print()
	if !d.UseSplitsAsCorridors {
		d.CarveDijkstraCorridors()
	}
	d.RemoveBSPDeadEnds()
	d.Print()
	fmt.Println("Seed:", seed)
	grid := d.getPathGrid()
	islands := findIslands(grid)
	fmt.Println(len(islands))
	for len(islands) > 1 {
		grid = connectIslands(grid, islands)
		islands = findIslands(grid)
	}
	d.CorrectTheGrid(grid)
	d.Print()
	// printGrid(grid)
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

func (d *BSPDungeon) CorrectTheGrid(grid [][]int) {
	for y := 0; y < d.ActualHeight; y++ {
		for x := 0; x < d.ActualWidth; x++ {
			if grid[y][x] == 2 && d.ExpandedGrid[y][x].Tile == Wall {
				d.ExpandedGrid[y][x].Tile = Floor
			}
		}
	}
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			switch cell {
			case 1:
				fmt.Print(" .")
			default:
				fmt.Printf("%2d", cell)
			}
		}
		fmt.Println()
	}
}

func (d *BSPDungeon) getPathGrid() [][]int {
	grid := make([][]int, d.ActualHeight)
	for i := range grid {
		grid[i] = make([]int, d.ActualWidth)
		for j := range grid[i] {
			if d.ExpandedGrid[i][j].PathValue > 1 {
				grid[i][j] = 2
			} else {
				grid[i][j] = d.ExpandedGrid[i][j].PathValue
			}
		}
	}
	return grid
}

func (d *BSPDungeon) Print() {
	// for _, row := range d.InitialGrid {
	// 	for _, cell := range row {
	// 		if cell == Wall {
	// 			print(". ")
	// 		} else if cell == Floor {
	// 			print("█ ")
	// 		} else if cell == SplitHorizontal {
	// 			print("— ")
	// 		} else if cell == SplitVertical {
	// 			print("| ")
	// 		} else if cell == Door {
	// 			print("D ")
	// 		} else {
	// 			print(fmt.Sprintf("%-2s", cell))
	// 		}
	// 	}
	// 	println()
	// }
	// println()
	// for _, row := range d.ExpandedGrid {
	// 	for _, cell := range row {
	// 		if cell.PathValue == blockedDijkstraWeight {
	// 			print("█ ")
	// 		} else {
	// 			print(fmt.Sprintf("%-2d", cell.PathValue))
	// 		}
	// 	}
	// 	println()
	// }
	// println()
	for _, row := range d.ExpandedGrid {
		for _, cell := range row {
			// if cell == Wall || cell == Perimeter {
			if cell.Tile == Wall || cell.Tile == Perimeter {
				print(". ")
			} else if cell.Tile == Floor {
				print("█ ")
			} else if cell.Tile == SplitHorizontal {
				print("— ")
			} else if cell.Tile == SplitVertical {
				print("| ")
			} else if cell.Tile == Door {
				print("D ")
			} else {
				print(fmt.Sprintf("%-2s", cell.Tile))
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

func (d *BSPDungeon) RemoveBSPDeadEnds() {
	for y := 1; y < d.ActualHeight-1; y++ {
		for x := 1; x < d.ActualWidth-1; x++ {
			if d.ExpandedGrid[y][x].Tile == SplitHorizontal || d.ExpandedGrid[y][x].Tile == SplitVertical {
				d.collapseDeadEnd(Point{X: x, Y: y})
			}
		}
	}
}

func (d *BSPDungeon) collapseDeadEnd(point Point) {
	if d.ExpandedGrid[point.Y][point.X].Tile == Wall || d.ExpandedGrid[point.Y][point.X].Tile == Perimeter {
		return
	}
	for direction, deadEnd := range deadEnds {
		if d.checkDeadEnd(point, deadEnd) {
			d.ExpandedGrid[point.Y][point.X].Tile = Wall
			d.collapseDeadEnd(Point{X: point.X + direction.X, Y: point.Y + direction.Y})
		}
	}
}

func (d *BSPDungeon) checkDeadEnd(point Point, deadEnd []Point) bool {
	for _, direction := range deadEnd {
		if d.ExpandedGrid[point.Y+direction.Y][point.X+direction.X].Tile != Wall && d.ExpandedGrid[point.Y+direction.Y][point.X+direction.X].Tile != Perimeter {
			return false
		}
	}
	return true
}

// Finds all islands of `2`s and their bordering `1`s
func findIslands(grid [][]int) [][][]int {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var islands [][][]int // List of islands, each containing (x,y) positions

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !visited[r][c] && (grid[r][c] == 2 || grid[r][c] == -1) {
				// New island found
				queue := list.New()
				queue.PushBack([]int{r, c})
				island := [][]int{}

				for queue.Len() > 0 {
					cell := queue.Remove(queue.Front()).([]int)
					x, y := cell[0], cell[1]

					if visited[x][y] {
						continue
					}

					visited[x][y] = true
					island = append(island, []int{x, y})

					// Explore neighbors
					for _, d := range directions {
						nx, ny := x+d.Y, y+d.X
						if nx >= 0 && ny >= 0 && nx < rows && ny < cols &&
							!visited[nx][ny] && (grid[nx][ny] == 2 || grid[nx][ny] == -1) {
							queue.PushBack([]int{nx, ny})
						}
					}
				}
				islands = append(islands, island)
			}
		}
	}
	return islands
}

// Finds the shortest path between disconnected `2` islands using BFS
func connectIslands(grid [][]int, islands [][][]int) [][]int {
	if len(islands) <= 1 {
		return grid // Already connected
	}

	rows, cols := len(grid), len(grid[0])
	queue := list.New()
	visited := make([][]bool, rows)
	prev := make([][]*[]int, rows) // Store previous position for path reconstruction

	for i := range visited {
		visited[i] = make([]bool, cols)
		prev[i] = make([]*[]int, cols)
	}

	// Start BFS from first island's `2`s
	for _, cell := range islands[0] {
		x, y := cell[0], cell[1]
		if grid[x][y] == 2 { // Only start from `2`s
			queue.PushBack([]int{x, y})
			visited[x][y] = true
		}
	}

	// Perform BFS to find the shortest path to another island
	var targetCell []int
	found := false
	for queue.Len() > 0 && !found {
		cell := queue.Remove(queue.Front()).([]int)
		x, y := cell[0], cell[1]

		for _, d := range directions {
			nx, ny := x+d.Y, y+d.X

			// Stay in bounds and avoid `-1`
			if nx >= 0 && ny >= 0 && nx < rows && ny < cols && !visited[nx][ny] && grid[nx][ny] != -1 {
				visited[nx][ny] = true
				prev[nx][ny] = &[]int{x, y} // Store previous step
				queue.PushBack([]int{nx, ny})

				// If we reached another `2` from a different island, stop BFS
				for _, island := range islands[1:] {
					for _, icell := range island {
						if icell[0] == nx && icell[1] == ny && grid[nx][ny] == 2 {
							targetCell = []int{nx, ny}
							found = true
							break
						}
					}
					if found {
						break
					}
				}
			}
		}
	}

	// Reconstruct the shortest path and convert necessary `1`s to `2`s
	for targetCell != nil {
		x, y := targetCell[0], targetCell[1]
		if grid[x][y] == 1 {
			grid[x][y] = 2 // Change `1` to `2` to create the connection
		}

		// Fix: Dereference prev[x][y] safely
		if prev[x][y] != nil {
			targetCell = *prev[x][y]
		} else {
			targetCell = nil // End path reconstruction
		}
	}

	return grid
}
