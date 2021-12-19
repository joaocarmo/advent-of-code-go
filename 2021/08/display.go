package main

import (
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

type SevenSegmentDisplay struct {
	top         string // requires 1, 7
	topLeft     string
	topRight    string
	middle      string // requires 0, 8
	bottom      string
	bottomLeft  string
	bottomRight string
}

func (ssd *SevenSegmentDisplay) inferFromSingals(signals map[int][]string) {
	// Start by assuming the topRight and bottomRight from 1
	signalsRight := strings.Split(signals[1][0], "")
	ssd.topRight = signalsRight[0]
	ssd.bottomRight = signalsRight[1]

	// Find the top using 1 and 7
	ssd.top = helpers.StringDiff(signals[1][0], signals[7][0])[0]

	// Start by assuming the topLeft and middle from 1 and 4
	signalsTopLeftMiddle := helpers.StringDiff(signals[1][0], signals[4][0])
	ssd.topLeft = signalsTopLeftMiddle[0]
	ssd.middle = signalsTopLeftMiddle[1]

	// Start by assuming the bottomLeft and bottom from 4, 7, and 8
	signalsBottomLeftBottom := helpers.StringDiff(signals[4][0]+signals[7][0], signals[8][0])
	ssd.bottomLeft = signalsBottomLeftBottom[0]
	ssd.bottom = signalsBottomLeftBottom[1]

	// We now have assigned a unique signal to a segment of the display, but we
	// are still unsure about: topRight/bottomRight, topLeft/middle, and
	// bottomLeft/bottom
}
