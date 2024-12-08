package main

import (
	"errors"
	"fmt"
	"math"
	"stokalas/advent-of-code/commonUtils"
	"strconv"
	"strings"
)

type Equation struct {
	Result  int64
	Numbers []int64
}

func main() {
	data, err := commonUtils.ReadData("data.txt")

	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}

	var result, result2nd int64 = 0, 0

	for _, row := range *data {
		equation, err := parseEquation(row)

		if err != nil {
			fmt.Println("Failed to parse equation:", err)
			return
		}

		// 1st task
		operations := make([]int, len(equation.Numbers)-1)
		for {
			if evaluateEquation(equation, operations) {
				result += equation.Result
				break
			}

			operations, err = simulateBinaryAdd(operations)
			if err != nil {
				break
			}
		}

		// 2nd task
		operations = make([]int, len(equation.Numbers)-1)
		for {
			if evaluateEquation(equation, operations) {
				result2nd += equation.Result
				break
			}

			operations, err = simulateTrinaryAdd(operations)
			if err != nil {
				break
			}
		}
	}

	fmt.Println("First task result: ", result)
	fmt.Println("Second task result:", result2nd)
}

func evaluateEquation(equation Equation, operations []int) bool {
	acc := equation.Numbers[0]

	for i := 1; i < len(equation.Numbers); i++ {
		operation := operations[i-1]
		if operation == 0 {
			acc += equation.Numbers[i]
		} else if operation == 1 {
			acc *= equation.Numbers[i]
		} else if operation == 2 {
			digits := 1
			decimalPoints := equation.Numbers[i] / 10
			for decimalPoints > 0 {
				digits++
				decimalPoints = decimalPoints / 10
			}
			acc *= int64(math.Pow10(digits))
			acc += equation.Numbers[i]
		}
	}

	return acc == equation.Result
}

func simulateBinaryAdd(bits []int) ([]int, error) {
	indexer := len(bits) - 1

	for {
		if indexer < 0 {
			break
		}

		if bits[indexer] == 0 {
			bits[indexer]++
			return bits, nil
		} else {
			bits[indexer] = 0
			indexer--
		}
	}
	return nil, errors.New("out of range for index")
}

func simulateTrinaryAdd(trits []int) ([]int, error) {
	indexer := len(trits) - 1

	for {
		if indexer < 0 {
			break
		}

		if trits[indexer] < 2 {
			trits[indexer]++
			return trits, nil
		} else {
			trits[indexer] = 0
			indexer--
		}
	}
	return nil, errors.New("out of range for index")
}

func parseEquation(row string) (Equation, error) {
	splitted := strings.Split(row, ":")

	if len(splitted) != 2 {
		return Equation{}, errors.New("invalid data row during split")
	}

	testValue, err := strconv.ParseInt(splitted[0], 10, 64)

	if err != nil {
		return Equation{}, errors.New("failed to parse test value")
	}

	numberFields := strings.Fields(splitted[1])
	numbers := make([]int64, 0, 5)

	for _, numberStr := range numberFields {
		parsed, err := strconv.ParseInt(numberStr, 10, 64)

		if err != nil {
			return Equation{}, errors.New("failed to parse number: " + numberStr)
		}

		numbers = append(numbers, parsed)
	}

	return Equation{Result: testValue, Numbers: numbers}, nil
}
