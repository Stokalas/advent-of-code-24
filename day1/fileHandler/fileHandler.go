package fileHandler

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func ReadData(fileName string) ([]int64, []int64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	firstColumn := make([]int64, 0, 1000)
	secondColumn := make([]int64, 0, 1000)

	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Fields(line)

		if len(numbers) < 2 {
			return nil, nil, errors.New("invalid line format")
		}

		num1, err1 := strconv.ParseInt(numbers[0], 10, 64)
		num2, err2 := strconv.ParseInt(numbers[1], 10, 64)

		if err1 != nil {
			return nil, nil, err1
		}

		if err2 != nil {
			return nil, nil, err2
		}

		firstColumn = append(firstColumn, num1)
		secondColumn = append(secondColumn, num2)
		index++
	}

	return firstColumn, secondColumn, nil
}
