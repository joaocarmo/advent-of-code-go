package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const VERBOSE = true
const INFINITY = int(^uint(0) >> 1)
const POINT_DELIMITER = ", "
const X_DELIMITER = "x="
const Y_DELIMITER = "y="
const COVERAGE_AT = 2000000
const MIN_COORDINATE = 0
const MAX_COORDINATE = 20
const TUNNING_FREQUENCY = 4000000

// Element represents the type of a point in the cave.
type Element int

const (
	Air Element = iota
	Beacon
	Sensor
	Covered
)

// String returns the string representation of an element.
func (e Element) String() string {
	return [...]string{"Beacon", "Covered", "Sensor"}[e]
}

// Point represents a point in the cave.
type Point struct {
	x, y     int
	closest  *Point
	element  Element
	maxRange int
}

// calculateTuningFrequency returns the tuning frequency of the point.
func (p *Point) calculateTuningFrequency() int {
	if p == nil {
		return 0
	}

	return p.x * TUNNING_FREQUENCY + p.y
}

// isAir returns true if the point is air.
func (p *Point) isAir() bool {
	return p.element == Air
}

// isBeacon returns true if the point is a beacon.
func (p *Point) isBeacon() bool {
	return p.element == Beacon
}

// isSensor returns true if the point is a sensor.
func (p *Point) isSensor() bool {
	return p.element == Sensor
}

// isCovered returns true if the point is covered by a sensor.
func (p *Point) isCovered() bool {
	return p.element == Covered
}

// isInvalid returns true if the point is valid.
func (p *Point) isInvalid() bool {
	return !p.isBeacon() && !p.isSensor()
}

// findRange returns the range of the point by calculating the maximum
// distance between the point and its closest point.
func (p *Point) findMaxRange() int {
	if p.maxRange > 0 {
		return p.maxRange
	}

	if p.closest == nil {
		return -1
	}

	p.maxRange = helpers.AbsDiffInt(p.x, p.closest.x) + helpers.AbsDiffInt(p.y, p.closest.y)

	return p.maxRange
}

// withinRange returns true if the point is within range of the sensor.
func (p *Point) withinRange(sensor *Point) bool {
	maxRange := sensor.findMaxRange()
	distance := helpers.AbsDiffInt(p.x, sensor.x) + helpers.AbsDiffInt(p.y, sensor.y)

	return distance <= maxRange
}

// cover covers the point with a sensor.
func (p *Point) cover(sensor *Point) {
	if p.isBeacon() {
		return
	}

	if p.isSensor() {
		return
	}

	if p.withinRange(sensor) {
		p.element = Covered
	}
}

// String returns the string representation of a point.
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

// Cave represents the cave.
type Cave struct {
	grid                   map[int]map[int]*Point
	xMin, xMax, yMin, yMax int
	sensors				   []*Point
	offset				   int
}

// exists returns true if the point exists in the cave.
func (c *Cave) exists(x, y int) bool {
	if _, ok := c.grid[y-c.yMin]; !ok {
		return false
	}

	if _, ok := c.grid[y-c.yMin][x-c.xMin]; !ok {
		return false
	}

	return true
}

// getPoint returns the point at the given coordinates.
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

// addPoint adds a point to the cave.
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

// getSensorCoverageAtY returns the number of points covered by a sensor at the given y coordinate.
func (c *Cave) getSensorCoverageAtY(y int, bounds ...int) int {
	coverage := 0
	xMin := c.xMin
	xMax := c.xMax

	if len(bounds) > 0 {
		xMin = bounds[0]
		xMax = bounds[1]
	}

	for x := xMin; x <= xMax; x++ {
		point := c.getPoint(x, y)

		for _, sensor := range c.sensors {
			if point.withinRange(sensor) {
				if point.isInvalid() {
					coverage++
				}
				break
			}
		}
	}

	return coverage
}

