package dungeon

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/genrpg/utils"
)

var dungeonLayouts = map[string][][]int{
	"Box": {
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	},
	"Cross": {
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	},
}

var corridorLayouts = map[string]int{
	"Labyrinth": 0,
	"Bent":      50,
	"Straight":  100,
}

var mapStyles = map[string]struct {
	Fill     color.RGBA
	Open     color.RGBA
	OpenGrid color.RGBA
}{
	"Standard": {
		Fill:     color.RGBA{0, 0, 0, 255},
		Open:     color.RGBA{255, 255, 255, 255},
		OpenGrid: color.RGBA{204, 204, 204, 255},
	},
}

type Cell struct {
	Blocked, Room, Corridor, Perimeter, Entrance                   bool
	Arch, Door, Locked, Trapped, Secret, Portc, StairDown, StairUp bool
	ID                                                             int
	Label                                                          string
}

type RoomData struct {
	ID, Row, Col, N, S, W, E, Height, Width, Area int
	Doors                                         map[string][]map[string]any
}

func NewCell() *Cell {
	return &Cell{}
}

type Proto struct {
	I, J, Height, Width *int
}

func (c *Cell) OpenSpace() bool {
	return c.Room || c.Corridor
}

func (c *Cell) DoorSpace() bool {
	return c.Arch || c.Door || c.Locked || c.Trapped || c.Secret || c.Portc
}

func (c *Cell) Espace() bool {
	return c.Entrance || c.DoorSpace()
}

func (c *Cell) NotEspace() {
	c.Entrance = false
	c.Arch = false
	c.Door = false
	c.Locked = false
	c.Trapped = false
	c.Secret = false
	c.Portc = false
}

func (c *Cell) Stairs() bool {
	return c.StairDown || c.StairUp
}

func (c *Cell) BlockRoom() bool {
	return c.Room || c.Blocked
}

func (c *Cell) BlockCorridor() bool {
	return c.Corridor || c.Perimeter || c.Blocked
}

func (c *Cell) BlockDoor() bool {
	return c.DoorSpace() || c.Blocked
}

func (c *Cell) IsOnlyCorridor() bool {
	return c.Corridor && !(c.Perimeter || c.Entrance ||
		c.Arch || c.Door || c.Locked || c.Trapped || c.Secret || c.Portc ||
		c.StairDown || c.StairUp)
}

var directionI = map[string]int{"N": -1, "S": 1, "E": 0, "W": 0}
var directionJ = map[string]int{"N": 0, "S": 0, "E": 1, "W": -1}

var opposite = map[string]string{"N": "S", "S": "N", "E": "W", "W": "E"}

var stairEnd = map[string]map[string]any{
	"N": {
		"walled":   [][]int{{1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}},
		"corridor": [][]int{{0, 0}, {1, 0}, {2, 0}},
		"stair":    []int{0, 0},
		"next":     []int{1, 0},
	},
	"S": {
		"walled":   [][]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}},
		"corridor": [][]int{{0, 0}, {-1, 0}, {-2, 0}},
		"stair":    []int{0, 0},
		"next":     []int{-1, 0},
	},
	"W": {
		"walled":   [][]int{{-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}},
		"corridor": [][]int{{0, 0}, {0, 1}, {0, 2}},
		"stair":    []int{0, 0},
		"next":     []int{0, 1},
	},
	"E": {
		"walled":   [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}},
		"corridor": [][]int{{0, 0}, {0, -1}, {0, -2}},
		"stair":    []int{0, 0},
		"next":     []int{0, -1},
	},
}

var closeEnd = map[string]map[string]any{
	"N": {
		"walled":  [][]int{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}},
		"close":   [][]int{{0, 0}},
		"recurse": []int{-1, 0},
	},
	"S": {
		"walled":  [][]int{{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}},
		"close":   [][]int{{0, 0}},
		"recurse": []int{1, 0},
	},
	"W": {
		"walled":  [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}},
		"close":   [][]int{{0, 0}},
		"recurse": []int{0, -1},
	},
	"E": {
		"walled":  [][]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}},
		"close":   [][]int{{0, 0}},
		"recurse": []int{0, 1},
	},
}

var directions = []string{"E", "N", "S", "W"}

type SeedOptions struct {
	Seed                                                      int64
	NRows, NCols, RoomMin, RoomMax, AddStairs, RemoveDeadEnds int
	DungeonLayout, RoomLayout, CorridorLayout                 string
}

