package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := readData("../data.txt")

	if err != nil {
		fmt.Println("Failed to read from file:", err.Error())
	}

	result := 0

	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Failed to parse system arg for Damper usage! Defaulting to simple mode")
		arg = 0
	}

	var safeCheckFunc func(*[]int) bool
	if arg == 1 {
		safeCheckFunc = isSafeLineWithDampener
	} else {
		safeCheckFunc = isSafeLine
	}

	for _, line := range *data {
		if safeCheckFunc(&line) {
			result++
		}
	}

	fmt.Println(result)
}

func isSafeLine(numbers *[]int) bool {
	if len(*numbers) < 2 {
		return true
	}

	increasing := (*numbers)[1]-(*numbers)[0] > 0
	for i := 0; i < len(*numbers)-1; i++ {
		difference := (*numbers)[i+1] - (*numbers)[i]
		inRange := checkIfInRange(difference)
		keepsTrend := checkIfKeepsTrend(difference, increasing)

		if !inRange || !keepsTrend {
			return false
		}
	}

	return true
}

func isSafeLineWithDampener(numbers *[]int) bool {
	if isSafeLine(numbers) {
		return true
	}

	for excluded := 0; excluded < len(*numbers); excluded++ {
		reducedList := make([]int, 0, len(*numbers)-1)
		for index, element := range *numbers {
			if index != excluded {
				reducedList = append(reducedList, element)
			}
		}
		if isSafeLine(&reducedList) {
			return true
		}
	}
	return false
}

func checkIfInRange(difference int) bool {
	return !(math.Abs(float64(difference)) > 3)
}

func checkIfKeepsTrend(difference int, increasing bool) bool {
	if !increasing && difference >= 0 {
		return false
	} else if increasing && difference <= 0 {
		return false
	}

	return true
}

func readData(fileName string) (*[][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	result := make([][]int, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		lineLevels := make([]int, 0, 6)

		numbers := strings.Fields(line)

		for _, element := range numbers {
			number, err := strconv.Atoi(element)

			if err != nil {
				return nil, errors.New("failed to parse integer")
			}

			lineLevels = append(lineLevels, number)
		}

		result = append(result, lineLevels)
	}

	return &result, nil
}