// getSensorCoverageAtX returns the number of points covered by a sensor at the given y coordinate.
func (c *Cave) getSensorCoverageAtX(x int, bounds ...int) int {
	coverage := 0
	yMin := c.yMin
	yMax := c.yMax

	if len(bounds) > 0 {
		yMin = bounds[0]
		yMax = bounds[1]
	}

	for y := yMin; y <= yMax; y++ {
		point := c.getPoint(x, y)

		for _, sensor := range c.sensors {
			if point.withinRange(sensor) {
				if point.isInvalid() {
					coverage++
				}
				break
			}
		}
	}

	return coverage
}

// findDistressSignal finds the distress signal within the given bounds.
func (c *Cave) findDistressSignal(min, max int) *Point {
	var point *Point
	coveragesX := make(map[int]int)
	coveragesY := make(map[int]int)

	for x := min; x <= max; x++ {
		coverageX := c.getSensorCoverageAtX(x, min, max)
		if coverageX < max - min {
			coveragesX[x] = coverageX
		}
	}

	for y := min; y <= max; y++ {
		coverageY := c.getSensorCoverageAtY(y, min, max)
		if coverageY < max - min {
			coveragesY[y] = coverageY
		}
	}

	xKeys := helpers.GetIntMapKeys(coveragesX)
	yKeys := helpers.GetIntMapKeys(coveragesY)
	sort.Ints(xKeys)
	sort.Ints(yKeys)

	for _, x := range xKeys {
		for _, y := range yKeys {
			point := c.getPoint(x, y)
			pointLeft := c.getPoint(x - 1, y)
			pointRight := c.getPoint(x + 1, y)

			if !point.isAir() || pointLeft.isAir() || pointRight.isAir() {
				continue
			}

			return point
		}
	}

	return point
}

// cover covers the cave with a sensor.
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

// getCoverageAt returns the number of points covered at the given y coordinate.
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

// findCoverage finds the coverage of the cave.
func (c *Cave) findCoverage() {
	for _, sensor := range c.sensors {
		c.cover(sensor)
	}
}

// String returns the string representation of the cave.
func (c *Cave) String() string {
	str := ""
	for y := c.yMin; y <= c.yMax; y++ {
		if VERBOSE {
			str += fmt.Sprintf("%d\t", (c.yMin + c.offset) + y)
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

// newCave returns a new cave.
func newCave(xMin, xMax, yMin, yMax, offset int) *Cave {
	c := Cave{
		xMin: xMin-offset,
		xMax: xMax+offset,
		yMin: yMin-offset,
		yMax: yMax+offset,
	}

	c.grid = make(map[int]map[int]*Point)
	c.offset = offset

	return &c
}

// getSensorAndBeaconFromLine returns the sensor and beacon from the given line.
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

// getMinMaxCoords returns the min and max x and y coordinates.
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

// getPointsFromFile returns the points from the given lines.
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

	// part 1
	coverageAt := COVERAGE_AT
	if len(args) > 1 {
		coverageAt, _ = strconv.Atoi(args[1])
	}
	caveOffset := coverageAt * 2
	points := getPointsFromFile(txtlines)
	xMin, xMax, yMin, yMax := getMinMaxCoords(points)
	cave := newCave(
		xMin,
		xMax,
		yMin,
		yMax,
		caveOffset,
	)
	for _, point := range points {
		cave.addPoint(point)
	}
	if VERBOSE {
		cave.findCoverage()
		fmt.Println(cave)
		numBeaconCannotBePresent := cave.getCoverageAt(coverageAt)
		fmt.Println(numBeaconCannotBePresent)
	} else {
		numBeaconNotPresent := cave.getSensorCoverageAtY(coverageAt)
		fmt.Printf(
			"[Part One] The answer is: %d\n",
			numBeaconNotPresent,
		)
	}

	// part 2
	distressSignal := cave.findDistressSignal(MIN_COORDINATE, MAX_COORDINATE)
	fmt.Println(distressSignal.x, distressSignal.y)
	fmt.Println(distressSignal.calculateTuningFrequency())
}
