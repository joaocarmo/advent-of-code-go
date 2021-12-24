package main

import (
	"fmt"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

type SevenSegmentDisplay struct {
	top         string
	topLeft     string
	topRight    string
	middle      string
	bottom      string
	bottomLeft  string
	bottomRight string
}

func (s SevenSegmentDisplay) getSingalForNum(num int) string {
	var signal string

	switch num {
	case 0:
		signal = s.top + s.topLeft + s.topRight + s.bottom + s.bottomLeft + s.bottomRight
	case 1:
		signal = s.topRight + s.bottomRight
	case 2:
		signal = s.top + s.topRight + s.middle + s.bottomLeft + s.bottom
	case 3:
		signal = s.top + s.topRight + s.middle + s.bottom + s.bottomRight
	case 4:
		signal = s.topLeft + s.middle + s.topRight + s.bottomRight
	case 5:
		signal = s.top + s.topLeft + s.middle + s.bottom + s.bottomRight
	case 6:
		signal = s.top + s.topLeft + s.middle + s.bottom + s.bottomRight + s.bottomLeft
	case 7:
		signal = s.top + s.topRight + s.bottomRight
	case 8:
		signal = s.top + s.topLeft + s.topRight + s.middle + s.bottom + s.bottomLeft + s.bottomRight
	case 9:
		signal = s.top + s.topLeft + s.topRight + s.middle + s.bottom + s.bottomRight
	}

	return helpers.SortString(signal)
}

func (s SevenSegmentDisplay) getNumForSignal(signal string) int {
	var num int
	sortedSignal := helpers.SortString(signal)

	switch sortedSignal {
	case s.getSingalForNum(0):
		num = 0
	case s.getSingalForNum(1):
		num = 1
	case s.getSingalForNum(2):
		num = 2
	case s.getSingalForNum(3):
		num = 3
	case s.getSingalForNum(4):
		num = 4
	case s.getSingalForNum(5):
		num = 5
	case s.getSingalForNum(6):
		num = 6
	case s.getSingalForNum(7):
		num = 7
	case s.getSingalForNum(8):
		num = 8
	case s.getSingalForNum(9):
		num = 9
	}

	return num
}

func (s *SevenSegmentDisplay) toString() string {
	output := fmt.Sprintf(" %s\n", strings.Repeat(s.top, 4))
	output += fmt.Sprintf("%s    %s\n", s.topLeft, s.topRight)
	output += fmt.Sprintf("%s    %s\n", s.topLeft, s.topRight)
	output += fmt.Sprintf(" %s\n", strings.Repeat(s.middle, 4))
	output += fmt.Sprintf("%s    %s\n", s.bottomLeft, s.bottomRight)
	output += fmt.Sprintf("%s    %s\n", s.bottomLeft, s.bottomRight)
	output += fmt.Sprintf(" %s\n", strings.Repeat(s.bottom, 4))

	return output
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

	// Find the topLeft and middle from 0, middle should appear twice and
	// topLeft three times
	numTopLeft := getNumOfOccurrences(ssd.topLeft, signals[0])
	numMiddle := getNumOfOccurrences(ssd.middle, signals[0])

	if numMiddle > numTopLeft {
		tmpTopLeft := ssd.topLeft
		ssd.topLeft = ssd.middle
		ssd.middle = tmpTopLeft
	}

	// Find the topRight and bottomRight from 0, topRight should appear twice
	// and bottomRight three times
	numTopRight := getNumOfOccurrences(ssd.topRight, signals[0])
	numBottomRight := getNumOfOccurrences(ssd.bottomRight, signals[0])

	if numTopRight > numBottomRight {
		tmpTopRight := ssd.topRight
		ssd.topRight = ssd.bottomRight
		ssd.bottomRight = tmpTopRight
	}

	// Find the bottomLeft and bottom from 0, bottomLeft should appear twice and
	// bottom three times
	numBottomLeft := getNumOfOccurrences(ssd.bottomLeft, signals[0])
	numBottom := getNumOfOccurrences(ssd.bottom, signals[0])

	if numBottomLeft > numBottom {
		tmpBottomLeft := ssd.bottomLeft
		ssd.bottomLeft = ssd.bottom
		ssd.bottom = tmpBottomLeft
	}
}
