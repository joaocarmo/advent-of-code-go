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

// MinMax returns the minimum and maximum values of an array.
func MinMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]

    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }

    return min, max
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

// GCD finds the greatest common divisor via the Euclidean algorithm.
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

// LCM returns the lowest common multiple of the given numbers.
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// FindLCM returns the lowest common multiple of the slice of integers.
func FindLCM(integers []int) int {
	return LCM(integers[0], integers[1], integers[2:]...)
}

// SumInts returns the sum of the given integers.
func SumInts(ints ...int) int {
	sum := 0

	for _, i := range ints {
		sum += i
	}

	return sum
}
