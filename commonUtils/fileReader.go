package commonUtils

import (
	"bufio"
	"os"
)

func ReadData(fileName string) (*[]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	rows := make([]string, 0, 150)

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	return &rows, nil
}