func getOptions() *SeedOptions {
	return &SeedOptions{
		Seed:          time.Now().UnixNano(),
		NRows:         41,
		NCols:         41,
		RoomMin:       3,
		RoomMax:       9,
		AddStairs:     2,
		DungeonLayout: "None",
		// RoomLayout:     "Packed",
		RoomLayout:     "Scattered",
		CorridorLayout: "Straight",
		RemoveDeadEnds: 100,
	}
}

type DumbDungeon struct {
	Seed                                                                          SeedOptions
	Ni, Nj, MaxRow, MaxCol, NRooms, RoomBase, RoomRadix, LastRoomID, NRows, NCols int
	Cells                                                                         [][]*Cell
	Rooms                                                                         map[int]*RoomData
	Stairs, Doors                                                                 []map[string]any
}

// TODO: Implement mask - dungeon shape
func initCells(nRows, nCols int) [][]*Cell {
	cells := make([][]*Cell, nRows)
	for i := 0; i < nRows; i++ {
		cells[i] = make([]*Cell, nCols)
		for j := 0; j < nCols; j++ {
			cells[i][j] = NewCell()
		}
	}
	return cells
}

func initCoords(nRows, nCols int) [][]int {
	cells := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		cells[i] = make([]int, nCols)
		for j := 0; j < nCols; j++ {
			cells[i][j] = 0
		}
	}
	return cells
}

func (d *DumbDungeon) Print() {
	output := ""
	for i := 0; i < d.Seed.NRows-1; i++ {
		content := ""
		// grid := ""
		for j := 0; j < d.Seed.NCols-1; j++ {
			if j == 0 {
				content += fmt.Sprintf("%2d", i)
				// grid += fmt.Sprint("  ")
				continue
			}
			if i == 0 {
				content += fmt.Sprintf("%-2d", j)
				continue
			}
			cell := d.Cells[i][j]
			if cell.Label != "" {
				content += fmt.Sprint(utils.U(utils.B(cell.Label)))
			} else if !(cell.Blocked || !cell.Room) {
				content += fmt.Sprint("█")
			} else {
				content += fmt.Sprint("_")
			}
			content += fmt.Sprint("|")
			// grid += fmt.Sprint("-+")
		}
		content += fmt.Sprintln()
		// grid += fmt.Sprintln()
		// output += content + grid
		output += content
	}
	fmt.Println(output)
}

func CreateDumbDungeon(seed *SeedOptions) *DumbDungeon {
	if seed == nil {
		seed = getOptions()
	}
	ni := seed.NRows / 2
	nj := seed.NCols / 2
	nRows := ni * 2
	nCols := nj * 2
	maxRow := nRows - 1
	maxCol := nCols - 1
	roomBase := (seed.RoomMin + 1) / 2
	roomRadix := (seed.RoomMax-seed.RoomMin)/2 + 1
	dungeon := &DumbDungeon{Seed: *seed, Ni: ni, Nj: nj, MaxRow: maxRow, MaxCol: maxCol, RoomBase: roomBase, RoomRadix: roomRadix, Rooms: map[int]*RoomData{}, Stairs: []map[string]any{}, Doors: []map[string]any{}, NRows: nRows, NCols: nCols}
	dungeon.Cells = initCells(seed.NRows, seed.NCols)
	dungeon.EmplaceRooms()
	dungeon.OpenRooms()
	dungeon.LabelRooms()
	dungeon.Corridors()
	if dungeon.Seed.AddStairs > 0 {
		dungeon.EmplaceStairs()
	}
	dungeon.CleanDungeon()
	return dungeon
}

func (d *DumbDungeon) EmplaceRooms() {
	if d.Seed.RoomLayout == "Packed" {
		d.PackRooms()
	} else {
		d.ScatterRooms()
	}
}

func (d *DumbDungeon) PackRooms() {
	for i := 0; i < d.Ni; i++ {
		r := i*2 + 1
		for j := 0; j < d.Nj; j++ {
			c := j*2 + 1
			if d.Cells[r][c].Room || ((i == 0 || j == 0) && rand.Intn(2) == 0) {
				continue
			}
			d.EmplaceRoom(&Proto{I: &i, J: &j})
		}
	}
}

func (d *DumbDungeon) ScatterRooms() {
	nRooms := allocRooms(d)
	for i := 0; i < nRooms; i++ {
		d.EmplaceRoom(nil)
	}
}

