package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const VERBOSE = false
const INFINITY = int(^uint(0) >> 1)
const POINT_DELIMITER = ","
const FLOOR_LEVEL = 2
const FLOOR_INFINITY = 9999

// Element is the type of the element in the cave.
type Element int

const (
	Rock Element = iota
	Air
	Sand
	SandSource
)

// String returns the string representation of the element.
func (e Element) String() string {
	return [...]string{"Rock", "Air", "Sand", "SandSource"}[e]
}

// Point is a point in the cave.
type Point struct {
	x, y    int
	element Element
}

// shouldStop returns true if the point should stop the sand.
func (p *Point) shouldStop() bool {
	if p.element == Rock {
		return true
	}

	if p.element == Sand {
		return true
	}

	if p.element == SandSource {
		return true
	}

	return false
}

// String returns the string representation of the point.
func (p *Point) String() string {
	switch p.element {
	case Rock:
		return "#"
	case Air:
		return "."
	case Sand:
		return "o"
	case SandSource:
		return "+"
	default:
		return " "
	}
}

// Cave is the cave.
type Cave struct {
	grid                   [][]*Point
	xMin, xMax, yMin, yMax int
	sandSource             *Point
	numFallenSand          int
}

// isSandSource returns true if the point is a sand source.
func (c *Cave) isSandSouce(point *Point) bool {
	return point.element == SandSource
}

// isAbyss returns true if the point is an abyss.
func (c *Cave) isAbyss(x, y int) bool {
	return y > c.yMax
}

// exists returns true if the point exists.
func (c *Cave) exists(x, y int) bool {
	if x < c.xMin || x > c.xMax || y < c.yMin || y > c.yMax {
		return true
	}

	return false
}

// getPoint returns the point at the given coordinates.
func (c *Cave) getPoint(x, y int) *Point {
	if c.exists(x, y) {
		return nil
	}

	return c.grid[y-c.yMin][x-c.xMin]
}

// addPoint adds a point to the cave.
func (c *Cave) addPoint(point *Point) {
	x := point.x - c.xMin
	y := point.y - c.yMin
	c.grid[y][x] = point
}

// addSandSource adds a sand source to the cave.
func (c *Cave) addSandSource(point *Point) {
	c.addPoint(point)
	c.sandSource = point
}

// addRock adds a rock to the cave.
func (c *Cave) addRockPath(rockPath []*Point) {
	for i := 0; i < len(rockPath)-1; i++ {
		firstPoint := rockPath[i]
		secondPoint := rockPath[i+1]

		for x := firstPoint.x; x <= secondPoint.x; x++ {
			c.addPoint(&Point{
				x:       x,
				y:       firstPoint.y,
				element: Rock,
			})
		}

		for y := firstPoint.y; y <= secondPoint.y; y++ {
			c.addPoint(&Point{
				x:       secondPoint.x,
				y:       y,
				element: Rock,
			})
		}

		for x := secondPoint.x; x <= firstPoint.x; x++ {
			c.addPoint(&Point{
				x:       x,
				y:       secondPoint.y,
				element: Rock,
			})
		}

		for y := secondPoint.y; y <= firstPoint.y; y++ {
			c.addPoint(&Point{
				x:       secondPoint.x,
				y:       y,
				element: Rock,
			})
		}
	}
}

// fillWithAir fills the cave with air.
func (c *Cave) fillWithAir() {
	for y := c.yMin; y <= c.yMax; y++ {
		for x := c.xMin; x <= c.xMax; x++ {
			c.addPoint(&Point{
				x:       x,
				y:       y,
				element: Air,
			})
		}
	}
}

// shouldStopAtPoint returns true if the sand should stop at the given point.
func (c *Cave) shouldStopAtPoint(x, y int) (bool, bool) {
	point := c.getPoint(x, y)

	if c.isAbyss(x, y) {
		return true, false
	}

	if point == nil || !point.shouldStop() {
		return false, true
	}

	return false, false
}

// fillWithSandFromPoint fills the cave with sand from the given point.
func (c *Cave) fillWithSandFromPoint(startPoint *Point) bool {
	x := startPoint.x
	for y := startPoint.y; y <= c.yMax; y++ {
		// Move down
		stop, fall := c.shouldStopAtPoint(x, y)
		if stop {
			return false
		}
		if fall {
			continue
		}

		// Move left and down
		x -= 1
		stop, fall = c.shouldStopAtPoint(x, y)
		if stop {
			return false
		}
		if fall {
			continue
		}

		// Move right and down
		x += 2
		stop, fall = c.shouldStopAtPoint(x, y)
		if stop {
			return false
		}
		if fall {
			continue
		}

		// Stop falling
		currentPoint := c.getPoint(x-1, y-1)
		c.numFallenSand++
		if c.isSandSouce(currentPoint) {
			currentPoint.element = Sand
			return false
		}
		currentPoint.element = Sand
		return true
	}

	return false
}

