package boardgame

func contains(slice []int, value int) bool {
	for _, sliceValue := range slice {
		if sliceValue == value {
			return true
		}
	}
	return false
}

func CheckBoard(board [][]int, numbers []int) bool {
	boardHasWon := false

	for i := 0; i < 5; i++ {
		hasRowBingo := true
		hasColumnBingo := true

		for j := 0; j < 5; j++ {
			if !contains(numbers, board[j][i]) {
				hasRowBingo = false
			}

			if !contains(numbers, board[i][j]) {
				hasColumnBingo = false
			}
		}

		if hasRowBingo || hasColumnBingo {
			boardHasWon = true
		}
	}

	return boardHasWon
}

func calculateScore(board [][]int, numbers []int) int {
	missedSum := 0
	lastNumber := numbers[len(numbers)-1]

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !contains(numbers, board[i][j]) {
				missedSum += board[i][j]
			}
		}
	}

	return missedSum * lastNumber
}

func playAllBoards(boards [][][]int, numbers []int) (scores []int) {
	playedNumbers := []int{}
	wonBoards := []int{}

	for _, number := range numbers {
		playedNumbers = append(playedNumbers, number)

		for boardIndex, board := range boards {
			if !contains(wonBoards, boardIndex) {
				if CheckBoard(board, playedNumbers) {
					scores = append(scores, calculateScore(board, playedNumbers))
					wonBoards = append(wonBoards, boardIndex)
				}
			}
		}
	}

	return
}

func Play(boards [][][]int, numbers []int) (score int) {
	playedNumbers := []int{}

	for _, number := range numbers {
		playedNumbers = append(playedNumbers, number)

		for _, board := range boards {
			if CheckBoard(board, playedNumbers) {
				score = calculateScore(board, playedNumbers)
				return
			}
		}
	}

	return
}

func PlayUntilTheEnd(boards [][][]int, numbers []int) (score int) {
	scores := playAllBoards(boards, numbers)
	score = scores[(len(scores) - 1)]

	return
}