func allocRooms(d *DumbDungeon) int {
	nRooms := (d.NCols * d.NRows) / (d.Seed.RoomMax * d.Seed.RoomMax)
	return nRooms
}

func (d *DumbDungeon) EmplaceRoom(proto *Proto) {
	if proto == nil {
		proto = &Proto{}
	}
	proto.SetRoom(d)
	r1 := *proto.I*2 + 1
	c1 := *proto.J*2 + 1
	r2 := (*proto.I+*proto.Height)*2 - 1
	c2 := (*proto.J+*proto.Width)*2 - 1
	if r1 <= 0 || r2 > d.MaxRow || c1 <= 0 || c2 > d.MaxCol {
		return
	}
	roomID := d.NRooms + 1

	if d.SoundRoom(r1, c1, r2, c2) {
		d.NRooms = roomID
		for i := r1; i <= r2; i++ {
			for j := c1; j <= c2; j++ {
				d.Cells[i][j].Room = true
			}
		}
	} else {
		return
	}
	d.LastRoomID = roomID

	for r := r1; r <= r2; r++ {
		for c := c1; c <= c2; c++ {
			cell := d.Cells[r][c]
			if cell.Entrance {
				cell.NotEspace()
			} else if cell.Perimeter {
				cell.Perimeter = false
			}
			cell.Room = true
			cell.ID = roomID
		}
	}
	height := (r2 - r1 + 1) * 10
	width := (c2 - c1 + 1) * 10
	d.Rooms[roomID] = &RoomData{ID: roomID, Row: r1, Col: c1, N: r1, S: r2, W: c1, E: c2, Height: height, Width: width, Area: height * width, Doors: map[string][]map[string]any{}}

	for r := r1 - 1; r <= r2+1; r++ {
		cell := d.Cells[r][c1-1]
		if !(cell.Room || cell.Entrance) {
			cell.Perimeter = true
		}
		cell = d.Cells[r][c2+1]
		if !(cell.Room || cell.Entrance) {
			cell.Perimeter = true
		}
	}
	for c := c1 - 1; c <= c2+1; c++ {
		cell := d.Cells[r1-1][c]
		if !(cell.Room || cell.Entrance) {
			cell.Perimeter = true
		}
		cell = d.Cells[r2+1][c]
		if !(cell.Room || cell.Entrance) {
			cell.Perimeter = true
		}
	}
}

func (p *Proto) SetRoom(d *DumbDungeon) {
	base := d.RoomBase
	radix := d.RoomRadix
	if p.Height == nil {
		p.Height = new(int)
		r := radix
		if p.I != nil {
			a := d.Ni - base - *p.I
			if a > radix {
				a = radix
			}
			r = a
		}
		if r <= 0 {
			r = 1
		}
		*p.Height = base + rand.Intn(r)
	}
	if p.Width == nil {
		p.Width = new(int)
		r := radix
		if p.J != nil {
			a := d.Nj - base - *p.J
			if a > radix {
				a = radix
			}
			r = a
		}
		if r <= 0 {
			r = 1
		}
		*p.Width = base + rand.Intn(r)
	}
	if p.I == nil {
		p.I = new(int)
		*p.I = rand.Intn(d.Ni - *p.Height)
	}
	if p.J == nil {
		p.J = new(int)
		*p.J = rand.Intn(d.Nj - *p.Width)
	}
}

func (d *DumbDungeon) SoundRoom(r1, c1, r2, c2 int) bool {
	for i := r1; i <= r2; i++ {
		for j := c1; j <= c2; j++ {
			if d.Cells[i][j].Blocked {
				return false
			}
			if d.Cells[i][j].Room {
				return false
			}
		}
	}
	return true
}

func (d *DumbDungeon) OpenRooms() {
	connections := map[string]int{}
	for id := 1; id <= d.NRooms; id++ {
		d.OpenRoom(d.Rooms[id], connections)
	}
}

