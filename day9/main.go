package main

import (
	"errors"
	"fmt"
	"stokalas/advent-of-code/commonUtils"
)

func main() {
	data, err := commonUtils.ReadData("data.txt")

	if err != nil {
		fmt.Println("Error while reading the file:", err)
		return
	}

	dataRow := (*data)[0]

	digits, sumOfDigits, err := parseDigits(dataRow)

	if err != nil {
		fmt.Println("Failed to get digits:", err)
	}

	fileSystem := createFileSystem(sumOfDigits, digits)
	compactFileSystem(fileSystem)

	altFileSystem := createFileSystem(sumOfDigits, digits)
	compactFileSystemWithoutFrag(altFileSystem)

	fmt.Println("First result:", calculateCheckSum(*fileSystem))
	fmt.Println("Second result:", calculateCheckSum(*altFileSystem))
}

func calculateCheckSum(fileSystem []int) int {
	result := 0

	for index, value := range fileSystem {
		if value == -1 {
			continue
		}

		result += index * value
	}

	return result
}

func compactFileSystemWithoutFrag(fileSystem *[]int) {
	endIndex := len(*fileSystem) - 1

	for endIndex >= 0 {
		start, end := findNextMoveCandidate(fileSystem, endIndex)

		if start == -1 {
			break
		}

		length := end - start + 1

		spot := findSpotToMoveCandidate(fileSystem, length, start)

		if spot != -1 {
			for i := 0; i < length; i++ {
				temp := (*fileSystem)[spot+i]
				(*fileSystem)[spot+i] = (*fileSystem)[start+i]
				(*fileSystem)[start+i] = temp
			}
		}
		endIndex = start - 1
	}
}

func findSpotToMoveCandidate(fileSystem *[]int, length, limit int) int {
	index := 0
	currentLength := 0
	for index < len(*fileSystem) && index < limit {
		if (*fileSystem)[index] != -1 {
			index++
			currentLength = 0
			continue
		}

		index++
		currentLength++

		if currentLength == length {
			return index - currentLength
		}
	}

	return -1
}

func findNextMoveCandidate(fileSystem *[]int, index int) (int, int) {
	for (*fileSystem)[index] == -1 {
		index--
		if index < 0 {
			return -1, -1
		}
	}

	value := (*fileSystem)[index]
	endIndex := index

	for (*fileSystem)[index-1] == value {
		index--
		if index == 0 {
			return index, endIndex
		}
	}

	return index, endIndex
}

func compactFileSystem(fileSystem *[]int) {
	startIndex, endIndex := 0, len(*fileSystem)-1
	for {
		if (*fileSystem)[startIndex] != -1 {
			startIndex++
			continue
		}

		if (*fileSystem)[endIndex] == -1 {
			endIndex--
			continue
		}

		if startIndex >= endIndex {
			break
		}

		temp := (*fileSystem)[startIndex]
		(*fileSystem)[startIndex] = (*fileSystem)[endIndex]
		(*fileSystem)[endIndex] = temp
	}
}

func createFileSystem(sumOfDigits int, digits []int) *[]int {
	fileSystem := make([]int, 0, sumOfDigits)
	isSpace := false

	index := 0
	for _, number := range digits {
		valueToAdd := index
		if isSpace {
			valueToAdd = -1
		}

		for i := 0; i < number; i++ {
			fileSystem = append(fileSystem, valueToAdd)
		}

		if !isSpace {
			index++
		}
		isSpace = !isSpace
	}

	return &fileSystem
}

func parseDigits(dataRow string) ([]int, int, error) {
	digits := make([]int, 0, len(dataRow))
	sumOfDigits := 0
	for _, digit := range dataRow {
		digitAsInt := commonUtils.ParseDigitFromRune(digit)
		if digitAsInt == -1 {
			return nil, 0, errors.New("non digit encountered inside of the row")
		}
		digits = append(digits, digitAsInt)
		sumOfDigits += digitAsInt
	}

	return digits, sumOfDigits, nil
}
