package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// Direction is a cardinal direction.
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Move struct {
	Direction Direction
	Steps     int
}

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}
	return ""
}

func (d Direction) Int() int {
	switch d {
	case Up:
		return 1
	case Down:
		return 2
	case Left:
		return 3
	case Right:
		return 4
	}
	return 0
}

type Point struct {
	X int
	Y int
}

type Rope struct {
	Start   Point
	Head    Point
	Tail    Point
	Visited map[Point]bool
}

func newRope(start Point) *Rope {
    rope := &Rope{start, start, start, make(map[Point]bool)}
	rope.Visited[start] = true
	return rope
}

func (r *Rope) moveTail(move Move) {
	distanceX, distanceY := distanceInt(&r.Head, &r.Tail)
	absDistanceX := helpers.AbsInt(distanceX)
	absDistanceY := helpers.AbsInt(distanceY)
	distance := distanceStraight(&r.Head, &r.Tail)

	if r.Head.X == r.Tail.X || r.Head.Y == r.Tail.Y {
		if absDistanceX > 1 {
			r.Tail.X += distanceX / absDistanceX
		}

		if absDistanceY > 1 {
			r.Tail.Y += distanceY / absDistanceY
		}
	} else if distance > math.Sqrt(2) {
		r.Tail.X += distanceX / absDistanceX
		r.Tail.Y += distanceY / absDistanceY
	}

	r.Visited[r.Tail] = true
}

func (r *Rope) moveHead(move Move) {
	for move.Steps > 0 {
		switch move.Direction {
		case Up:
			r.Head.Y += 1
		case Down:
			r.Head.Y -= 1
		case Left:
			r.Head.X -= 1
		case Right:
			r.Head.X += 1
		}

		r.moveTail(move)

		move.Steps -= 1
	}
}

func (r *Rope) move(moves []Move) {
	for _, move := range moves {
		r.moveHead(move)
	}
}

func distanceInt(p1, p2 *Point) (int, int) {
	xDistanceFromP1ToP2 := p1.X - p2.X
	yDistanceFromP1ToP2 := p1.Y - p2.Y

	return xDistanceFromP1ToP2, yDistanceFromP1ToP2
}

func distanceStraight(p1 *Point, p2 *Point) float64 {
	squares := helpers.Square(p1.X-p2.X) + helpers.Square(p1.Y-p2.Y)

	return math.Sqrt(float64(squares))
}

func strToDirection(s string) Direction {
	switch s {
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	case "R":
		return Right
	}
	return 0
}

func getMovesFromFile(lines []string) []Move {
	moves := []Move{}

	for _, line := range lines {
		// split the line by whitespace
		directionAndSteps := strings.Split(line, " ")
		direction := strToDirection(directionAndSteps[0])
		steps, _ := strconv.Atoi(directionAndSteps[1])
		moves = append(moves, Move{direction, steps})
	}

	return moves
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	moves := getMovesFromFile(txtlines)

	// part 1
	rope := newRope(Point{0, 0})
	rope.move(moves)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		len(rope.Visited),
	)
}
