package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const VERBOSE = false
const START = "S"
const END = "E"
const EDGE_LENGTH = 1
const INFINITY = 999999
const POINT_DELIMITER = ","

// Move represents a move in the map.
type Move int

const (
	Up Move = iota
	Down
	Left
	Right
)

// String returns the string representation of the move.
func (m Move) String() string {
	switch m {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}

	return "Unknown"
}

// Heights represents the height of the map.
type Heights int

const (
	a Heights = iota
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	l
	m
	n
	o
	p
	q
	r
	s
	t
	u
	v
	w
	x
	y
	z
)

// String returns the string representation of the heightmap.
func (h Heights) String() string {
	return string('a' + h)
}

// convertToHeightmap converts a rune to a heightmap.
func convertToHeightmap(c rune) Heights {
	return Heights(c - 'a')
}

// canClimb returns true if the current height can climb to the next height.
func canClimb(current, next Heights) bool {
	diff := next - current

	return diff <= EDGE_LENGTH
}

// canGoDown returns true if the current height can go down to the next height.
func canGoDown(current, next Heights) bool {
	diff := current - next

	return diff <= EDGE_LENGTH
}

// Point represents a point in the map.
type Point struct {
	x        int
	y        int
	distance int
}

// Mapheight represents a map of heights.
type Mapheight struct {
	grid      [][]Heights
	points    [][]*Point
	visited   [][]bool
	unvisited [][]bool
	start     *Point
	end       *Point
	current   *Point
	route     []*Point
	testFn    func(current, next Heights) bool
}

// newMapheight creates a new mapheight.
func newMapheight(lines []string) *Mapheight {
	numRows := len(lines)
	numCols := len(lines[0])
	grid := make([][]Heights, numRows)
	points := make([][]*Point, numRows)
	visited := make([][]bool, numRows)
	unvisited := make([][]bool, numRows)
	// route has a maximum size of numRows * numCols
	route := make([]*Point, 0, numRows * numCols)

	for i := 0; i < numRows; i++ {
		grid[i] = make([]Heights, numCols)
		points[i] = make([]*Point, numCols)
		visited[i] = make([]bool, numCols)
		unvisited[i] = make([]bool, numCols)
	}

	m := &Mapheight{
		grid:      grid,
		points:    points,
		visited:   visited,
		unvisited: unvisited,
		start:     &Point{},
		end:       &Point{},
		current:   &Point{},
		route:     route,
		testFn:    canClimb,
	}

	m.addGrid(lines)

	return m
}

// sclear clears the map.
func (m *Mapheight) clear() {
	m.grid = nil
	m.points = nil
	m.visited = nil
	m.unvisited = nil
	m.start = nil
	m.end = nil
	m.current = nil
	m.testFn = nil
}

// setHeight sets the height of the given position.
func (m *Mapheight) setHeight(x, y int, h Heights) {
	m.grid[y][x] = h
}

// getHeight returns the height of the given position.
func (m *Mapheight) getHeight(x, y int) Heights {
	return m.grid[y][x]
}

// setPoint sets the point of the given position.
func (m *Mapheight) setPoint(x, y int, p *Point) {
	m.points[y][x] = p
}

// getPoint returns the point of the given position.
func (m *Mapheight) getPoint(x, y int) *Point {
	return m.points[y][x]
}

// canMove returns true if the move is possible, considering the current position and the surrounding height.
func (m *Mapheight) canMove(move Move) bool {
	currentHeight := m.getHeight(m.current.x, m.current.y)
	var nextHeight Heights

	switch move {
	case Up:
		if m.current.y == 0 {
			return false
		}

		nextHeight = m.getHeight(m.current.x, m.current.y-1)
	case Down:
		if m.current.y == len(m.grid) - 1 {
			return false
		}

		nextHeight = m.getHeight(m.current.x, m.current.y+1)
	case Left:
		if m.current.x == 0 {
			return false
		}

		nextHeight = m.getHeight(m.current.x-1, m.current.y)
	case Right:
		if len(m.grid) > 0 && m.current.x == len(m.grid[0]) - 1 {
			return false
		}

		nextHeight = m.getHeight(m.current.x+1, m.current.y)
	}

	return m.testFn(currentHeight, nextHeight)
}

