package dungeon

import (
	"container/heap"
	"math"
	"slices"
)

type Node struct {
	point             Point
	distance          int
	previousDirection Point
}

const (
	blockedDijkstraWeight  = -1
	neutralDijkstraWeight  = 0
	normalDijkstraWeight   = 1
	corridorDijkstraWeight = 2
)

var N = Point{X: 0, Y: -1}
var NE = Point{X: 1, Y: -1}
var E = Point{X: 1, Y: 0}
var SE = Point{X: 1, Y: 1}
var S = Point{X: 0, Y: 1}
var SW = Point{X: -1, Y: 1}
var W = Point{X: -1, Y: 0}
var NW = Point{X: -1, Y: -1}

var directions = []Point{N, W, E, S}
var dirMap = map[Point]int{N: 1, W: 2, E: 3, S: 4}
var oppositeDirections = map[Point]Point{N: S, E: W, W: E, S: N}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func isValid(x, y, rows, cols, wrongWeight, doorID int, grid [][]Cell) bool {
	if x < 0 || y < 0 || y >= rows || x >= cols {
		return false
	}
	if grid[y][x].PathValue == blockedDijkstraWeight || grid[y][x].ID == wrongWeight {
		return false
	}
	if grid[y][x].DoorPerimeter.DoorID != doorID && grid[y][x].DoorPerimeter.RoomID == wrongWeight {
		return false
	}
	return true
}

func dijkstraFindNearest(d BSPDungeon, start Point, originID, doorID int) (int, Point, []Point) {
	grid := d.ExpandedGrid
	rows := len(grid)
	cols := len(grid[0])

	distance := make([][]int, rows)
	previous := make([][]Point, rows)
	for i := range distance {
		// for i := range previous {
		distance[i] = make([]int, cols)
		previous[i] = make([]Point, cols)
		for j := range distance[i] {
			// for j := range previous[i] {
			distance[i][j] = math.MaxInt32
			previous[i][j] = Point{-1, -1}
		}
	}
	distance[start.Y][start.X] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{start, 0, Point{-1, -1}})

	starting := true
	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)
		currentPoint := node.point
		currentDist := node.distance
		// previousDir := node.previousDirection

		if !starting && grid[currentPoint.Y][currentPoint.X].PathValue == originID+1 {
			continue
		}
		starting = false

		if grid[currentPoint.Y][currentPoint.X].ID != originID &&
			grid[currentPoint.Y][currentPoint.X].PathValue != normalDijkstraWeight &&
			grid[currentPoint.Y][currentPoint.X].PathValue != blockedDijkstraWeight &&
			currentPoint != start {
			if !d.Corridors[grid[currentPoint.Y][currentPoint.X].CorridorID][originID] &&
				!slices.Contains(d.Rooms[originID].ConnectedRooms, d.ExpandedGrid[currentPoint.Y][currentPoint.X].ID) {
				path := []Point{}
				for p := currentPoint; p != (Point{-1, -1}); p = previous[p.Y][p.X] {
					path = append([]Point{p}, path...)
				}
				return currentDist, currentPoint, path
			} else {
				continue
			}
		}

		for _, dir := range directions {
			if dir == oppositeDirections[node.previousDirection] {
				continue
			}
			nx, ny := currentPoint.X+dir.X, currentPoint.Y+dir.Y

			if isValid(nx, ny, rows, cols, originID, doorID, grid) {
				newDist := currentDist + 1
				if newDist < distance[ny][nx] {
					distance[ny][nx] = newDist
					previous[ny][nx] = currentPoint
					heap.Push(pq, Node{Point{nx, ny}, newDist, dir})
				}
			}
		}
	}

	return -1, Point{-1, -1}, nil // If no target point exists
}
