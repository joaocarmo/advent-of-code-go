package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const COORDINATES_DELIMITER = ","
const FACES_PER_CUBE = 6

type Point struct {
	x, y, z int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.x, p.y, p.z)
}

type Cube struct {
	position *Point
}

func (c *Cube) String() string {
	return c.position.String()
}

type Grid struct {
	cubes []*Cube
}

func (g *Grid) getAdjacentFaces() int {
	adjacentFaces := 0
	for _, cube := range g.cubes {
		for _, otherCube := range g.cubes {
			cubesFaces := 0

			if cube == otherCube {
				continue
			}

			if cube.position.x == otherCube.position.x && cube.position.y == otherCube.position.y {
				cubesFaces++
			}

			if cube.position.x == otherCube.position.x && cube.position.z == otherCube.position.z {
				cubesFaces++
			}

			if cube.position.y == otherCube.position.y && cube.position.z == otherCube.position.z {
				cubesFaces++
			}

			adjacentFaces += cubesFaces
		}
	}

	return adjacentFaces
}

func (g *Grid) getSurfaceArea() int {
	totalFaces := len(g.cubes) * FACES_PER_CUBE
	adjacentFaces := g.getAdjacentFaces()
	return totalFaces - adjacentFaces
}

func (g *Grid) String() string {
	str := ""
	for _, cube := range g.cubes {
		str += cube.String() + "\n"
	}
	return str
}

func getCoordinatesFromLine(line string) (int, int, int) {
	coordinates := strings.Split(line, COORDINATES_DELIMITER)
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	z, _ := strconv.Atoi(coordinates[2])
	return x, y, z
}

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
	surfaceArea := grid.getSurfaceArea()
	fmt.Println(grid)
	fmt.Println(surfaceArea)
}