func (d *DumbDungeon) OpenRoom(room *RoomData, connections map[string]int) {
	doorSills := d.DoorSills(room)
	if len(doorSills) <= 0 {
		return
	}
	nOpens := d.allocOpens(room)

	for i := 0; i < nOpens; {
		if len(doorSills) == 0 {
			break
		}
		sills, _ := utils.Splice(&doorSills, rand.Intn(len(doorSills)), 1)
		if len(sills) == 0 {
			break
		}
		sill := sills[0]
		doorR := sill["doorR"].(int)
		doorC := sill["doorC"].(int)
		doorCell := d.Cells[doorR][doorC]
		if doorCell.DoorSpace() {
			continue
		}
		outID := sill["outID"].(int)
		if outID != 0 {
			ids := []int{room.ID, outID}
			sort.Ints(ids)
			connect := fmt.Sprintf("%d-%d", ids[0], ids[1])
			if connections[connect] != 0 {
				continue
			}
			connections[connect]++
		}

		openR := sill["sillR"].(int)
		openC := sill["sillC"].(int)
		openDir := sill["dir"].(string)

		for ii := 0; ii < 3; ii++ {
			r := openR + (directionI[openDir] * ii)
			c := openC + (directionJ[openDir] * ii)
			d.Cells[r][c].Perimeter = false
			d.Cells[r][c].Entrance = true
		}

		door := map[string]any{"row": doorR, "col": doorC}
		doorType := rand.Intn(110)
		switch {
		case doorType < 15:
			d.Cells[doorR][doorC].Arch = true
			door["key"] = "Arch"
			door["type"] = "Archway"
		case doorType < 60:
			d.Cells[doorR][doorC].Door = true
			door["key"] = "Door"
			door["type"] = "Unlocked Door"
			d.Cells[doorR][doorC].Label = "o"
		case doorType < 75:
			d.Cells[doorR][doorC].Locked = true
			door["key"] = "Locked"
			door["type"] = "Locked Door"
			d.Cells[doorR][doorC].Label = "x"
		case doorType < 90:
			d.Cells[doorR][doorC].Trapped = true
			door["key"] = "Trapped"
			door["type"] = "Trapped Door"
			d.Cells[doorR][doorC].Label = "t"
		case doorType < 100:
			d.Cells[doorR][doorC].Secret = true
			door["key"] = "Secret"
			door["type"] = "Secret Door"
			d.Cells[doorR][doorC].Label = "s"
		default:
			d.Cells[doorR][doorC].Portc = true
			door["key"] = "Portc"
			door["type"] = "Portcullis"
			d.Cells[doorR][doorC].Label = "#"
		}
		if outID != 0 {
			door["outID"] = outID
		}
		if _, ok := room.Doors[openDir]; !ok {
			room.Doors[openDir] = []map[string]any{}
		}
		room.Doors[openDir] = append(room.Doors[openDir], door)
		//
		i++
	}
}

func (d *DumbDungeon) allocOpens(room *RoomData) int {
	roomH := (room.S-room.N)/2 + 1
	roomW := (room.E-room.W)/2 + 1
	flumph := int(math.Sqrt(float64(roomH) * float64(roomW)))
	return flumph + rand.Intn(flumph)
}

func (d *DumbDungeon) DoorSills(room *RoomData) []map[string]any {
	doorSills := make([]map[string]any, 0)
	if room.N >= 3 {
		for c := room.W; c <= room.E; c += 2 {
			doorSill := checkSill(d.Cells, room, room.N, c, "N")
			if doorSill != nil {
				doorSills = append(doorSills, doorSill)
			}
		}
	}
	if room.S <= d.NRows-3 {
		for c := room.W; c <= room.E; c += 2 {
			doorSill := checkSill(d.Cells, room, room.S, c, "S")
			if doorSill != nil {
				doorSills = append(doorSills, doorSill)
			}
		}
	}
	if room.W >= 3 {
		for r := room.N; r <= room.S; r += 2 {
			doorSill := checkSill(d.Cells, room, r, room.W, "W")
			if doorSill != nil {
				doorSills = append(doorSills, doorSill)
			}
		}
	}
	if room.E <= d.NCols-3 {
		for r := room.N; r <= room.S; r += 2 {
			doorSill := checkSill(d.Cells, room, r, room.E, "E")
			if doorSill != nil {
				doorSills = append(doorSills, doorSill)
			}
		}
	}
	return doorSills
}

func checkSill(cells [][]*Cell, room *RoomData, r, c int, dir string) map[string]any {
	doorR := r + directionI[dir]
	doorC := c + directionJ[dir]
	doorCell := cells[doorR][doorC]
	if !doorCell.Perimeter || doorCell.BlockDoor() {
		return nil
	}
	outR := doorR + directionI[dir]
	outC := doorC + directionJ[dir]
	outCell := cells[outR][outC]
	if outCell.Blocked {
		return nil
	}
	if outCell.Room && outCell.ID == room.ID {
		return nil
	}
	return map[string]any{
		"sillR": r,
		"sillC": c,
		"doorR": doorR,
		"doorC": doorC,
		"outID": outCell.ID,
		"dir":   dir,
	}
}

