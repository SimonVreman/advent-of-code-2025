package days_4

import "simonvreman/advent-of-code-2025/src/util"

const FirstExpected int = 13

type paperMatrix = [][]int
type position = struct {
	x int
	y int
}

const empty = -1
const moveThreshold = 4

func createPaperMatrix(lines [][]byte) paperMatrix {
	matrix := paperMatrix{}

	for _, line := range lines {
		vector := []int{}
		for _, position := range line {
			value := empty
			if position == '@' {
				value = 0
			}
			vector = append(vector, value)
		}
		matrix = append(matrix, vector)
	}

	return matrix
}

func findNeighbours(x int, y int, maxX int, maxY int) []position {
	neighbours := []position{}
	radius := 1

	for neighbourX := util.Max(x-radius, 0); neighbourX <= util.Min(x+radius, maxX-1); neighbourX++ {
		for neighbourY := util.Max(y-radius, 0); neighbourY <= util.Min(y+radius, maxY-1); neighbourY++ {
			if neighbourX == x && neighbourY == y {
				continue
			}
			neighbours = append(neighbours, position{neighbourX, neighbourY})
		}
	}

	return neighbours
}

func calculateNeighbours(matrix paperMatrix) paperMatrix {
	maxX := len(matrix)
	maxY := len(matrix[0])

	for x, vector := range matrix {
		for y, value := range vector {
			if value == empty {
				continue
			}

			for _, neighbour := range findNeighbours(x, y, maxX, maxY) {
				if matrix[neighbour.x][neighbour.y] != empty {
					matrix[neighbour.x][neighbour.y] = matrix[neighbour.x][neighbour.y] + 1
				}
			}
		}
	}

	return matrix
}

func countWhereFewerThanThreshold(matrix paperMatrix) int {
	count := 0

	for _, vector := range matrix {
		for _, value := range vector {
			if value != empty && value < moveThreshold {
				count++
			}
		}
	}

	return count
}

func First(input []byte) int {
	lines := util.Split(input, '\n')
	matrix := createPaperMatrix(lines)
	matrix = calculateNeighbours(matrix)
	return countWhereFewerThanThreshold(matrix)
}

const SecondExpected int = 43

func removeWhereFewerThanThreshold(matrix paperMatrix) (paperMatrix, int) {
	removed := 0

	for x, vector := range matrix {
		for y, value := range vector {
			if value == empty {
				continue
			}
			if value < moveThreshold {
				removed++
				matrix[x][y] = empty
			} else {
				matrix[x][y] = 0
			}
		}
	}

	return matrix, removed
}

func Second(input []byte) int {
	lines := util.Split(input, '\n')
	matrix := createPaperMatrix(lines)

	removed := 0
	count := 0

	for {
		matrix = calculateNeighbours(matrix)
		matrix, count = removeWhereFewerThanThreshold(matrix)
		removed += count

		if count == 0 {
			break
		}
	}

	return removed
}
