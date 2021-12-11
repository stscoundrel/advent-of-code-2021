package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day8/reader"
	"github.com/stscoundrel/advent-of-code-2021/day8/signals"
)

func CountUniquesInOutput(fileName string) int {
	lines := reader.ReadSignalInputFromFile(fileName)
	signalGroup := signals.LinesToSignals(lines)
	count := signals.CountUniquesInOutputs(signalGroup)

	return count
}