func (d *DumbDungeon) LabelRooms() {
	for id := 1; id <= d.NRooms; id++ {
		room := d.Rooms[id]
		label := fmt.Sprintf("%d", room.ID)
		labelLen := len(label)
		labelRow := room.N
		labelCol := room.W

		for c := 0; c < labelLen; c++ {
			char := string(label[c])
			d.Cells[labelRow][labelCol+c].Label = char
		}
	}
}

func (d *DumbDungeon) Corridors() {
	for i := 1; i < d.Ni; i++ {
		r := i*2 + 1
		for j := 1; j < d.Nj; j++ {
			c := j*2 + 1
			if d.Cells[r][c].Corridor {
				continue
			}
			d.tunnel(i, j, "")
		}
	}
}

func (d *DumbDungeon) tunnel(i, j int, lastDir string) {
	dirs := d.tunnelDirs(lastDir)
	for _, dir := range dirs {
		if d.openTunnel(i, j, dir) {
			nextI := i + directionI[dir]
			nextJ := j + directionJ[dir]
			d.tunnel(nextI, nextJ, dir)
		}
	}
}

func (d *DumbDungeon) tunnelDirs(lastDir string) []string {
	p := corridorLayouts[d.Seed.CorridorLayout]
	dirs := []string{"E", "N", "S", "W"}
	rand.Shuffle(len(dirs), func(i, j int) {
		dirs[i], dirs[j] = dirs[j], dirs[i]
	})
	if lastDir != "" && rand.Intn(100) < p {
		dirs = append([]string{lastDir}, dirs...)
	}
	return dirs
}

func (d *DumbDungeon) openTunnel(i, j int, dir string) bool {
	thisR := i*2 + 1
	thisC := j*2 + 1
	nextR := (i+directionI[dir])*2 + 1
	nextC := (j+directionJ[dir])*2 + 1
	midR := (thisR + nextR) / 2
	midC := (thisC + nextC) / 2
	if d.soundTunnel(midR, midC, nextR, nextC) {
		return d.delveTunnel(thisR, thisC, nextR, nextC)
	}
	return false
}

func (d *DumbDungeon) soundTunnel(midR, midC, nextR, nextC int) bool {
	if nextR < 0 || nextC < 0 || nextR > d.NRows || nextC > d.NCols {
		return false
	}
	r1, r2 := midR, nextR
	if midR > nextR {
		r1, r2 = nextR, midR
	}
	c1, c2 := midC, nextC
	if midC > nextC {
		c1, c2 = nextC, midC
	}
	for r := r1; r <= r2; r++ {
		for c := c1; c <= c2; c++ {
			if d.Cells[r][c].BlockCorridor() {
				return false
			}
		}
	}
	return true
}

func (d *DumbDungeon) delveTunnel(thisR, thisC, nextR, nextC int) bool {
	r1, r2 := thisR, nextR
	if thisR > nextR {
		r1, r2 = nextR, thisR
	}
	c1, c2 := thisC, nextC
	if thisC > nextC {
		c1, c2 = nextC, thisC
	}
	for r := r1; r <= r2; r++ {
		for c := c1; c <= c2; c++ {
			d.Cells[r][c].Entrance = false
			d.Cells[r][c].Corridor = true
			if d.Cells[r][c].Label == "" {
				d.Cells[r][c].Label = "█"
			}
		}
	}
	return true
}

func (d *DumbDungeon) EmplaceStairs() {
	n := d.Seed.AddStairs
	list := d.stairEnds()
	if len(list) == 0 {
		return
	}

	for i := 0; i < n; i++ {
		stairs, _ := utils.Splice(&list, rand.Intn(len(list)), 1)
		if len(stairs) == 0 {
			break
		}
		stair := stairs[0]
		r := stair["row"].(int)
		c := stair["col"].(int)
		typ := i
		if i >= 2 {
			typ = rand.Intn(2)
		}
		if typ == 0 {
			d.Cells[r][c].StairDown = true
			d.Cells[r][c].Label = "d"
			stair["type"] = "down"
		} else {
			d.Cells[r][c].StairUp = true
			d.Cells[r][c].Label = "u"
			stair["type"] = "up"
		}
		d.Stairs = append(d.Stairs, stair)
	}
}

