package main

import (
	"fmt"
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
	cubes []*Cube
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
	grid := &Grid{
		cubes: make([]*Cube, len(txtlines)),
	}
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
}
