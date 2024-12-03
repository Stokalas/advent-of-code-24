package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := readData("data.txt")

	if err != nil {
		fmt.Println("Error while reading file:", err)
	}

	result := 0
	mulIndexes := findAllIndexesOfString(data, "mul(")

	for _, index := range mulIndexes {
		value, err := processMul(data, index)

		if err != nil {
			continue
		}

		result += value
	}

	fmt.Println(result)
}

func findAllIndexesOfString(input, subStr string) []int {
	result := make([]int, 0, 100)
	subStrLen := len(subStr)

	index := strings.Index(input, subStr)
	for index != -1 {
		result = append(result, index)

		index = strings.Index(input[index+subStrLen:], subStr)
		if index != -1 {
			index += result[len(result)-1] + subStrLen
		}
	}
	return result
}

func processMul(input string, targetIndex int) (int, error) {
	targetRange := 12

	if len(input) < targetIndex+targetRange {
		targetRange = len(input) - targetIndex
	}

	if targetRange < 8 {
		return 0, errors.New("not enough space to have mul")
	}

	return getDigit(input[targetIndex : targetIndex+targetRange])
}

func getDigit(input string) (int, error) {
	separatorIndex := strings.Index(input, ",")

	if separatorIndex == -1 {
		return 0, errors.New("no separator found")
	}

	numberRunes := input[4:separatorIndex]

	firstNumber, err := strconv.Atoi(numberRunes)

	if err != nil {
		return 0, errors.New("failed to parse first int")
	}

	closingIndex := strings.Index(input, ")")

	if closingIndex == -1 {
		return 0, errors.New("no closing parentheses found")
	}

	numberRunes = input[separatorIndex+1 : closingIndex]

	secondNumber, err := strconv.Atoi(numberRunes)

	if err != nil {
		return 0, errors.New("failes to parse second int")
	}

	return firstNumber * secondNumber, nil
}

func readData(fileName string) (string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	result := ""

	for scanner.Scan() {
		lineContent := scanner.Text()

		result += lineContent
	}

	return result, nil
}
