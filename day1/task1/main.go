package main

import (
	"fmt"
	"stokalas/aoc-1/fileHandler"
)

func main() {
	firstCol, secondCol, err := fileHandler.ReadData("../data.txt")

	if err != nil {
		fmt.Println("Failed reading file: ", err.Error())
	}

	var result int64 = 0
	index1, index2 := 0, 0

	for i := 0; i < len(firstCol); i++ {
		index1 = findNextLowest(index1, firstCol)
		current1 := firstCol[index1]
		firstCol[index1] = -1

		index2 = findNextLowest(index2, secondCol)
		current2 := secondCol[index2]
		secondCol[index2] = -1

		result += absDiff(current1, current2)
	}

	fmt.Println(result)
}

func findNextLowest(index int, values []int64) int {
	for i := 0; i < len(values); i++ {
		if values[index] == -1 && values[i] > -1 {
			index = i
		}
		if values[i] != -1 && values[i] < values[index] {
			index = i
		}
	}
	return index
}

func absDiff(x, y int64) int64 {
	if x < y {
		return y - x
	}
	return x - y
}
