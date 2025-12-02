package days_2

import (
	"bytes"
	"math"
	"simonvreman/advent-of-code-2025/src/util"
	"strconv"
)

const FirstExpected int = 1227775554

type parsedRange = struct {
	start string
	end   string
}

func parseRanges(ranges [][]byte) []parsedRange {
	parsed := []parsedRange{}

	for _, single := range ranges {
		parts := util.Split(single, '-')
		start := string(parts[0])
		end := string(parts[1])
		parsed = append(parsed, parsedRange{start, end})
	}

	return parsed
}

// Considers only two repetitions exactly
func sumInvalidIdsInRanges(ranges []parsedRange) int {
	invalidSum := 0

	for _, single := range ranges {
		minLength := len(single.start)
		maxLength := len(single.end)

		for i := minLength; i <= maxLength; i++ {
			if i%2 != 0 {
				continue
			}

			halfLength := i / 2
			lowerBound := int(math.Pow10(halfLength - 1))
			upperBound := lowerBound * 10
			halfLengthMultiplier := upperBound

			if i == minLength {
				lowerStart, _ := strconv.Atoi(single.start[:halfLength])
				lowerEnd, _ := strconv.Atoi(single.start[halfLength:])

				lowerBound = lowerStart
				if lowerStart < lowerEnd {
					lowerBound++
				}
			}

			if i == maxLength {
				upperStart, _ := strconv.Atoi(single.end[:halfLength])
				upperEnd, _ := strconv.Atoi(single.end[halfLength:])

				upperBound = upperStart
				if upperStart <= upperEnd {
					upperBound++
				}
			}

			for id := lowerBound; id < upperBound; id++ {
				addToSum := id + id*halfLengthMultiplier
				invalidSum += addToSum
			}
		}
	}

	return invalidSum
}

func First(input []byte) int {
	ranges := parseRanges(util.Split(input, ','))
	return sumInvalidIdsInRanges(ranges)
}

const SecondExpected int = 4174379265

var divisorCount = map[int][]int{}

func findDivisors(length int) []int {
	count := divisorCount[length]
	if len(count) != 0 {
		return count
	}

	// Minimum divisor is two (otherwise, no repetion...)
	for i := 2; i <= length; i++ {
		if length%i == 0 {
			count = append(count, i)
		}
	}

	divisorCount[length] = count
	return count
}

// Guarantee all ranges have same starting and ending number length
func splitRangesByLength(ranges []parsedRange) []parsedRange {
	splittedRanges := []parsedRange{}

	for _, single := range ranges {
		minLength := len(single.start)
		maxLength := len(single.end)

		for i := minLength; i <= maxLength; i++ {
			start := single.start
			end := single.end

			if i != minLength {
				start = strconv.Itoa(int(math.Pow10(i - 1)))
			}
			if i != maxLength {
				end = strconv.Itoa(int(math.Pow10(i)) - 1)
			}

			splittedRanges = append(splittedRanges, parsedRange{start, end})
		}
	}

	return splittedRanges
}

func firstDistinctOrLast(ints []int, from int) int {
	for _, k := range ints {
		if k != from {
			return k
		}
	}
	return ints[len(ints)-1]
}

func sumInvalidIdsInRangesMultipleDivisors(ranges []parsedRange) int {
	invalidSum := 0

	for _, single := range ranges {
		length := len(single.start)
		divisors := findDivisors(length)

		for _, divisor := range divisors {
			chunkSize := length / divisor
			lower := []int{}
			upper := []int{}

			for i := range divisor {
				offset := i * chunkSize
				startChunk, _ := strconv.Atoi(single.start[offset : offset+chunkSize])
				endChunk, _ := strconv.Atoi(single.end[offset : offset+chunkSize])
				lower = append(lower, startChunk)
				upper = append(upper, endChunk)
			}

			repetitiveChunkModulo, _ := strconv.Atoi(string(bytes.Repeat([]byte{'1'}, chunkSize)))

			for chunk := lower[0]; chunk <= upper[0]; chunk++ {
				stringifiedChunk := strconv.Itoa(chunk)
				if (chunkSize > 1 && chunk%repetitiveChunkModulo == 0) ||
					(chunkSize%2 == 0 && stringifiedChunk[chunkSize/2:] == stringifiedChunk[:chunkSize/2]) ||
					(chunk == lower[0] && firstDistinctOrLast(lower, chunk) > chunk) ||
					(chunk == upper[0] && firstDistinctOrLast(upper, chunk) < chunk) {
					continue
				}

				for offset := range divisor {
					invalidSum += chunk * int(math.Pow10(offset*chunkSize))
				}
			}
		}
	}

	return invalidSum
}

func Second(input []byte) int {
	ranges := parseRanges(util.Split(input, ','))
	ranges = splitRangesByLength(ranges)
	return sumInvalidIdsInRangesMultipleDivisors(ranges)
}
