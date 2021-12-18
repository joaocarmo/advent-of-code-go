package main

import (
	"github.com/joaocarmo/advent-of-code/helpers"
)

type Crab struct {
	positionX int
	positionY int
	positionZ int
}

func (c *Crab) new(positionX int, positionY int, positionZ int) {
	c.positionX = positionX
	c.positionY = positionY
	c.positionZ = positionZ
}

func (crab *Crab) getX() int {
	return crab.positionX
}

func (crab *Crab) getY() int {
	return crab.positionY
}

func (crab *Crab) getZ() int {
	return crab.positionZ
}

func (crab *Crab) getDistanceTo(positionX int, positionY int, positionZ int) int {
	point1 := &helpers.Point{}
	point1.New(crab.getX(), crab.getY(), crab.getZ())

	point2 := &helpers.Point{}
	point2.New(positionX, positionY, positionZ)

	return int(helpers.Distance(point1, point2))
}

func (crab *Crab) getWeightedDistanceTo(positionX int, positionY int, positionZ int) int {
	weightedDistance := 0
	linearDistance := crab.getDistanceTo(positionX, positionY, positionZ)

	// for each step in the distance, the weight of a change in position is
	// increased by 1.
	for i := 0; i <= linearDistance; i++ {
		weightedDistance += i
	}

	return weightedDistance
}