func (d *DumbDungeon) stairEnds() []map[string]any {
	list := []map[string]any{}
	for i := 0; i < d.Ni; i++ {
		r := i*2 + 1
		for j := 0; j < d.Nj; j++ {
			c := j*2 + 1
			if !d.Cells[r][c].IsOnlyCorridor() || d.Cells[r][c].Stairs() {
				continue
			}
			for _, stairs := range stairEnd {
				if d.checkTunnel(r, c, stairs) {
					end := map[string]any{"row": r, "col": c}
					n := stairs["next"].([]int)
					end["nextRow"] = end["row"].(int) + n[0]
					end["nextCol"] = end["col"].(int) + n[1]
					list = append(list, end)
					break
				}
			}
		}
	}
	return list
}

func (d *DumbDungeon) CleanDungeon() {
	if d.Seed.RemoveDeadEnds > 0 {
		d.removeDeadEnds()
	}
	d.fixDoors()
	d.emptyBlocks()
}

func (d *DumbDungeon) removeDeadEnds() {
	all := d.Seed.RemoveDeadEnds == 100

	for i := 0; i < d.Ni; i++ {
		r := i*2 + 1
		for j := 0; j < d.Nj; j++ {
			c := j*2 + 1
			if d.Cells[r][c].Stairs() || !(all || rand.Intn(100) < d.Seed.RemoveDeadEnds) {
				continue
			}
			d.collapse(r, c)
		}
	}
}

func (d *DumbDungeon) collapse(r, c int) {
	if !d.Cells[r][c].OpenSpace() {
		return
	}

	for _, value := range closeEnd {
		if d.checkTunnel(r, c, value) {
			for _, p := range value["close"].([][]int) {
				d.Cells[r+p[0]][c+p[1]] = NewCell()
			}
			v := value["recurse"].([]int)
			d.collapse(r+v[0], c+v[1])
		}
	}
}

func (d *DumbDungeon) checkTunnel(r, c int, check map[string]any) bool {
	if list, ok := check["corridor"].([][]int); !ok || list != nil {
		for _, coord := range list {
			if !d.Cells[r+coord[0]][c+coord[1]].IsOnlyCorridor() {
				return false
			}
		}
	}
	if list := check["walled"].([][]int); list != nil {
		for _, coord := range list {
			if d.Cells[r+coord[0]][c+coord[1]].OpenSpace() {
				return false
			}
		}
	}
	return true
}

func (d *DumbDungeon) fixDoors() {
	fixed := initCoords(d.MaxRow+1, d.MaxCol+1)
	for _, room := range d.Rooms {
		for dir, doors := range room.Doors {
			shiny := []map[string]any{}
			for _, door := range doors {
				doorR := door["row"].(int)
				doorC := door["col"].(int)
				doorCell := d.Cells[doorR][doorC]
				if !doorCell.OpenSpace() {
					d.Cells[doorR][doorC] = NewCell()
					continue
				}
				if fixed[doorR][doorC] != 0 {
					shiny = append(shiny, door)
					continue
				}
				if door["outID"] != nil {
					outID := door["outID"].(int)
					if outID == 0 {
						continue
					}
					outDir := opposite[dir]
					if _, ok := d.Rooms[outID]; !ok {
						continue
					}
					if _, ok := d.Rooms[outID].Doors[outDir]; !ok {
						d.Rooms[outID].Doors[outDir] = []map[string]any{}
					}
					d.Rooms[outID].Doors[outDir] = append(d.Rooms[outID].Doors[outDir], door)
				}
				shiny = append(shiny, door)
				fixed[doorR][doorC]++
			}
			if len(shiny) == 0 {
				if _, ok := room.Doors[dir]; ok {
					room.Doors[dir] = []map[string]any{}
				}
				room.Doors[dir] = shiny
				d.Doors = append(d.Doors, shiny...)
			} else {
				delete(room.Doors, dir)
			}
		}
	}
}

func (d *DumbDungeon) emptyBlocks() {
	for r := 0; r < d.NRows; r++ {
		for c := 0; c < d.NCols; c++ {
			if d.Cells[r][c].Blocked {
				d.Cells[r][c] = NewCell()
			}
		}
	}
}
