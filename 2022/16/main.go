package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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
	// if the valve was visited
	visited  bool
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
		visited:  false,
	}
}

type Position struct {
	value   *Valve
	leadsTo []*Valve
}

func (p *Position) getBestNext() *Valve {
	best := p.leadsTo[0]
	for _, next := range p.leadsTo {
		if !next.visited {
			best = next
		}
	}
	return best
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

func (c *Cave) totalFlowRate() int {
	total := 0
	for _, position := range c.positions {
		if position.value.open {
			total += position.value.flowRate
		}
	}
	return total
}

func (c *Cave) getOpenValves() []string {
	openValves := make([]string, 0)
	for _, position := range c.positions {
		if position.value.open {
			openValves = append(openValves, position.value.label)
		}
	}
	return openValves
}

func (c *Cave) getPosition(v *Valve) *Position {
	for _, position := range c.positions {
		if position.value == v {
			return position
		}
	}

	return nil
}

func (c *Cave) openValve() {
	if !c.position.value.open && c.position.value.flowRate > 0 {
		c.moveInTime()
		c.position.value.open = true
		c.position.value.openedAt = c.elapsed

		if VERBOSE {
			fmt.Println("You open valve", c.position.value.label)
		}
	}
	c.position.value.visited = true
}

func (c *Cave) moveInTime() {
	if VERBOSE {
		currentPressure := c.totalFlowRate()
		openValves := strings.Join(c.getOpenValves(), ", ")
		fmt.Println("== Minute", c.elapsed, "==")
		if len(openValves) == 0 {
			fmt.Println("No valves are open.")
		} else {
			fmt.Println("Valves", openValves, "are open, releasing", currentPressure, "pressure.")
		}
	}
	c.elapsed++
}

func (c *Cave) move(v *Valve) {
	c.moveInTime()

	c.position = c.getPosition(v)
	if VERBOSE {
		fmt.Println("You move to valve", c.position.value.label)
	}
}

func (c *Cave) moveNext() {
	bestNextPosition := c.position.getBestNext()
	c.move(bestNextPosition)
}

func (c *Cave) findPath() {
	for c.elapsed < MINUTES_REMAINING {
		c.openValve()
		c.moveNext()
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
