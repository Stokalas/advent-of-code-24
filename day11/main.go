package main

import (
	"errors"
	"fmt"
	"math"
	"stokalas/advent-of-code/commonUtils"
	"strconv"
	"strings"
)

func main() {
	data, err := commonUtils.ReadData("data.txt")

	if err != nil {
		fmt.Println("failed to read file:", err)
		return
	}

	pebbles, err := parsePebbles((*data)[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 25; i++ {
		pebbles = handleBlink(pebbles)
	}

	fmt.Println(len(*pebbles))
}

func handleBlink(pebbles *[]int) *[]int {
	result := make([]int, 0, len(*pebbles))
	for _, val := range *pebbles {
		if val == 0 {
			result = append(result, 1)
			continue
		}

		if val >= 10 {
			digitCount := findDigitCount(val)
			if digitCount%2 == 0 {
				firstHalf, secondHalf := splitInHalf(val, digitCount)
				result = append(result, firstHalf)
				result = append(result, secondHalf)
				continue
			}
		}

		result = append(result, val*2024)
	}

	return &result
}

func findDigitCount(num int) int {
	result := 1

	for {
		num = num / 10
		if num > 0 {
			result++
		} else {
			break
		}
	}

	return result
}

func splitInHalf(num int, digitCount int) (int, int) {
	tenPower := int(math.Pow10(digitCount / 2))
	return num / tenPower, num % tenPower
}

func parsePebbles(data string) (*[]int, error) {
	fields := strings.Fields(data)

	result := make([]int, 0, len(fields))

	for _, field := range fields {
		num, err := strconv.Atoi(field)

		if err != nil {
			return nil, errors.New("failed to parse numberå")
		}

		result = append(result, num)
	}

	return &result, nil
}
