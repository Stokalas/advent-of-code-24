package main

import (
	"fmt"
	"stokalas/aoc-1/fileHandler"
)

func main() {
	firstCol, secondCol, err := fileHandler.ReadData("../data.txt")

	var result int64 = 0
	if err != nil {
		fmt.Println("Failed reading file: ", err.Error())
	}

	for i := 0; i < len(firstCol); i++ {
		temp := firstCol[i]
		appearances := findAmountOfRecordsInArray(temp, secondCol)
		result += temp * appearances
	}

	fmt.Println(result)
}

func findAmountOfRecordsInArray(item int64, array []int64) int64 {
	var result int64 = 0
	for i := 0; i < len(array); i++ {
		if item == array[i] {
			result++
		}
	}
	return result
}
