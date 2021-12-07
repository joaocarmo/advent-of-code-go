package main

import (
	"fmt"
	"strconv"
	"strings"
)

type VentPoint struct {
	x     int
	y     int
	count int
}

func (p *VentPoint) toString() string {
	if p.count == 0 {
		return "."
	}

	return fmt.Sprintf("%d", p.count)
}

type Vent struct {
	start VentPoint
	end   VentPoint
}

func (v *Vent) new(start string, end string) *Vent {
	// parse the start and end points
	startPoints := strings.Split(start, ",")
	endPoints := strings.Split(end, ",")

	// convert the points to integers
	startX, _ := strconv.Atoi(startPoints[0])
	startY, _ := strconv.Atoi(startPoints[1])
	endX, _ := strconv.Atoi(endPoints[0])
	endY, _ := strconv.Atoi(endPoints[1])

	// add the start and end points
	v.start = VentPoint{startX, startY, 1}
	v.end = VentPoint{endX, endY, 1}

	return v
}

type Board struct {
	grid [][]VentPoint
}

func (b *Board) new(vents []Vent) *Board {
	// find the size of the board
	maxX, maxY := b.findBoardSizeFromVents(vents)

	// create the grid
	b.grid = make([][]VentPoint, maxY+1)

	for i := 0; i < maxY+1; i++ {
		b.grid[i] = make([]VentPoint, maxX+1)
	}

	// add the vent points to the grid
	for _, vent := range vents {
		b.addVentPoints(vent)
	}

	return b
}

func (b *Board) findBoardSizeFromVents(vents []Vent) (int, int) {
	// find the max x and y values
	var maxX, maxY int
	for _, vent := range vents {
		if vent.start.y > maxY {
			maxY = vent.start.y
		}
		if vent.end.y > maxY {
			maxY = vent.end.y
		}

		if vent.start.x > maxX {
			maxX = vent.start.x
		}
		if vent.end.x > maxX {
			maxX = vent.end.x
		}
	}

	// return the max and min x and y values
	return maxX, maxY
}

func (b *Board) addVentPoint(point VentPoint) {
	currentPoint := b.grid[point.y][point.x]

	// add the point to the grid
	b.grid[point.y][point.x] = VentPoint{point.x, point.y, currentPoint.count + point.count}
}

func (b *Board) addVentPoints(vent Vent) {
	if vent.start.x == vent.end.x {
		// add the vertical points
		start := minOf(vent.start.y, vent.end.y)
		end := maxOf(vent.start.y, vent.end.y)
		for i := start; i <= end; i++ {
			b.addVentPoint(VentPoint{vent.start.x, i, vent.start.count})
		}
	}

	if vent.start.y == vent.end.y {
		// add the horizontal points
		start := minOf(vent.start.x, vent.end.x)
		end := maxOf(vent.start.x, vent.end.x)
		for i := start; i <= end; i++ {
			b.addVentPoint(VentPoint{i, vent.start.y, vent.start.count})
		}
	}
}

func (b *Board) getOverlap(n int) int {
	var count int

	for _, row := range b.grid {
		for _, cell := range row {
			if cell.count >= n {
				count++
			}
		}
	}

	return count
}

func (b *Board) toString() string {
	var s string

	for _, row := range b.grid {
		for _, cell := range row {
			s += cell.toString()
		}
		s += "\n"
	}

	return s
}
