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
}

func (c *Cave) isAbyss(x, y int) bool {
	if x < c.xMin || x > c.xMax || y < c.yMin || y > c.yMax {
		return true
	}

	return false
}

func (c *Cave) getPoint(x, y int) *Point {
	if c.isAbyss(x, y) {
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

func (c *Cave) fillWithSandFromPoint(xStart, yStart int) {
	if c.isAbyss(xStart, yStart) {
		return
	}

	x := xStart
	for y := yStart; y <= c.yMax; y++ {
		point := c.getPoint(x, y)

		// We reached rock or sand
		if point.shouldStop() {
			// Attempt to move one step down and to the left
			nextPoint := c.getPoint(x-1, y+1)

			if nextPoint == nil || nextPoint.shouldStop() {
				// Attempt to move one step down and to the right
				nextPoint = c.getPoint(x+1, y+1)

				if nextPoint == nil || nextPoint.shouldStop() {
					// Sand comes to a rest in the previous point
					point = c.getPoint(x, y-1)
					point.element = Sand
				}
			}
		}
	}
}

func (c *Cave) fillWithSand() {
	c.fillWithSandFromPoint(c.sandSource.x, c.sandSource.y)
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
	fmt.Println(cave)
}
