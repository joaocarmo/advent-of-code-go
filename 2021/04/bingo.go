package main

type BingoCard struct {
	sequence        []int
	winningSequence []int
	card            [][]int
	marked          [][]bool
}

func (b *BingoCard) new(sequence []int, card [][]int) *BingoCard {
	b.sequence = sequence
	b.card = card
	b.marked = make([][]bool, len(card))

	for i := 0; i < len(b.marked); i++ {
		b.marked[i] = make([]bool, len(card[i]))
	}

	return b
}

func (b *BingoCard) getCard() [][]int {
	return b.card
}

func (b *BingoCard) getWinningSequence() []int {
	return b.winningSequence
}

func (b *BingoCard) getNumberInCardIndex(number int) (int, int) {
	for i := 0; i < len(b.card); i++ {
		for j := 0; j < len(b.card[i]); j++ {
			if b.card[i][j] == number {
				return i, j
			}
		}
	}

	return -1, -1
}

func (b *BingoCard) markCardWithIndex(i int, j int) {
	if i > -1 && j > -1 {
		b.marked[i][j] = true
	}
}

func (b *BingoCard) isRowMarked(i int) bool {
	if i < 0 {
		return false
	}

	row := b.marked[i]

	for _, isMarked := range row {
		if !isMarked {
			return false
		}
	}

	return true
}

func (b *BingoCard) isColumnMarked(i int) bool {
	if i < 0 {
		return false
	}

	for _, row := range b.marked {
		if !row[i] {
			return false
		}
	}

	return true
}

func (b *BingoCard) getUnmarkedNumbers() []int {
	var unmarkedNumbers []int

	for i := 0; i < len(b.marked); i++ {
		for j := 0; j < len(b.marked[i]); j++ {
			if !b.marked[i][j] {
				unmarkedNumbers = append(unmarkedNumbers, b.card[i][j])
			}
		}
	}

	return unmarkedNumbers
}

func (b *BingoCard) getSumOfUnmarkedNumbers() int {
	var sum int

	for _, number := range b.getUnmarkedNumbers() {
		sum += number
	}

	return sum
}

func (b *BingoCard) getScore() int {
	sumOfUnmarkedNumbers := b.getSumOfUnmarkedNumbers()
	finalNumberInWinningSequence := b.winningSequence[len(b.winningSequence)-1]

	return sumOfUnmarkedNumbers * finalNumberInWinningSequence
}

func (b *BingoCard) findWinningSequence() {
	if len(b.winningSequence) > 0 {
		return
	}

	var tempWinningSequence []int

	// loop through the sequence
	for _, number := range b.sequence {
		// find the index of the number in the card
		i, j := b.getNumberInCardIndex(number)

		// mark the number in the card
		b.markCardWithIndex(i, j)

		// record the winning sequence
		tempWinningSequence = append(tempWinningSequence, number)

		// check if either a row or a column is marked
		if b.isRowMarked(i) || b.isColumnMarked(j) {
			b.winningSequence = tempWinningSequence
			break
		}
	}
}

func (b *BingoCard) isWinner() bool {
	if len(b.winningSequence) == 0 {
		b.findWinningSequence()
	}

	return len(b.winningSequence) > 0
}
