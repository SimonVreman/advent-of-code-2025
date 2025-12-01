package days

import (
	days_1 "simonvreman/advent-of-code-2025/src/days/1"
)

type PuzzleFunc func(input []byte) int
type Puzzle struct {
	Fn       PuzzleFunc
	Expected int
}

var Solutions = map[int][]Puzzle{
	1: {{days_1.First, days_1.FirstExpected}, {days_1.First, days_1.SecondExpected}},
}