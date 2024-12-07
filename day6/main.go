package main

import (
	"errors"
	"fmt"
	"stokalas/advent-of-code/commonUtils"
	"strings"
)

func main() {
	data, err := commonUtils.ReadData("data.txt")

	if err != nil {
		fmt.Println("Failed to read the file", err)
		return
	}

	// 0 - up; 1 - right; 2 - down; 3 - left
	direction := 0
	x, y, err := findStartingPos(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	secondTaskResult := bruteForceObstacles(data)
	result := 0
	for {
		if !stillInside(data, x, y) {
			break
		}

		unique, err := processStep(data, &direction, &x, &y)

		if err != nil {
			fmt.Println("error while processing step", err)
			return
		}

		result += unique
	}

	fmt.Println(result)
	fmt.Println(secondTaskResult)
}

func bruteForceObstacles(data *[]string) int {
	result := 0
	x, y, _ := findStartingPos(data)

	for i := 0; i < len(*data); i++ {
		for j := 0; j < len((*data)[0]); j++ {
			var dataCopy = make([]string, len(*data))
			copy(dataCopy[:], (*data)[:])
			currRow := dataCopy[i]

			if currRow[j] == '#' || currRow[j] == '^' {
				continue
			}
			dataCopy[i] = currRow[:j] + "#" + currRow[j+1:]
			xCurr := x
			yCurr := y
			direction := 0
			noNewPlaces := 0

			for {
				if !stillInside(&dataCopy, xCurr, yCurr) {
					break
				}
				if noNewPlaces >= 20000 {
					result++
					break
				}

				unique, err := processStep(&dataCopy, &direction, &xCurr, &yCurr)
				if err != nil {
					fmt.Println("error while processing step", err)
					return -1
				}

				if unique == 0 {
					noNewPlaces++
				} else {
					noNewPlaces = 0
				}
			}
		}
	}
	return result
}

func stillInside(data *[]string, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if y >= len(*data) || x >= len((*data)[0]) {
		return false
	}
	return true
}

func processStep(data *[]string, direction *int, x *int, y *int) (int, error) {
	currRow := (*data)[*y]
	result := 0
	if currRow[*x] != 'X' {
		result++
		(*data)[*y] = currRow[:*x] + "X" + currRow[*x+1:]
	}

	shouldTurn := false

	switch *direction {
	case 0:
		shouldTurn = processUpStep(data, x, y)
	case 1:
		shouldTurn = processRightStep(data, x, y)
	case 2:
		shouldTurn = processDownStep(data, x, y)
	case 3:
		shouldTurn = processLeftStep(data, x, y)
	default:
		return 0, errors.New("invalid direction value")
	}

	if shouldTurn {
		if *direction == 3 {
			*direction = 0
		} else {
			*direction += 1
		}
	}

	return result, nil
}

func processUpStep(data *[]string, x *int, y *int) bool {
	if *y > 0 {
		if (*data)[*y-1][*x] == '#' {
			return true
		}
	}

	*y -= 1
	return false
}

func processRightStep(data *[]string, x *int, y *int) bool {
	if *x < len((*data)[0])-1 {
		if (*data)[*y][*x+1] == '#' {
			return true
		}
	}

	*x += 1
	return false
}

func processDownStep(data *[]string, x *int, y *int) bool {
	if *y < len(*data)-1 {
		if (*data)[*y+1][*x] == '#' {
			return true
		}
	}

	*y += 1
	return false
}

func processLeftStep(data *[]string, x *int, y *int) bool {
	if *x > 0 {
		if (*data)[*y][*x-1] == '#' {
			return true
		}
	}

	*x -= 1
	return false
}

func findStartingPos(data *[]string) (int, int, error) {
	for yIndex, row := range *data {
		xIndex := strings.Index(row, "^")
		if xIndex != -1 {
			return xIndex, yIndex, nil
		}
	}

	return -1, -1, errors.New("were not able to find guard")
}