// fillWithSand fills the cave with sand.
func (c *Cave) fillWithSand() {
	x := c.sandSource.x
	y := c.sandSource.y + 1
	startPoint := c.getPoint(x, y)
	do := c.fillWithSandFromPoint(startPoint)

	for do {
		do = c.fillWithSandFromPoint(startPoint)

		if VERBOSE {
			fmt.Println(c.numFallenSand)
			fmt.Println(c)
		}
	}
}

// String returns the string representation of the cave.
func (c *Cave) String() string {
	str := ""
	for _, row := range c.grid {
		for _, point := range row {
			if point == nil {
				str += " "
				continue
			}
			str += point.String()
		}
		str += "\n"
	}
	return str
}

// newCave creates a new cave.
func newCave(xMin, xMax, yMin, yMax int) *Cave {
	c := Cave{
		xMin: xMin,
		xMax: xMax,
		yMin: yMin,
		yMax: yMax,
	}

	c.grid = make([][]*Point, c.yMax-c.yMin+1)

	for i := range c.grid {
		c.grid[i] = make([]*Point, c.xMax-c.xMin+1)
	}

	c.fillWithAir()

	return &c
}

// getRockPath returns the rock path.
func getRockPath(rockPathPoints [][]string) [][]*Point {
	rockPath := make([][]*Point, len(rockPathPoints))

	for i, points := range rockPathPoints {
		rockPath[i] = make([]*Point, len(points))

		for j, pointString := range points {
			point := strings.Split(pointString, POINT_DELIMITER)
			x, _ := strconv.Atoi(point[0])
			y, _ := strconv.Atoi(point[1])
			rockPath[i][j] = &Point{
				x:       x,
				y:       y,
				element: Rock,
			}
		}
	}

	return rockPath
}

// getMinMaxCoords returns the min and max coordinates of the rock path.
func getMinMaxCoords(rockPathPoints [][]*Point) (int, int, int, int) {
	xMin := INFINITY
	xMax := 0
	yMin := INFINITY
	yMax := 0

	for _, points := range rockPathPoints {
		for _, point := range points {
			if point.x < xMin {
				xMin = point.x
			}

			if point.x > xMax {
				xMax = point.x
			}

			if point.y < yMin {
				yMin = point.y
			}

			if point.y > yMax {
				yMax = point.y
			}
		}
	}

	return xMin, xMax, yMin, yMax
}

// getPointsFromFile returns the points from the file.
func getPointsFromFile(txtlines []string) [][]string {
	points := make([][]string, len(txtlines))

	for i, line := range txtlines {
		re := regexp.MustCompile(`(\d+),(\d+)`)
		matches := re.FindAllString(line, -1)
		points[i] = matches
	}

	return points
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// create the sand source
	sandSource := &Point{
		x:       500,
		y:       0,
		element: SandSource,
	}

	// part 1
	rockPathPoints := getPointsFromFile(txtlines)
	rockPaths := getRockPath(rockPathPoints)
	rockPathsWithSandSource := append(rockPaths, []*Point{sandSource})
	xMin, xMax, yMin, yMax := getMinMaxCoords(rockPathsWithSandSource)
	cave := newCave(xMin, xMax, yMin, yMax)
	for _, rockPath := range rockPaths {
		cave.addRockPath(rockPath)
	}
	cave.addSandSource(sandSource)
	cave.fillWithSand()
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		cave.numFallenSand,
	)

	// part 2
	rockPathPoints = getPointsFromFile(txtlines)
	rockPaths = getRockPath(rockPathPoints)
	rockPathsWithSandSource = append(rockPaths, []*Point{sandSource})
	xMin, xMax, yMin, yMax = getMinMaxCoords(rockPathsWithSandSource)
	xMin -= FLOOR_INFINITY
	xMax += FLOOR_INFINITY
	yMax += FLOOR_LEVEL
	floor := []*Point{
		&Point{
			x:       xMin,
			y:       yMax,
			element: Rock,
		},
		&Point{
			x:       xMax,
			y:       yMax,
			element: Rock,
		},
	}
	cave = newCave(xMin, xMax, yMin, yMax)
	for _, rockPath := range rockPaths {
		cave.addRockPath(rockPath)
	}
	cave.addRockPath(floor)
	cave.addSandSource(sandSource)
	cave.fillWithSand()
	fmt.Printf(
		"[Part Two] The answer is: %d\n",
		cave.numFallenSand,
	)
}
