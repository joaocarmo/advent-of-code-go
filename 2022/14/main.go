package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const VERBOSE = true
const INFINITY = int(^uint(0) >> 1)
const POINT_DELIMITER = ","

type Element int

const (
	Rock Element = iota
	Air
	Sand
	SandSource
)

func (e Element) String() string {
	return [...]string{"Rock", "Air", "Sand", "SandSource"}[e]
}

type Point struct {
	x, y    int
	element Element
}

func (p *Point) shouldStop() bool {
	if p.element == Rock {
		return true
	}

	if p.element == Sand {
		return true
	}

	return false
}

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

type Cave struct {
	grid                   [][]*Point
	xMin, xMax, yMin, yMax int
	sandSource             *Point
	numFallenSand          int
}

func (c *Cave) isAbyss(x, y int) bool {
	return y > c.yMax
}

func (c *Cave) exists(x, y int) bool {
	if x < c.xMin || x > c.xMax || y < c.yMin || y > c.yMax {
		return true
	}

	return false
}

func (c *Cave) getPoint(x, y int) *Point {
	if c.exists(x, y) {
		return nil
	}

	return c.grid[y-c.yMin][x-c.xMin]
}

func (c *Cave) addPoint(point *Point) {
	x := point.x - c.xMin
	y := point.y - c.yMin
	c.grid[y][x] = point
}

func (c *Cave) addSandSource(point *Point) {
	c.addPoint(point)
	c.sandSource = point
}

func (c *Cave) addRockPath(rockPath []*Point) {
	for i := 0; i < len(rockPath)-1; i++ {
		firstPoint := rockPath[i]
		secondPoint := rockPath[i+1]

		for x := secondPoint.x; x <= firstPoint.x; x++ {
			c.addPoint(&Point{
				x:       x,
				y:       secondPoint.y,
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
	}
}

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

func (c *Cave) shouldStopAtPoint(x, y int) (bool, bool) {
	point := c.getPoint(x, y)

	if point == nil {
		return true, false
	}

	if point != nil && !point.shouldStop() {
		return false, true
	}

	return false, false
}

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
		currentPoint.element = Sand
		c.numFallenSand++
		return true
	}

	return false
}

func (c *Cave) fillWithSand() {
	do := c.fillWithSandFromPoint(c.sandSource)

	for do {
		do = c.fillWithSandFromPoint(c.sandSource)

		if VERBOSE {
			fmt.Println(c.numFallenSand)
			fmt.Println(c)
		}
	}
}

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

	// part 1
	sandSource := &Point{
		x:       500,
		y:       0,
		element: SandSource,
	}
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
}
