package main

import (
	"errors"
	"fmt"
	"stokalas/advent-of-code/commonUtils"
)

type Coordinates struct {
	X int
	Y int
}

func main() {
	data, err := commonUtils.ReadData("data.txt")

	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	warehouse, commands, err := parseData(data)

	if err != nil {
		fmt.Println("Failed to parse data:", err)
		return
	}

	robotAt, err := findRobot(&warehouse)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range commands {
		for _, move := range row {
			err := processMove(&warehouse, &robotAt, move)
			if err != nil {
				fmt.Println(err)
				return
			}
			// if i == 0 {
			// 	fmt.Println(string(move), "----------------")
			// 	for _, ff := range warehouse {
			// 		fmt.Println(ff)
			// 	}
			// }

		}
	}
	// fmt.Println(move, "----------------")
	// for _, ff := range warehouse {
	// 	fmt.Println(ff)
	// }
	fmt.Println(sumGPS(&warehouse))
}

func sumGPS(warehouse *[][]int) int {
	result := 0
	for y, row := range *warehouse {
		for x, item := range row {
			if item != 1 {
				continue
			}
			result += y*100 + x
		}
	}

	return result
}

func processMove(warehouse *[][]int, robotAt *Coordinates, move rune) error {
	switch move {
	case '^':
		processUp(warehouse, robotAt)
	case 'v':
		processDown(warehouse, robotAt)
	case '<':
		processLeft(warehouse, robotAt)
	case '>':
		processRight(warehouse, robotAt)
	default:
		return errors.New("unexpected move")
	}

	return nil
}

func processUp(warehouse *[][]int, robotAt *Coordinates) {
	if (*warehouse)[robotAt.Y-1][robotAt.X] == 0 {
		(*warehouse)[robotAt.Y-1][robotAt.X] = 2
		(*warehouse)[robotAt.Y][robotAt.X] = 0

		robotAt.Y = robotAt.Y - 1
		return
	} else if (*warehouse)[robotAt.Y-1][robotAt.X] == -1 {
		return
	}

	emptyAt := -1
	for i := robotAt.Y - 2; i >= 0; i-- {
		if (*warehouse)[i][robotAt.X] == -1 {
			break
		}

		if (*warehouse)[i][robotAt.X] == 0 {
			emptyAt = i
			break
		}
	}

	if emptyAt == -1 {
		return
	}

	(*warehouse)[robotAt.Y-1][robotAt.X] = 2
	(*warehouse)[emptyAt][robotAt.X] = 1
	(*warehouse)[robotAt.Y][robotAt.X] = 0
	robotAt.Y = robotAt.Y - 1
}

func processDown(warehouse *[][]int, robotAt *Coordinates) {
	if (*warehouse)[robotAt.Y+1][robotAt.X] == 0 {
		(*warehouse)[robotAt.Y+1][robotAt.X] = 2
		(*warehouse)[robotAt.Y][robotAt.X] = 0

		robotAt.Y = robotAt.Y + 1
		return
	} else if (*warehouse)[robotAt.Y+1][robotAt.X] == -1 {
		return
	}

	emptyAt := -1
	for i := robotAt.Y + 2; i < len(*warehouse); i++ {
		if (*warehouse)[i][robotAt.X] == -1 {
			break
		}

		if (*warehouse)[i][robotAt.X] == 0 {
			emptyAt = i
			break
		}
	}

	if emptyAt == -1 {
		return
	}

	(*warehouse)[robotAt.Y+1][robotAt.X] = 2
	(*warehouse)[emptyAt][robotAt.X] = 1
	(*warehouse)[robotAt.Y][robotAt.X] = 0
	robotAt.Y = robotAt.Y + 1
}

func processLeft(warehouse *[][]int, robotAt *Coordinates) {
	if (*warehouse)[robotAt.Y][robotAt.X-1] == 0 {
		(*warehouse)[robotAt.Y][robotAt.X-1] = 2
		(*warehouse)[robotAt.Y][robotAt.X] = 0

		robotAt.X = robotAt.X - 1
		return
	} else if (*warehouse)[robotAt.Y][robotAt.X-1] == -1 {
		return
	}

	emptyAt := -1
	for i := robotAt.X - 2; i >= 0; i-- {
		if (*warehouse)[robotAt.Y][i] == -1 {
			break
		}

		if (*warehouse)[robotAt.Y][i] == 0 {
			emptyAt = i
			break
		}
	}

	if emptyAt == -1 {
		return
	}

	(*warehouse)[robotAt.Y][robotAt.X-1] = 2
	(*warehouse)[robotAt.Y][emptyAt] = 1
	(*warehouse)[robotAt.Y][robotAt.X] = 0
	robotAt.X = robotAt.X - 1
}

func processRight(warehouse *[][]int, robotAt *Coordinates) {
	if (*warehouse)[robotAt.Y][robotAt.X+1] == 0 {
		(*warehouse)[robotAt.Y][robotAt.X+1] = 2
		(*warehouse)[robotAt.Y][robotAt.X] = 0

		robotAt.X = robotAt.X + 1
		return
	} else if (*warehouse)[robotAt.Y][robotAt.X+1] == -1 {
		return
	}

	emptyAt := -1
	for i := robotAt.X + 2; i < len((*warehouse)[0]); i++ {
		if (*warehouse)[robotAt.Y][i] == -1 {
			break
		}

		if (*warehouse)[robotAt.Y][i] == 0 {
			emptyAt = i
			break
		}
	}

	if emptyAt == -1 {
		return
	}

	(*warehouse)[robotAt.Y][robotAt.X+1] = 2
	(*warehouse)[robotAt.Y][emptyAt] = 1
	(*warehouse)[robotAt.Y][robotAt.X] = 0
	robotAt.X = robotAt.X + 1
}

func findRobot(warehouse *[][]int) (Coordinates, error) {
	for y, row := range *warehouse {
		for x, item := range row {
			if item == 2 {
				return Coordinates{X: x, Y: y}, nil
			}
		}
	}

	return Coordinates{}, errors.New("failed to find robot")
}

func parseData(data *[]string) ([][]int, []string, error) {
	warehouse := make([][]int, 0, 5)

	for index, row := range *data {
		if len(row) == 0 {
			return warehouse, (*data)[index+1:], nil
		}
		wRow, err := parseWarehouseRow(row)
		if err != nil {
			return nil, nil, err
		}

		warehouse = append(warehouse, wRow)
	}

	return nil, nil, errors.New("failed to find separator")
}

func parseWarehouseRow(row string) ([]int, error) {
	result := make([]int, 0, len(row))
	for _, char := range row {
		switch char {
		case '#':
			result = append(result, -1)
		case '.':
			result = append(result, 0)
		case 'O':
			result = append(result, 1)
		case '@':
			result = append(result, 2)
		default:
			return nil, errors.New("unexpected character")
		}
	}

	return result, nil
}
