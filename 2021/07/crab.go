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
	point1.New(crab.positionX, crab.positionY, crab.positionZ)

	point2 := &helpers.Point{}
	point2.New(positionX, positionY, positionZ)

	return int(helpers.Distance(point1, point2))
}
