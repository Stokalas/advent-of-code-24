package main

import (
	"fmt"
	"stokalas/day3/utils"
)

func main() {
	data, err := utils.ReadData("../data.txt")

	if err != nil {
		fmt.Println("Error while reading file:", err)
	}

	result := 0
	mulIndexes := utils.FindAllIndexesOfString(data, "mul(")

	for _, index := range mulIndexes {
		value, err := utils.ProcessMul(data, index)

		if err != nil {
			continue
		}

		result += value
	}

	fmt.Println(result)
}