// move moves the current position in the map.
func (m *Mapheight) move(move Move) (int, int, error) {
	if !m.canMove(move) {
		return -1, -1, errors.New("can't move")
	}

	switch move {
	case Up:
		return m.current.x, m.current.y - 1, nil
	case Down:
		return m.current.x, m.current.y + 1, nil
	case Left:
		return m.current.x - 1, m.current.y, nil
	case Right:
		return m.current.x + 1, m.current.y, nil
	}

	return -1, -1, errors.New("can't move")
}

// moveTo return the new point.
func (m *Mapheight) moveTo(move Move) *Point {
	x, y, err := m.move(move)

	if err != nil {
		return nil
	}

	return m.getPoint(x, y)
}

// addPoint adds a point to the map.
func (m *Mapheight) addPoint(x, y int) {
	point := &Point{x, y, INFINITY}
	m.setPoint(x, y, point)
}

// markUnvisited marks the current position as unvisited.
func (m *Mapheight) markUnvisited(x, y int) {
	m.unvisited[y][x] = true

	if m.visited[y][x] {
		m.visited[y][x] = false
	}
}

// markVisited marks the current position as visited.
func (m *Mapheight) markVisited(x, y int) {
	m.visited[y][x] = true

	if m.unvisited[y][x] {
		m.unvisited[y][x] = false
	}

	if VERBOSE {
		fmt.Printf("Marked (%d%s%d) as visited\n", x, POINT_DELIMITER, y)
	}
}

// isVisited returns true if the current position has been visited.
func (m *Mapheight) isVisited(x, y int) bool {
	return m.visited[y][x]
}

// getNeighbours returns the neighbours of the current position.
func (m *Mapheight) getNeighbours() []*Point {
	// We can move up, down, left or right. At most we can have 4 neighbours.
	neighbours := make([]*Point, 0, 4)

	for _, move := range []Move{Up, Down, Left, Right} {
		if !m.canMove(move) {
			continue
		}

		neighbour := m.moveTo(move)

		if neighbour != nil {
			neighbours = append(neighbours, neighbour)
		}
	}

	return neighbours
}

// getUnvisitedNeighbours returns the neighbours of the current position that have not been visited.
func (m *Mapheight) getUnvisitedNeighbours() []*Point {
	// We can move up, down, left or right. At most we can have 4 neighbours.
	neighbours := make([]*Point, 0, 4)

	for _, move := range []Move{Up, Down, Left, Right} {
		if !m.canMove(move) {
			continue
		}

		neighbour := m.moveTo(move)

		if neighbour != nil && !m.isVisited(neighbour.x, neighbour.y) {
			neighbours = append(neighbours, neighbour)
		}
	}

	return neighbours
}

// getSmallestUnvisited returns the smallest unvisited point.
func (m *Mapheight) getSmallestUnvisited() *Point {
	smallest := INFINITY
	var smallestPoint *Point

	for y := 0; y < len(m.points); y++ {
		for x := 0; x < len(m.points[y]); x++ {
			if m.isVisited(x, y) {
				continue
			}

			point := m.getPoint(x, y)

			if point.distance < smallest {
				smallest = point.distance
				smallestPoint = point
			}
		}
	}


	return smallestPoint
}

// findPathAlgorithm finds the path from the start to the end using Dijkstra's algorithm.
func (m *Mapheight) findPathAlgorithm() {
	currentNeighbours := m.getUnvisitedNeighbours()

	for _, neighbour := range currentNeighbours {
		tentativeDistance := m.current.distance + EDGE_LENGTH

		if neighbour.distance > tentativeDistance {
			neighbour.distance = tentativeDistance
		}
	}

	m.markVisited(m.current.x, m.current.y)

	destination := m.getSmallestUnvisited()

	if destination == nil {
		if VERBOSE {
			fmt.Println("No destination found")
		}
		return
	}

	if m.isVisited(destination.x, destination.y) || destination.distance == INFINITY {
		if VERBOSE {
			fmt.Println("Destination is visited or has no distance")
		}
		return
	}

	if destination.x == m.end.x && destination.y == m.end.y {
		if VERBOSE {
			fmt.Println("Destination is the end")
		}
		return
	}

	m.current = destination

	m.findPathAlgorithm()
}

