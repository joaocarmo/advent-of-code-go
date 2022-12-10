package main

import (
	"fmt"
	"strconv"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// Matrix is a matrix of integers.
type Matrix [][]int

// get returns the value of a point in the matrix.
func (m Matrix) get(x, y int) int {
	return m[y][x]
}

// getColumnLength returns the length of the first column.
func (m Matrix) getColumnLength() int {
	return len(m)
}

// getRowLength returns the length of the first row.
func (m Matrix) getRowLength() int {
	if m.getColumnLength() < 1 {
		return 0
	}

	return len(m[0])
}

// isEdge returns true if the point is on the edge of the matrix.
func (m Matrix) isEdge(x, y int) bool {
	columnLength := m.getColumnLength()
	rowLength := m.getRowLength()

	if x == 0 || y == 0 || x == rowLength-1 || y == columnLength-1 {
		return true
	}

	return false
}

// isVisible returns true if the point is visible from the outside.
func (m Matrix) isVisible(x, y int) bool {
	if m.isEdge(x, y) {
		return true
	}

	height := m.get(x, y)
	isVisible := true

	// search the top
	for i := y - 1; i >= 0; i-- {
		if m.get(x, i) >= height {
			isVisible = false
			break
		}
	}

	if isVisible {
		return true
	}

	isVisible = true

	// search the bottom
	for i := y + 1; i < m.getColumnLength(); i++ {
		if m.get(x, i) >= height {
			isVisible = false
			break
		}
	}

	if isVisible {
		return true
	}

	isVisible = true

	// search the left
	for i := x - 1; i >= 0; i-- {
		if m.get(i, y) >= height {
			isVisible = false
			break
		}
	}

	if isVisible {
		return true
	}

	isVisible = true

	// search the right
	for i := x + 1; i < m.getRowLength(); i++ {
		if m.get(i, y) >= height {
			isVisible = false
			break
		}
	}

	return isVisible
}

func (m Matrix) String() string {
	s := ""

	for _, row := range m {
		for _, value := range row {
			s += fmt.Sprintf("%d ", value)
		}
		s += "\n"
	}

	return s
}

// calculateNumVisibleFromOutside returns the number of visible points from the outside.
func calculateNumVisibleFromOutside(matrix Matrix) int {
	numVisibleFromOutside := 0

	for y := 0; y < matrix.getColumnLength(); y++ {
		for x := 0; x < matrix.getRowLength(); x++ {
			if matrix.isVisible(x, y) {
				numVisibleFromOutside++
			}
		}
	}

	return numVisibleFromOutside
}

// getMatrixFromFile returns a matrix from a slice of strings.
func getMatrixFromFile(txtlines []string) Matrix {
	matrix := make(Matrix, len(txtlines))

	for i, line := range txtlines {
		matrix[i] = make([]int, len(line))

		for j, char := range line {
			matrix[i][j], _ = strconv.Atoi(string(char))
		}
	}

	return matrix
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	matrix := getMatrixFromFile(txtlines)

	// part 1
	numVisibleFromOutside := calculateNumVisibleFromOutside(matrix)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		numVisibleFromOutside,
	)
}
