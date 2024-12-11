package main

import (
	"errors"
	"fmt"
	"stokalas/advent-of-code/commonUtils"
)

type Peak struct {
	Row int
	Col int
}

func main() {
	data, err := commonUtils.ReadData("data.txt")

	if err != nil {
		fmt.Println("Error while reading the file:", err)
		return
	}

	topography, err := parseTopography(data)

	if err != nil {
		fmt.Println("Error while parsing data:", err)
	}

	score, rating := calculateScore(topography)
	fmt.Println(score)
	fmt.Println(rating)
}

func calculateScore(topography *[][]int8) (int, int) {
	result := 0
	rating := 0
	for rowIndex, row := range *topography {
		startIndexes := findIndexes(row, 0)
		for _, startIndex := range startIndexes {
			var reachablePeaks []Peak
			rating += navigateTrail(topography, rowIndex, startIndex, 0, &reachablePeaks)
			result += len(reachablePeaks)
		}
	}

	return result, rating
}

func navigateTrail(topography *[][]int8, rowIndex, startIndex int, currentVal int8, acc *[]Peak) int {
	if currentVal == 9 {
		addPeak(acc, Peak{Row: rowIndex, Col: startIndex})
		return 1
	}

	searchVal := currentVal + 1
	result := 0

	if rowIndex > 0 { // top
		if (*topography)[rowIndex-1][startIndex] == searchVal {
			result += navigateTrail(topography, rowIndex-1, startIndex, searchVal, acc)
		}
	}
	if rowIndex < len(*topography)-1 { // bottom
		if (*topography)[rowIndex+1][startIndex] == searchVal {
			result += navigateTrail(topography, rowIndex+1, startIndex, searchVal, acc)
		}
	}

	if startIndex > 0 { // left
		if (*topography)[rowIndex][startIndex-1] == searchVal {
			result += navigateTrail(topography, rowIndex, startIndex-1, searchVal, acc)
		}
	}
	if startIndex < len((*topography)[0])-1 { // right
		if (*topography)[rowIndex][startIndex+1] == searchVal {
			result += navigateTrail(topography, rowIndex, startIndex+1, searchVal, acc)
		}
	}

	return result
}

func addPeak(array *[]Peak, item Peak) {
	for _, existing := range *array {
		if item.Col == existing.Col && item.Row == existing.Row {
			return
		}
	}

	*array = append(*array, item)
}

func findIndexes(array []int8, searchItem int8) []int {
	var result []int

	for index, item := range array {
		if item == searchItem {
			result = append(result, index)
		}
	}

	return result
}

func parseTopography(data *[]string) (*[][]int8, error) {
	result := make([][]int8, 0, len(*data))

	for _, row := range *data {
		rowAsInts := make([]int8, 0, len(row))

		for _, val := range row {
			intVal := commonUtils.ParseDigitFromRune(val)

			if intVal == -1 {
				return nil, errors.New("failed to parse digit from rune")
			}

			rowAsInts = append(rowAsInts, int8(intVal))
		}
		result = append(result, rowAsInts)
	}

	return &result, nil
}