// findPathRoute finds the shortest path from the start to the end using the points' distance.
func (m *Mapheight) findPathRoute() {
	m.route = append(m.route, m.end)
	m.current = m.end

	// find the smallest neighbour of the current point
	smallest := INFINITY
	var smallestPoint *Point

	for _, neighbour := range m.getNeighbours() {
		if neighbour.distance < smallest {
			smallest = neighbour.distance
			smallestPoint = neighbour
		}
	}

	m.current = smallestPoint

	for {
		for _, neighbour := range m.getNeighbours() {
			if m.current.distance == neighbour.distance + EDGE_LENGTH {
				m.current = neighbour
				m.route = append(m.route, neighbour)

				if VERBOSE {
					fmt.Printf("Added (%d%s%d) to the route\n", neighbour.x, POINT_DELIMITER, neighbour.y)
				}

				break
			}
		}

		if m.current.x == m.start.x && m.current.y == m.start.y {
			break
		}
	}
}

// findPath finds the path from the start to the end using Dijkstra's algorithm.
func (m *Mapheight) findPath() {
	m.start.distance = 0
	m.current = m.start
	m.markVisited(m.end.x, m.end.y)

	m.findPathAlgorithm()

	m.testFn = canGoDown

	m.findPathRoute()
}

// getPossibleStartingPoints returns the possible starting points.
func (m *Mapheight) getPossibleStartingPoints() []*Point {
	startingPoints := make([]*Point, 0)

	for y := 0; y < len(m.points); y++ {
		for x := 0; x < len(m.points[y]); x++ {
			point := m.getPoint(x, y)

			if m.getHeight(point.x, point.y) == a {
				startingPoints = append(startingPoints, point)
			}
		}
	}

	return startingPoints
}

// addGrid adds a grid to the map.
func (m *Mapheight) addGrid(lines []string) {
	for y, line := range lines {
		for x, c := range line {
			m.addPoint(x, y)
			m.markUnvisited(x, y)

			if string(c) == START {
				m.start = m.getPoint(x, y)
				// start is always at the lowest point
				m.setHeight(x, y, a)
			} else if string(c) == END {
				m.end = m.getPoint(x, y)
				// end is always at the highest point
				m.setHeight(x, y, z)
			} else {
				m.setHeight(x, y, convertToHeightmap(c))
			}
		}
	}
}

// String returns the string representation of the map.
func (m *Mapheight) String() string {
	result := ""

	for _, row := range m.grid {
		for _, c := range row {
			result += c.String()
		}

		result += "\n"
	}

	result += "\n"

	result += fmt.Sprintf("Start: %v\n", m.start)
	result += fmt.Sprintf("End: %v\n", m.end)
	result += fmt.Sprintf("Current: %v\n", m.current)

	return result
}

// getMapheightsForStartingPoints returns the mapheights for the starting points.
func getMapheightsForStartingPoints(lines []string, startingPoints []*Point) []int {
	numPoints := len(startingPoints)
	routeLengths := make([]int, numPoints)

	for i, startingPoint := range startingPoints {
		mapheight := newMapheight(lines)
		point := mapheight.getPoint(startingPoint.x, startingPoint.y)
		mapheight.start = point

		mapheight.findPath()
		mapheight.clear()

		routeLengths[i] = len(mapheight.route)
	}

	return routeLengths
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	var startingPoint string
	if len(args) > 1 {
		startingPoint = args[1]
	}
	txtlines := helpers.ReadFile(filename)

	// part 1
	mapheight := newMapheight(txtlines)
	mapheight.findPath()
	if startingPoint == "" {
		fmt.Printf(
			"[Part One] The answer is: %d\n",
			len(mapheight.route),
		)
		fmt.Println("[Part Two] The answer can be obtained with the following command:")
		fmt.Printf("\n\tgo run . %s all\n\n", filename)
	}

	// part 2
	var possibleStartingPoints []*Point
	if startingPoint == "all" {
		possibleStartingPoints = mapheight.getPossibleStartingPoints()
		startingPoints := make([]string, len(possibleStartingPoints))
		for i, point := range possibleStartingPoints {
			startingPoints[i] = fmt.Sprintf("%d%s%d", point.x, POINT_DELIMITER, point.y)
		}
		fmt.Print(strings.Join(startingPoints, "\n"))
	} else if startingPoint != "" {
		// split and parse the starting point
		xy := strings.Split(startingPoint, POINT_DELIMITER)
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		point := mapheight.getPoint(x, y)
		possibleStartingPoints = []*Point{point}
		routeLengths := getMapheightsForStartingPoints(txtlines, possibleStartingPoints)
		shortestRoute := helpers.MinOf(routeLengths...)
		fmt.Println(shortestRoute)
	}
	}
