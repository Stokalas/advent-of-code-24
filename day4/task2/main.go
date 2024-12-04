package main

import (
	"fmt"
	"stokalas/advent-of-code/commonUtils"
	"stokalas/advent-of-code/day4/utils"
)

func main() {
	data, err := utils.ReadData("../data.txt")

	if err != nil {
		fmt.Println("Error while reading the file", err)
	}

	result := 0

	for index, row := range *data {
		if index == 0 || index == len(*data)-1 {
			continue
		}
		indexes := commonUtils.FindAllIndexesOfString(row, "A")

		topRow := index - 1

		bottomRow := index + 2

		result += processRow((*data)[topRow:bottomRow], indexes)
	}

	fmt.Println(result)
}

func processRow(rows []string, aIndexes []int) int {
	result := 0

	for _, index := range aIndexes {
		if index == 0 || index == len(rows[0])-1 {
			continue
		}
		if checkIfCorrect(rows, index) {
			result++
		}
	}

	return result
}

func checkIfCorrect(rows []string, index int) bool {
	lToR := false
	rToL := false
	topLeft := rows[0][index-1]
	topRight := rows[0][index+1]
	bottomLeft := rows[2][index-1]
	bottomRight := rows[2][index+1]
	if topLeft == 'M' && bottomRight == 'S' {
		lToR = true
	} else if topLeft == 'S' && bottomRight == 'M' {
		lToR = true
	}

	if !lToR {
		return false
	}

	if topRight == 'M' && bottomLeft == 'S' {
		rToL = true
	} else if topRight == 'S' && bottomLeft == 'M' {
		rToL = true
	}

	return rToL
}
