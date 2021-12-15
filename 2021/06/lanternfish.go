package main

const daysToNewFish = 7

type LanternFish struct {
	daysLeft    int
	totalResets int
}

func (lf *LanternFish) new(initialDays int) *LanternFish {
	if initialDays > -1 {
		lf.daysLeft = initialDays
	} else {
		lf.daysLeft = 8
	}

	lf.totalResets = 0

	return lf
}

func (lf *LanternFish) getDaysLeft() int {
	return lf.daysLeft
}

func (lf *LanternFish) getTotalResets() int {
	return lf.totalResets
}

func (lf *LanternFish) resetDays() {
	lf.daysLeft = daysToNewFish - 1
	lf.totalResets++
}

func (lf *LanternFish) advanceDays(days int) {
	lf.daysLeft -= days % daysToNewFish
	lf.totalResets += days / daysToNewFish
}

func (lf *LanternFish) nextDay() {
	lf.advanceDays(1)
}

func (lf *LanternFish) getNextDayFish() *LanternFish {
	// calculate the next day
	lf.nextDay()

	if lf.daysLeft < 0 {
		// reset the state of the initial fish
		lf.resetDays()

		// create the new fish
		newFish := &LanternFish{}
		newFish.new(-1)

		// return the new fish
		return newFish
	}

	return nil
}
