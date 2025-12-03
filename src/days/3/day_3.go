package days_3

import (
	"strconv"
)

const FirstExpected int = 357

func parseBanks(input []byte) [][]int {
	banks := [][]int{{}}
	current := 0

	for _, item := range input {
		if item == '\n' {
			current++
			banks = append(banks, []int{})
			continue
		}

		battery, _ := strconv.Atoi(string(item))
		banks[current] = append(banks[current], battery)
	}

	return banks
}

func lowestMaxIndex(bank []int) int {
	lowestIndex := -1
	maxCharge := -1

	for i, charge := range bank {
		if charge > maxCharge {
			lowestIndex = i
			maxCharge = charge
		}
	}

	return lowestIndex
}

func First(input []byte) int {
	banks := parseBanks(input)
	sum := 0

	for _, bank := range banks {
		firstIndex := lowestMaxIndex(bank[:len(bank)-1])
		secondIndex := lowestMaxIndex(bank[firstIndex+1:]) + firstIndex + 1
		charge, _ := strconv.Atoi(strconv.Itoa(bank[firstIndex]) + strconv.Itoa(bank[secondIndex]))
		sum += charge
	}

	return sum
}

const SecondExpected int = 3121910778619

const batteryCount = 12

func Second(input []byte) int {
	banks := parseBanks(input)
	sum := 0

	for _, bank := range banks {
		lastIndex := -1
		charge := ""

		for battery := range batteryCount {
			minIndex := lastIndex + 1
			maxIndex := len(bank) - (batteryCount - (battery + 1))
			index := lowestMaxIndex(bank[minIndex:maxIndex]) + minIndex
			lastIndex = index
			charge += strconv.Itoa(bank[index])
		}

		parsedCharge, _ := strconv.Atoi(charge)
		sum += parsedCharge
	}

	return sum
}
