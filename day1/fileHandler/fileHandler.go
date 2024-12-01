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

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	firstColumn := make([]int64, lineCount)
	secondColumn := make([]int64, lineCount)

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, nil, errors.New("failed to seek file start")
	}
	scanner = bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Fields(line)

		a, err1 := strconv.ParseInt(numbers[0], 10, 64)
		b, err2 := strconv.ParseInt(numbers[1], 10, 64)

		if err1 != nil {
			return nil, nil, err1
		}

		if err2 != nil {
			return nil, nil, err2
		}

		firstColumn[index] = a
		secondColumn[index] = b
		index++
	}

	return firstColumn, secondColumn, nil
}
