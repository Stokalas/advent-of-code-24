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
	data, err := commonUtils.ReadData("../data.txt")

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
	fmt.Println("initial", "-----------------------")
	printWarehouse(&warehouse)

	for index, row := range commands {
		for _, move := range row {
			err := processMove(&warehouse, &robotAt, move)
			if err != nil {
				fmt.Println(err)
				return
			}

			if index == 2 {
				fmt.Println(string(move), "-----------------------")
				printWarehouse(&warehouse)
			}
		}
	}
	fmt.Println(sumGPS(&warehouse))
}

func printWarehouse(t *[][]int) {
	for _, row := range *t {
		result := ""
		for _, char := range row {
			switch char {
			case -1:
				result += "#"
			case 0:
				result += "."
			case 1:
				result += "["
			case 2:
				result += "]"
			case 5:
				result += "@"
			}
		}
		fmt.Println(result)
	}
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
		(*warehouse)[robotAt.Y-1][robotAt.X] = 5
		(*warehouse)[robotAt.Y][robotAt.X] = 0

		robotAt.Y = robotAt.Y - 1
		return
	} else if (*warehouse)[robotAt.Y-1][robotAt.X] == -1 {
		return
	}

	emptyAt := -1
	tiles := [][]int{}

	if (*warehouse)[robotAt.Y-1][robotAt.X] == 1 {
		tiles = append(tiles, []int{robotAt.X, robotAt.X + 1})
	} else {
		tiles = append(tiles, []int{robotAt.X - 1, robotAt.X})
	}

	tilesIndex := 0
	for i := robotAt.Y - 2; i >= 0; i-- {
		newTiles := []int{}
		emptyAt = i
		for _, bottomTile := range tiles[tilesIndex] {

			if (*warehouse)[i][bottomTile] == -1 {
				return
			}

			if (*warehouse)[i][bottomTile] != 0 {
				if (*warehouse)[i][bottomTile] == 1 {
					newTiles = append(newTiles, []int{bottomTile, bottomTile + 1}...)
				} else {
					newTiles = append(newTiles, []int{bottomTile - 1, bottomTile}...)
				}
				emptyAt = -1
			}
		}
		if emptyAt != -1 {
			break
		}

		tiles = append(tiles, newTiles)
		tilesIndex++
	}
	tilesIndex = len(tiles) - 1
	for i := emptyAt; i <= robotAt.Y-2; i++ {
		for _, tile := range tiles[tilesIndex] {
			(*warehouse)[i][tile] = (*warehouse)[i+1][tile]
		}
		for _, tile := range tiles[tilesIndex] {
			(*warehouse)[i+1][tile] = 0
		}
		tilesIndex--
	}

	(*warehouse)[robotAt.Y-1][robotAt.X] = 5
	(*warehouse)[robotAt.Y][robotAt.X] = 0
	robotAt.Y = robotAt.Y - 1
}

func processDown(warehouse *[][]int, robotAt *Coordinates) {
	if (*warehouse)[robotAt.Y+1][robotAt.X] == 0 {
		(*warehouse)[robotAt.Y+1][robotAt.X] = 5
		(*warehouse)[robotAt.Y][robotAt.X] = 0

		robotAt.Y = robotAt.Y + 1
		return
	} else if (*warehouse)[robotAt.Y+1][robotAt.X] == -1 {
		return
	}

	emptyAt := -1
	tiles := [][]int{}

	if (*warehouse)[robotAt.Y+1][robotAt.X] == 1 {
		tiles = append(tiles, []int{robotAt.X, robotAt.X + 1})
	} else {
		tiles = append(tiles, []int{robotAt.X - 1, robotAt.X})
	}

	tilesIndex := 0
	for i := robotAt.Y + 2; i < len(*warehouse); i++ {
		newTiles := []int{}
		emptyAt = i
		for _, bottomTile := range tiles[tilesIndex] {

			if (*warehouse)[i][bottomTile] == -1 {
				return
			}

			if (*warehouse)[i][bottomTile] != 0 {
				if (*warehouse)[i][bottomTile] == 1 {
					newTiles = append(newTiles, []int{bottomTile, bottomTile + 1}...)
				} else {
					newTiles = append(newTiles, []int{bottomTile - 1, bottomTile}...)
				}
				emptyAt = -1
			}
		}
		if emptyAt != -1 {
			break
		}

		tiles = append(tiles, newTiles)
		tilesIndex++
	}

	if emptyAt == -1 {
		return
	}
	tilesIndex = len(tiles) - 1
	for i := emptyAt; i >= robotAt.Y+2; i-- {
		for _, tile := range tiles[tilesIndex] {
			(*warehouse)[i][tile] = (*warehouse)[i-1][tile]
		}
		for _, tile := range tiles[tilesIndex] {
			(*warehouse)[i-1][tile] = 0
		}
		tilesIndex--
	}

	(*warehouse)[robotAt.Y+1][robotAt.X] = 5
	(*warehouse)[robotAt.Y][robotAt.X] = 0
	robotAt.Y = robotAt.Y + 1
}

func processLeft(warehouse *[][]int, robotAt *Coordinates) {
	if (*warehouse)[robotAt.Y][robotAt.X-1] == 0 {
		(*warehouse)[robotAt.Y][robotAt.X-1] = 5
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

	boxStart := true
	for i := emptyAt; i < robotAt.X-1; i++ {
		if boxStart {
			(*warehouse)[robotAt.Y][i] = 1
			boxStart = false
		} else {
			(*warehouse)[robotAt.Y][i] = 2
			boxStart = true
		}
	}

	(*warehouse)[robotAt.Y][robotAt.X-1] = 5
	(*warehouse)[robotAt.Y][robotAt.X] = 0
	robotAt.X = robotAt.X - 1
}

func processRight(warehouse *[][]int, robotAt *Coordinates) {
	if (*warehouse)[robotAt.Y][robotAt.X+1] == 0 {
		(*warehouse)[robotAt.Y][robotAt.X+1] = 5
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

	boxStart := true
	for i := robotAt.X + 2; i <= emptyAt; i++ {
		if boxStart {
			(*warehouse)[robotAt.Y][i] = 1
			boxStart = false
		} else {
			(*warehouse)[robotAt.Y][i] = 2
			boxStart = true
		}
	}

	(*warehouse)[robotAt.Y][robotAt.X+1] = 5
	(*warehouse)[robotAt.Y][robotAt.X] = 0
	robotAt.X = robotAt.X + 1
}

func findRobot(warehouse *[][]int) (Coordinates, error) {
	for y, row := range *warehouse {
		for x, item := range row {
			if item == 5 {
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
	result := make([]int, 0, len(row)*2)
	for _, char := range row {
		switch char {
		case '#':
			result = append(result, -1)
			result = append(result, -1)
		case '.':
			result = append(result, 0)
			result = append(result, 0)
		case 'O':
			result = append(result, 1)
			result = append(result, 2)
		case '@':
			result = append(result, 5)
			result = append(result, 0)
		default:
			return nil, errors.New("unexpected character")
		}
	}

	return result, nil
}
