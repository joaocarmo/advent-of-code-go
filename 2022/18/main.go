package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const VERBOSE = false
const COORDINATES_DELIMITER = ","
const FACES_PER_CUBE = 6

// Point represents a point in 3D space.
type Point struct {
	x, y, z int
}

// String returns a string representation of the point.
func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.x, p.y, p.z)
}

// Cube represents a cube.
type Cube struct {
	position *Point
}

// String returns a string representation of the cube.
func (c *Cube) String() string {
	return c.position.String()
}

// Grid represents a grid of cubes.
type Grid struct {
	cubes      []*Cube
	cubesMap   map[int]map[int]map[int]bool
	minX, maxX int
	minY, maxY int
	minZ, maxZ int
}

// getAdjacentFaces returns the number of adjacent faces.
func (g *Grid) getAdjacentFaces() int {
	adjacentFaces := 0
	for _, cube := range g.cubes {
		for _, otherCube := range g.cubes {
			cubesFaces := 0

			if cube == otherCube {
				continue
			}

			if cube.position.x == otherCube.position.x && cube.position.y == otherCube.position.y {
				if cube.position.z == otherCube.position.z+1 || cube.position.z == otherCube.position.z-1 {
					cubesFaces++
				}
			}

			if cube.position.x == otherCube.position.x && cube.position.z == otherCube.position.z {
				if cube.position.y == otherCube.position.y+1 || cube.position.y == otherCube.position.y-1 {
					cubesFaces++
				}
			}

			if cube.position.y == otherCube.position.y && cube.position.z == otherCube.position.z {
				if cube.position.x == otherCube.position.x+1 || cube.position.x == otherCube.position.x-1 {
					cubesFaces++
				}
			}

			adjacentFaces += cubesFaces
		}
	}

	return adjacentFaces
}

// getMinMax returns the min and max values for the grid.
func (g *Grid) getMinMax() {
	minX, maxX := math.MaxInt32, math.MinInt32
	minY, maxY := math.MaxInt32, math.MinInt32
	minZ, maxZ := math.MaxInt32, math.MinInt32

	for _, cube := range g.cubes {
		if cube.position.x < minX {
			minX = cube.position.x
		}
		if cube.position.x > maxX {
			maxX = cube.position.x
		}
		if cube.position.y < minY {
			minY = cube.position.y
		}
		if cube.position.y > maxY {
			maxY = cube.position.y
		}
		if cube.position.z < minZ {
			minZ = cube.position.z
		}
		if cube.position.z > maxZ {
			maxZ = cube.position.z
		}
	}

	g.minX = minX
	g.maxX = maxX
	g.minY = minY
	g.maxY = maxY
	g.minZ = minZ
	g.maxZ = maxZ
}

func (g *Grid) withinBounds(p *Point) bool {
	return p.x >= g.minX && p.x <= g.maxX && p.y >= g.minY && p.y <= g.maxY && p.z >= g.minZ && p.z <= g.maxZ
}

func (g *Grid) exists(p *Point) bool {
	for _, cube := range g.cubes {
		if cube.position.x == p.x && cube.position.y == p.y && cube.position.z == p.z {
			return true
		}
	}

	return false
}

func (g *Grid) addToInternalGrid(p *Point, v bool) {
	if _, ok := g.cubesMap[p.x]; !ok {
		g.cubesMap[p.x] = make(map[int]map[int]bool)
	}
	if _, ok := g.cubesMap[p.x][p.y]; !ok {
		g.cubesMap[p.x][p.y] = make(map[int]bool)
	}
	g.cubesMap[p.x][p.y][p.z] = v
}

func (g *Grid) inInternalGrid(p *Point) bool {
	if _, ok := g.cubesMap[p.x]; !ok {
		return false
	}
	if _, ok := g.cubesMap[p.x][p.y]; !ok {
		return false
	}
	if _, ok := g.cubesMap[p.x][p.y][p.z]; !ok {
		return false
	}

	return true
}

// getInternalFaces returns the number of internal faces.
func (g *Grid) getInternalCubes() []*Point {
	queue := []*Point{}
	start := &Point{g.minX, g.minY, g.minZ}
	queue = append(queue, start)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		g.addToInternalGrid(current, true)

		x, y, z := current.x, current.y, current.z
		directions := []*Point{
			&Point{x, y, z + 1},
			&Point{x, y, z - 1},
			&Point{x, y + 1, z},
			&Point{x, y - 1, z},
			&Point{x + 1, y, z},
			&Point{x - 1, y, z},
		}

		for _, direction := range directions {
			if alreadyInQueue(queue, direction) {
				continue
			}

			if g.inInternalGrid(direction) {
				continue
			}

			if !g.withinBounds(direction) {
				continue
			}

			if g.exists(direction) {
				continue
			}

			queue = append(queue, direction)
		}
	}

	internalCubes := []*Point{}
	for x := g.minX; x <= g.maxX; x++ {
		for y := g.minY; y <= g.maxY; y++ {
			for z := g.minZ; z <= g.maxZ; z++ {
				current := &Point{x, y, z}

				if g.inInternalGrid(current) {
					continue
				}

				if g.exists(current) {
					continue
				}

				internalCubes = append(internalCubes, current)
			}
		}
	}

	return internalCubes
}

// getExternalSurfaceArea returns the external surface area of the grid.
func (g *Grid) getInternalSurfaceArea() int {
	internalCubes := g.getInternalCubes()
	newGrid := newGrid(len(internalCubes))

	for i, cube := range internalCubes {
		newGrid.cubes[i] = &Cube{position: cube}
	}

	surfaceArea := newGrid.getSurfaceArea()

	return surfaceArea
}

// getSurfaceArea returns the surface area of the grid.
func (g *Grid) getSurfaceArea() int {
	totalFaces := len(g.cubes) * FACES_PER_CUBE
	adjacentFaces := g.getAdjacentFaces()
	return totalFaces - adjacentFaces
}

// String returns a string representation of the grid.
func (g *Grid) String() string {
	str := ""
	for _, cube := range g.cubes {
		str += cube.String() + "\n"
	}
	return str
}

// newGrid returns a new grid.
func newGrid(n int) *Grid {
	g := &Grid{
		cubes: make([]*Cube, n),
		cubesMap: make(map[int]map[int]map[int]bool),
	}
	return g
}

func alreadyInQueue(queue []*Point, direction *Point) bool {
	for _, p := range queue {
		if p.x == direction.x && p.y == direction.y && p.z == direction.z {
			return true
		}
	}
	return false
}

// getCoordinatesFromLine returns the coordinates from a line.
func getCoordinatesFromLine(line string) (int, int, int) {
	coordinates := strings.Split(line, COORDINATES_DELIMITER)
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	z, _ := strconv.Atoi(coordinates[2])
	return x, y, z
}

// getGridFromFile returns a grid from a file.
func getGridFromFile(txtlines []string) *Grid {
	grid := newGrid(len(txtlines))
	for i, line := range txtlines {
		x, y, z := getCoordinatesFromLine(line)
		grid.cubes[i] = &Cube{
			position: &Point{
				x: x,
				y: y,
				z: z,
			},
		}
	}
	return grid
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	grid := getGridFromFile(txtlines)

	// part 1
	surfaceArea := grid.getSurfaceArea()
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		surfaceArea,
	)

	// part 2
	grid.getMinMax()
	internalSurfaceArea := grid.getInternalSurfaceArea()
	fmt.Println(surfaceArea, "-", internalSurfaceArea, "=", surfaceArea - internalSurfaceArea)
}
