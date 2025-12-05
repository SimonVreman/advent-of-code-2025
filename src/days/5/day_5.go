package days_5

import (
	"bytes"
	"cmp"
	"simonvreman/advent-of-code-2025/src/util"
	"slices"
	"strconv"
)

const FirstExpected int = 3

type idRange = struct {
	start int
	end   int
}

func parseInput(input []byte) ([]idRange, []int) {
	sections := bytes.Split(input, []byte{'\n', '\n'})

	ranges := util.Map(util.Split(sections[0], '\n'), func(line []byte) idRange {
		parts := util.Split(line, '-')
		start, _ := strconv.Atoi(string(parts[0]))
		end, _ := strconv.Atoi(string(parts[1]))
		return idRange{start, end}
	})

	ids := util.Map(util.Split(sections[1], '\n'), func(line []byte) int {
		value, _ := strconv.Atoi(string(line))
		return value
	})

	slices.SortFunc(ranges, func(a idRange, b idRange) int { return cmp.Compare(a.start, b.start) })
	slices.Sort(ids)

	return ranges, ids
}

func First(input []byte) int {
	ranges, ids := parseInput(input)
	valid := 0

	for _, id := range ids {
		for _, fresh := range ranges {
			if id <= fresh.end {
				if id >= fresh.start {
					valid++
				}
				break
			}
		}
	}

	return valid
}

const SecondExpected int = 14

func mergeRanges(ranges []idRange) []idRange {
	merged := []idRange{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		previous := merged[len(merged)-1]
		current := ranges[i]

		if current.start <= (previous.end + 1) {
			if current.end > previous.end {
				merged[len(merged)-1] = idRange{previous.start, current.end}
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

func Second(input []byte) int {
	ranges, _ := parseInput(input)
	ranges = mergeRanges(ranges)
	count := 0

	for _, idRange := range ranges {
		// +1 because end is inclusive
		count += 1 + idRange.end - idRange.start
	}

	return count
}
