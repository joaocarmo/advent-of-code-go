package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const VERBOSE = true
const MINUTES_REMAINING = 30

type Valve struct {
	// valve label
	label    string
	// pressure per minute
	flowRate int
	// state of the valve
	open     bool
	// when the valve was opened
	openedAt int
}

func (v *Valve) calculatePressureReleased() int {
	return v.flowRate * (MINUTES_REMAINING - v.openedAt)
}

func (v *Valve) String() string {
	return fmt.Sprintf("%s: %2d", v.label, v.flowRate)
}

func newValve(label string, flowRate int) *Valve {
	return &Valve{
		label:    label,
		flowRate: flowRate,
		open:     false,
		openedAt: 0,
	}
}

type Position struct {
	value   *Valve
	leadsTo []*Valve
}

func (p *Position) String() string {
	return fmt.Sprintf("%s\t-> %s", p.value, p.leadsTo)
}

type Cave struct {
	positions []*Position
	position  *Position
	start     *Position
	elapsed   int
}

func (c *Cave) getPosition(v *Valve) *Position {
	for _, position := range c.positions {
		if position.value == v {
			return position
		}
	}

	return nil
}

func (c *Cave) moveInTime() {
	c.elapsed++
}

func (c *Cave) move(v *Valve) {
	for _, next := range c.position.leadsTo {
		if next == v {
			c.position = c.getPosition(next)
			c.openValve()
			break
		} else if VERBOSE {
			fmt.Println("-- Skipped", next.label)
		}
	}

	if VERBOSE {
		fmt.Println("Moved to", c.position.value.label, "at", c.elapsed, "min")
	}

	c.moveInTime()
}

func (c *Cave) moveNext() bool {
	for _, next := range c.position.leadsTo {
		if !next.open {
			c.move(next)
			return true
		} else if VERBOSE {
			fmt.Println("- Skipped", next.label)
		}
	}
	return false
}

func (c *Cave) openValve() {
	if !c.position.value.open {
		c.position.value.open = true
		c.position.value.openedAt = c.elapsed
		c.moveInTime()
	}
}

func (c *Cave) findPath() {
	for c.elapsed < MINUTES_REMAINING {
		if !c.position.value.open && c.position.value.flowRate > 0 {
			c.openValve()
		}

		if !c.moveNext() {
			return
		}
	}
}

func (c *Cave) String() string {
	return fmt.Sprintf("%s (%d min)", c.position.value.label, c.elapsed)
}

func newCave(positions []*Position) *Cave {
	start := positions[0]

	return &Cave{
		positions: positions,
		position:  start,
		start:     start,
		elapsed:   0,
	}
}

func getPressureReleased(positions []*Position) int {
	pressureReleased := 0
	for _, position := range positions {
		if position.value.open {
			pressureReleased += position.value.calculatePressureReleased()
		}
	}
	return pressureReleased
}

func getPositionsFromFile(txtlines []string) []*Position {
	valveMap := make(map[string]*Valve, len(txtlines))
	positions := make([]*Position, len(txtlines))

	for i, line := range txtlines {
		re1 := regexp.MustCompile(`(\d+)`)
		matches1 := re1.FindAllString(line, -1)
		flowRate, _ := strconv.Atoi(matches1[0])

		re2 := regexp.MustCompile(`([A-Z]{2})`)
		matches2 := re2.FindAllString(line, -1)
		valveLabel := matches2[0]
		positionNext := matches2[1:]

		if valve, ok := valveMap[valveLabel]; ok {
			valve.flowRate = flowRate
		} else {
			valveMap[valveLabel] = newValve(valveLabel, flowRate)
		}

		position := &Position{
			value:   valveMap[valveLabel],
			leadsTo: make([]*Valve, len(positionNext)),
		}

		for j, nextValveLabel := range positionNext {
			if valve, ok := valveMap[nextValveLabel]; ok {
				position.leadsTo[j] = valve
			} else {
				valveMap[nextValveLabel] = newValve(nextValveLabel, 0)
				position.leadsTo[j] = valveMap[nextValveLabel]
			}
		}

		positions[i] = position
	}

	return positions
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	positions := getPositionsFromFile(txtlines)
	for _, position := range positions {
		fmt.Println(position)
	}
	cave := newCave(positions)
	fmt.Println(cave)
	cave.findPath()
	pressureReleased := getPressureReleased(positions)
	fmt.Println(cave)
	fmt.Println(pressureReleased)
}
