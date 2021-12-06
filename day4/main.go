package main

import (
	"fmt"

	"github.com/stscoundrel/advent-of-code-2021/day4/boardgame"
	"github.com/stscoundrel/advent-of-code-2021/day4/reader"
)

func PlaWithTheSquid() int {
	numbers := reader.ReadBingoNumbersFromFile("./resources/bingos.txt")
	boards := reader.ReadBingoBoardsFromFile("./resources/bingos.txt")
	score := boardgame.Play(boards, numbers)

	return score
}

func main() {
	fmt.Println(PlaWithTheSquid())
}
