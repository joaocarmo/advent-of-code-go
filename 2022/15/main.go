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
const POINT_DELIMITER = ", "
const X_DELIMITER = "x="
const Y_DELIMITER = "y="
const CAVE_OFFSET = 12

type Element int

const (
	Air Element = iota
	Beacon
	Sensor
	Covered
)

func (e Element) String() string {
	return [...]string{"Beacon", "Covered", "Sensor"}[e]
}

type Point struct {
	x, y    int
	closest *Point
	element Element
}

func (p *Point) isSensor() bool {
	return p.element == Sensor
}

func (p *Point) isCovered() bool {
	return p.element == Covered
}

// findRange returns the range of the point by calculating the maximum
// distance between the point and its closest point.
func (p *Point) findMaxRange() int {
	if p.closest == nil {
		return -1
	}

	return helpers.AbsDiffInt(p.x, p.closest.x) + helpers.AbsDiffInt(p.y, p.closest.y)
}

func (p *Point) withinRange(sensor *Point) bool {
	maxRange := sensor.findMaxRange()
	distance := helpers.AbsDiffInt(p.x, sensor.x) + helpers.AbsDiffInt(p.y, sensor.y)

	return distance <= maxRange
}

func (p *Point) cover(sensor *Point) {
	if p.element == Beacon {
		return
	}

	if p.element == Sensor {
		return
	}

	if p.withinRange(sensor) {
		p.element = Covered
	}
}

func (p *Point) String() string {
	switch p.element {
	case Air:
		return "."
	case Beacon:
		return "B"
	case Covered:
		return "#"
	case Sensor:
		return "S"
	default:
		return "?"
	}
}

type Cave struct {
	grid                   map[int]map[int]*Point
	xMin, xMax, yMin, yMax int
	sensors				   []*Point
}

func (c *Cave) exists(x, y int) bool {
	if _, ok := c.grid[y-c.yMin]; !ok {
		return false
	}

	if _, ok := c.grid[y-c.yMin][x-c.xMin]; !ok {
		return false
	}

	return true
}

func (c *Cave) getPoint(x, y int) *Point {
	if c.exists(x, y) {
		return c.grid[y-c.yMin][x-c.xMin]
	}

	return &Point{
		x: x,
		y: y,
		element: Air,
	}
}

func (c *Cave) addPoint(point *Point) {
	x := point.x - c.xMin
	y := point.y - c.yMin

	if _, ok := c.grid[y]; !ok {
		c.grid[y] = make(map[int]*Point)
	}

	c.grid[y][x] = point

	if point.isSensor() {
		c.sensors = append(c.sensors, point)
	}
}

func (c *Cave) cover(sensor *Point) {
	sensorRange := sensor.findMaxRange()
	xMin := helpers.MaxOf(sensor.x-sensorRange, c.xMin)
	xMax := helpers.MinOf(sensor.x+sensorRange, c.xMax)
	yMin := helpers.MaxOf(sensor.y-sensorRange, c.yMin)
	yMax := helpers.MinOf(sensor.y+sensorRange, c.yMax)

	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			point := c.getPoint(x, y)

			point.cover(sensor)

			c.addPoint(point)
		}
	}
}

func (c *Cave) getCoverageAt(y int) int {
	coverage := 0

	for x := c.xMin; x <= c.xMax; x++ {
		point := c.getPoint(x, y)

		if point.isCovered() {
			coverage++
		}
	}

	return coverage
}

func (c *Cave) findCoverage() {
	for _, sensor := range c.sensors {
		c.cover(sensor)
	}
}

func (c *Cave) String() string {
	str := ""
	for y := c.yMin; y <= c.yMax; y++ {
		if VERBOSE {
			str += fmt.Sprintf("%d\t", c.yMin + y)
		}
		for x := c.xMin; x <= c.xMax; x++ {
			point := c.getPoint(x, y)
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

	c.grid = make(map[int]map[int]*Point)

	return &c
}

func getSensorAndBeaconFromLine(line string) (*Point, *Point) {
	sensor := &Point{}
	beacon := &Point{}

	re := regexp.MustCompile(`x=(-?\d+), y=(-?\d+)`)
	matches := re.FindAllString(line, -1)

	sensorXY := strings.Split(matches[0], POINT_DELIMITER)
	sensorX, _ := strconv.Atoi(strings.Replace(sensorXY[0], X_DELIMITER, "", 1))
	sensorY, _ := strconv.Atoi(strings.Replace(sensorXY[1], Y_DELIMITER, "", 1))

	beaconXY := strings.Split(matches[1], POINT_DELIMITER)
	beaconX, _ := strconv.Atoi(strings.Replace(beaconXY[0], X_DELIMITER, "", 1))
	beaconY, _ := strconv.Atoi(strings.Replace(beaconXY[1], Y_DELIMITER, "", 1))

	sensor.x = sensorX
	sensor.y = sensorY
	sensor.element = Sensor

	beacon.x = beaconX
	beacon.y = beaconY
	beacon.element = Beacon

	return sensor, beacon
}

func getMinMaxCoords(points []*Point) (int, int, int, int) {
	xMin := INFINITY
	xMax := 0
	yMin := INFINITY
	yMax := 0

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

	return xMin, xMax, yMin, yMax
}

func getPointsFromFile(lines []string) []*Point {
	points := []*Point{}

	for _, line := range lines {
		sensor, beacon := getSensorAndBeaconFromLine(line)
		sensor.closest = beacon
		points = append(points, sensor, beacon)
	}

	return points
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// print the text lines
	points := getPointsFromFile(txtlines)
	xMin, xMax, yMin, yMax := getMinMaxCoords(points)
	cave := newCave(
		xMin-CAVE_OFFSET,
		xMax+CAVE_OFFSET,
		yMin-CAVE_OFFSET,
		yMax+CAVE_OFFSET,
	)
	for _, point := range points {
		cave.addPoint(point)
	}
	cave.findCoverage()
	fmt.Println(cave)
	numBeaconCannotBePresent := cave.getCoverageAt(10)
	fmt.Println(numBeaconCannotBePresent)
}
