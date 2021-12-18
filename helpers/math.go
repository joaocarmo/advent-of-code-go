package helpers

import "math"

// AbsDiffInt returns the absolute difference between two integers.
func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

// absInt returns the absolute value of the given integer.
func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

// MinOf returns the minimum of the given values.
func MinOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

// MaxOf returns the maximum of the given values.
func MaxOf(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

// Square returns the square of the given integer.
func Square(x int) int {
	return x * x
}

// Distance returns the distance between two points.
func Distance(p1 *Point, p2 *Point) float64 {
	squares := Square(p1.X-p2.X) + Square(p1.Y-p2.Y) + Square(p1.Z-p2.Z)
	return math.Sqrt(float64(squares))
}
