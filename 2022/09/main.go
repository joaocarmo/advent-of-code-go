package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const NUM_OF_KNOTS_PART_1 = 2
const NUM_OF_KNOTS_PART_2 = 10

// Direction is a cardinal direction.
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// Move is a direction and the number of steps to take.
type Move struct {
	Direction Direction
	Steps     int
}

// String returns the string representation of a Direction.
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

// Point is a point in a 2D plane.
type Point struct {
	X int
	Y int
}

// Rope is a rope with a head and a tail.
type Rope struct {
	Start   Point
	Head    Point
	Body    []*Point
	Tail    Point
	Visited map[Point]bool
}

// newRope creates a new rope with a given start point and number of knots.
func newRope(start Point, numOfKnots int) *Rope {
	knots := make([]*Point, numOfKnots)
    for i := range knots {
        knots[i] = &Point{start.X, start.Y}
    }
    rope := &Rope{
			start,
			start,
			knots,
			start,
			make(map[Point]bool),
		}
	rope.Visited[start] = true
	return rope
}

// moveBody moves the body of the rope.
func (r *Rope) moveBody(move Move) {
	for i := 0; i < len(r.Body) - 1; i++ {
		moveTail(move, r.Body[i], r.Body[i+1])
	}
	r.Tail = *r.Body[len(r.Body)-1]
	r.Visited[r.Tail] = true
}

// moveHead moves the head of the rope.
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

		r.Body[0] = &Point{r.Head.X, r.Head.Y}
		r.moveBody(move)

		move.Steps -= 1
	}
}

// move moves the rope's head and body.
func (r *Rope) move(moves []Move) {
	for _, move := range moves {
		r.moveHead(move)
	}
}

// moveTail moves a piece of the tail of the rope.
func moveTail(move Move, head, tail *Point) {
	distanceX, distanceY := distanceInt(head, tail)
	absDistanceX := helpers.AbsInt(distanceX)
	absDistanceY := helpers.AbsInt(distanceY)
	distance := distanceStraight(head, tail)

	if head.X == tail.X || head.Y == tail.Y {
		if absDistanceX > 1 {
			tail.X += distanceX / absDistanceX
		}

		if absDistanceY > 1 {
			tail.Y += distanceY / absDistanceY
		}
	} else if distance > math.Sqrt(2) {
		tail.X += distanceX / absDistanceX
		tail.Y += distanceY / absDistanceY
	}
}

// distanceInt returns the distance between two points.
func distanceInt(p1, p2 *Point) (int, int) {
	xDistanceFromP1ToP2 := p1.X - p2.X
	yDistanceFromP1ToP2 := p1.Y - p2.Y

	return xDistanceFromP1ToP2, yDistanceFromP1ToP2
}

// distanceStraight returns the straight distance between two points.
func distanceStraight(p1 *Point, p2 *Point) float64 {
	squares := helpers.Square(p1.X-p2.X) + helpers.Square(p1.Y-p2.Y)

	return math.Sqrt(float64(squares))
}

// strToDirection converts a string to a Direction.
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

// getMovesFromFile returns a list of moves from a list of strings.
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
	ropeA := newRope(Point{0, 0}, NUM_OF_KNOTS_PART_1)
	ropeA.move(moves)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		len(ropeA.Visited),
	)

	// part 2
	ropeB := newRope(Point{0, 0}, NUM_OF_KNOTS_PART_2)
	ropeB.move(moves)
	fmt.Printf(
		"[Part Two] The answer is: %d\n",
		len(ropeB.Visited),
	)
}
