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

	for _, line := range *data {
		if isSafeLine(&line) {
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

		if difference == 0 {
			return false
		}

		if increasing && difference < 0 {
			return false
		} else if !increasing && difference > 0 {
			return false
		}

		if math.Abs(float64(difference)) > 3 {
			return false
		}
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
				return nil, errors.New("Failed to parse integer")
			}

			lineLevels = append(lineLevels, number)
		}

		result = append(result, lineLevels)
	}

	return &result, nil
}
