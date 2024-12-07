package main

import (
	"fmt"
	"stokalas/advent-of-code/commonUtils"
)

func main() {
	data, err := commonUtils.ReadData("../data.txt")

	if err != nil {
		fmt.Println("Error while reading the file", err)
	}

	searchRunes := []rune("XMAS")

	result := 0
	for index, row := range *data {
		indexes := commonUtils.FindAllIndexesOfString(row, "X")

		currentRowPos := 0

		topRow := index
		if index > 2 {
			topRow = index - 3
			currentRowPos += 3
		}

		bottomRow := index + 1
		if index+3 < len(*data) {
			bottomRow = index + 4
		}

		result += processRow((*data)[topRow:bottomRow], currentRowPos, indexes, searchRunes)
	}

	fmt.Println(result)
}

func processRow(rows []string, processedRow int, xIndexes []int, searchRunes []rune) int {
	result := 0
	for _, index := range xIndexes {
		result += checkHorizontal(rows[processedRow], index, searchRunes)
		result += checkVertical(rows, processedRow, index, searchRunes)
		result += checkDiagonal(rows, processedRow, index, searchRunes)
	}

	return result
}

func checkDiagonal(rows []string, processedRow int, index int, searchRunes []rune) int {
	result := 0
	if processedRow > 0 && index > 2 {
		result += 1
		for i, currentRune := range searchRunes {
			if rows[processedRow-i][index-i] != byte(currentRune) {
				result -= 1
				break
			}
		}
	}

	if processedRow > 0 && index < len(rows[0])-3 {
		result += 1
		for i, currentRune := range searchRunes {
			if rows[processedRow-i][index+i] != byte(currentRune) {
				result -= 1
				break
			}
		}
	}

	if len(rows) > processedRow+3 && index > 2 {
		result += 1
		for i, currentRune := range searchRunes {
			if rows[processedRow+i][index-i] != byte(currentRune) {
				result -= 1
				break
			}
		}
	}

	if len(rows) > processedRow+3 && index < len(rows[0])-3 {
		result += 1
		for i, currentRune := range searchRunes {
			if rows[processedRow+i][index+i] != byte(currentRune) {
				result -= 1
				break
			}
		}
	}

	return result
}

func checkVertical(rows []string, processedRow int, index int, searchRunes []rune) int {
	result := 0
	if processedRow > 0 {
		result += 1
		for i, currentRune := range searchRunes {
			if rows[processedRow-i][index] != byte(currentRune) {
				result -= 1
				break
			}
		}
	}

	if len(rows) > processedRow+3 {
		result += 1
		for i, currentRune := range searchRunes {
			if rows[processedRow+i][index] != byte(currentRune) {
				result -= 1
				break
			}
		}
	}

	return result
}

func checkHorizontal(row string, index int, searchRunes []rune) int {
	result := 0
	runeStr := []rune(row)

	if index > 2 {
		result += 1
		for i, currentRune := range searchRunes {
			if runeStr[index-i] != currentRune {
				result -= 1
				break
			}
		}
	}

	if index < len(runeStr)-3 {
		result += 1
		for i, currentRune := range searchRunes {
			if runeStr[index+i] != currentRune {
				result -= 1
				break
			}
		}
	}
	return result
}
