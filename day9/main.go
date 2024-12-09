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
	fmt.Println(calculateCheckSum(*fileSystem))
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
		digitAsInt := parseDigit(digit)
		if digitAsInt == -1 {
			return nil, 0, errors.New("non digit encountered inside of the row")
		}
		digits = append(digits, digitAsInt)
		sumOfDigits += digitAsInt
	}

	return digits, sumOfDigits, nil
}

func parseDigit(char rune) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	}

	return -1
}
