package test

import (
	"fmt"
	"os"
	"path/filepath"
	"simonvreman/advent-of-code-2025/src/days"
	"strconv"
	"testing"
)

func TestAllDays(t *testing.T) {
	for day, puzzles := range days.Solutions {
		for part, puzzle := range puzzles {
			runTestForPuzzle(t, day, part, puzzle)
		}
	}
}

func runTestForPuzzle(t *testing.T, day int, part int, puzzle days.Puzzle) {
	t.Run(fmt.Sprintf("day %v part %v", day, part+1), func(t *testing.T) {
		inputPath := filepath.Join("..", "days", strconv.Itoa(day), fmt.Sprintf("example_%v.txt", part+1))
		input, err := os.ReadFile(inputPath)

		if err != nil {
			t.Fatal("failed to read test input file", err.Error())
		}

		expected := puzzle.Expected
		actual := puzzle.Fn(input)

		if expected != actual {
			t.Errorf("expected %v, got %v instead", expected, actual)
		}
	})
}
