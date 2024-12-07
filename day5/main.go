package main

import (
	"errors"
	"fmt"
	"stokalas/advent-of-code/commonUtils"
	"strconv"
	"strings"
)

func main() {
	data, err := commonUtils.ReadData("data.txt")

	if err != nil {
		fmt.Println("Failed to read the file", err)
	}

	sequences, sepIndex, err := parseSequences(*data)
	if err != nil {
		fmt.Println("Failed to parse sequences:", err)
	}

	result := 0
	for i := sepIndex + 1; i < len(*data); i++ {
		rowNumbers, err := parseRow((*data)[i])

		if err != nil {
			fmt.Println("Failed to parse number row", err)
		}

		if checkIfRowInOrder(rowNumbers, sequences) {
			result += getMiddleElement(rowNumbers)
		}
	}
	fmt.Println(result)
}

func getMiddleElement(rowNumbers []int) int {
	middleIndex := len(rowNumbers) / 2
	if len(rowNumbers)%2 != 0 {
		middleIndex++
	}
	return rowNumbers[middleIndex-1]
}

func checkIfRowInOrder(rowNumbers []int, sequences *map[int][]int) bool {
	for i := len(rowNumbers) - 1; i > 0; i-- {
		cannotBeAfter := (*sequences)[rowNumbers[i]]
		for j := i - 1; j >= 0; j-- {
			if arrayContains(rowNumbers[j], cannotBeAfter) {
				return false
			}
		}
	}
	return true
}

func arrayContains(element int, array []int) bool {
	for _, current := range array {
		if element == current {
			return true
		}
	}
	return false
}

func parseRow(input string) ([]int, error) {
	numbersAsString := strings.Split(input, ",")

	result := make([]int, 0, 25)

	for _, item := range numbersAsString {
		number, err := strconv.Atoi(item)

		if err != nil {
			return nil, err
		}

		result = append(result, number)
	}

	return result, nil
}

func parseSequences(input []string) (*map[int][]int, int, error) {
	result := make(map[int][]int)
	separatorRowIndex := 0

	for index, row := range input {
		if len(row) == 0 {
			separatorRowIndex = index
			break
		}

		numbersAsStrings := strings.Split(row, "|")

		if len(numbersAsStrings) != 2 {
			return nil, 0, errors.New("corrupted sequence condition encountered")
		}

		beforeNum, err1 := strconv.Atoi(numbersAsStrings[0])
		afterNum, err2 := strconv.Atoi(numbersAsStrings[1])

		if err1 != nil || err2 != nil {
			return nil, 0, errors.New("failed to parse number")
		}

		result[beforeNum] = append(result[beforeNum], afterNum)
	}

	return &result, separatorRowIndex, nil
}
