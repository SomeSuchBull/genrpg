package dungeon

import (
	"container/heap"
	"math"
)

type Node struct {
	point             Point
	distance          int
	previousDirection Point
}

const (
	blockedDijkstraWeight = -1
	neutralDijkstraWeight = 0
	normalDijkstraWeight  = 1
)

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

func isValid(x, y, rows, cols, wrongWeight int, grid [][]int) bool {
	if x < 0 || y < 0 || y >= rows || x >= cols {
		return false
	}
	if grid[y][x] == blockedDijkstraWeight || grid[y][x] == wrongWeight {
		return false
	}
	return true
}

func dijkstraFindNearest(grid [][]int, start Point, startingDijkstraWeight int) (int, Point, []Point) {
	rows := len(grid)
	cols := len(grid[0])
	directions := []Point{N, E, W, S}
	oppositeDirections := map[Point]Point{N: S, E: W, W: E, S: N}

	distance := make([][]int, rows)
	previous := make([][]Point, rows)
	for i := range distance {
		distance[i] = make([]int, cols)
		previous[i] = make([]Point, cols)
		for j := range distance[i] {
			distance[i][j] = math.MaxInt32
			previous[i][j] = Point{-1, -1}
		}
	}
	distance[start.Y][start.X] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{start, 0, Point{-1, -1}})

	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)
		currentPoint := node.point
		currentDist := node.distance
		// previousDir := node.previousDirection

		if grid[currentPoint.Y][currentPoint.X] != startingDijkstraWeight &&
			grid[currentPoint.Y][currentPoint.X] != normalDijkstraWeight &&
			grid[currentPoint.Y][currentPoint.X] != blockedDijkstraWeight &&
			currentPoint != start {
			path := []Point{}
			for p := currentPoint; p != (Point{-1, -1}); p = previous[p.Y][p.X] {
				path = append([]Point{p}, path...)
			}
			return currentDist, currentPoint, path
		}

		for _, dir := range directions {
			if dir == oppositeDirections[node.previousDirection] {
				continue
			}
			nx, ny := currentPoint.X+dir.X, currentPoint.Y+dir.Y

			if isValid(nx, ny, rows, cols, startingDijkstraWeight, grid) {
				newDist := currentDist + grid[ny][nx]
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
