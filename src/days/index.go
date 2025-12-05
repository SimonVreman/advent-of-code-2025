package days

import (
	days_1 "simonvreman/advent-of-code-2025/src/days/1"
	days_2 "simonvreman/advent-of-code-2025/src/days/2"
	days_3 "simonvreman/advent-of-code-2025/src/days/3"
	days_4 "simonvreman/advent-of-code-2025/src/days/4"
	days_5 "simonvreman/advent-of-code-2025/src/days/5"
)

type PuzzleFunc func(input []byte) int
type Puzzle struct {
	Fn       PuzzleFunc
	Expected int
}

var Solutions = map[int][]Puzzle{
	1: {{days_1.First, days_1.FirstExpected}, {days_1.Second, days_1.SecondExpected}},
	2: {{days_2.First, days_2.FirstExpected}, {days_2.Second, days_2.SecondExpected}},
	3: {{days_3.First, days_3.FirstExpected}, {days_3.Second, days_3.SecondExpected}},
	4: {{days_4.First, days_4.FirstExpected}, {days_4.Second, days_4.SecondExpected}},
	5: {{days_5.First, days_5.FirstExpected}, {days_5.Second, days_5.SecondExpected}},
}