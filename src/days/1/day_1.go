package days_1

import (
	"simonvreman/advent-of-code-2025/src/util"
	"strconv"
)

const FirstExpected int = 3

const dialModulo = 100

func calculateDial(dial int, line []byte) int {
	lineString := string(line)
	value, _ := strconv.Atoi(lineString[1:])

	factor := -1
	if lineString[0] == 'R' {
		factor = 1
	}

	return (dial + factor*value) % dialModulo
}

func First(input []byte) int {
	lines := util.Split(input, '\n')
	dial := 50
	timesHitZero := 0

	for _, line := range lines {
		dial = calculateDial(dial, line)
		if dial == 0 {
			timesHitZero++
		}
	}

	return timesHitZero
}

const SecondExpected int = 6

func calculateDialWithZeroPasses(dial int, line []byte) (int, int) {
	lineString := string(line)
	value, _ := strconv.Atoi(lineString[1:])

	factor := -1
	if lineString[0] == 'R' {
		factor = 1
	}

	fullTurns := value / dialModulo
	partialTurn := value % dialModulo
	raw := dial + partialTurn*factor

	if (raw <= 0 && dial > 0) || raw >= dialModulo {
		fullTurns++
	}

	return (raw + dialModulo) % dialModulo, fullTurns
}

func Second(input []byte) int {
	lines := util.Split(input, '\n')
	dial := 50
	timesHitZero := 0

	for _, line := range lines {
		newDial, zeroPasses := calculateDialWithZeroPasses(dial, line)
		dial = newDial
		timesHitZero += zeroPasses
	}

	return timesHitZero
}
