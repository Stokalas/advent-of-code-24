package main

import (
	"fmt"
	"stokalas/day3/utils"
)

func main() {
	data, err := utils.ReadData("../data.txt")

	if err != nil {
		fmt.Println("failed to read the file")
		return
	}

	result := 0

	mulIndexes := utils.FindAllIndexesOfString(data, "mul(")
	doIndexes := utils.FindAllIndexesOfString(data, "do()")
	dontIndexes := utils.FindAllIndexesOfString(data, "don't()")

	for _, mulIndex := range mulIndexes {
		closestDo := utils.FindClosestPrevIndex(mulIndex, &doIndexes)
		closestDont := utils.FindClosestPrevIndex(mulIndex, &dontIndexes)

		if closestDont > closestDo {
			continue
		}

		number, err := utils.ProcessMul(data, mulIndex)

		if err != nil {
			continue
		}

		result += number
	}

	fmt.Println(result)
}
