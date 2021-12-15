package main

type LanternFish struct {
	daysLeft int
}

func (lf *LanternFish) new(initialDays int) *LanternFish {
	if initialDays > -1 {
		lf.daysLeft = initialDays
	} else {
		lf.daysLeft = 8
	}

	return lf
}

func (lf *LanternFish) getDaysLeft() int {
	return lf.daysLeft
}

func (lf *LanternFish) resetDays() {
	lf.daysLeft = 6
}

func (lf *LanternFish) advanceDays(days int) {
	lf.daysLeft -= days
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
