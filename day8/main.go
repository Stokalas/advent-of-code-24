package main

import (
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
		fmt.Println("Failed to read the file:", err)
		return
	}

	antennasMap := make(map[rune][]Coordinates, 32)

	for yIndex, row := range *data {
		for xIndex, char := range row {
			if char != '.' {
				antennasMap[char] = append(antennasMap[char], Coordinates{X: xIndex, Y: yIndex})
			}
		}
	}

	xLength := len((*data)[0])
	yLength := len(*data)

	antinodes := make([]Coordinates, 0, 200)
	for _, v := range antennasMap {
		processType(v, &antinodes, xLength, yLength)
	}

	fmt.Println(len(antinodes))
}

func processType(locations []Coordinates, antinodes *[]Coordinates, xLength, yLength int) {
	for i := 0; i < len(locations)-1; i++ {
		for j := i + 1; j < len(locations); j++ {
			x1 := locations[j].X - locations[i].X
			y1 := locations[j].Y - locations[i].Y

			newX := locations[j].X + x1
			newY := locations[j].Y + y1

			if newX >= 0 && newX < xLength && newY >= 0 && newY < yLength {
				pos := Coordinates{X: newX, Y: newY}
				if !antinodeAlreadyExists(pos, *antinodes) {
					*antinodes = append(*antinodes, pos)
				}
			}

			x2 := x1 * -1
			y2 := y1 * -1

			newX = locations[i].X + x2
			newY = locations[i].Y + y2

			if newX >= 0 && newX < xLength && newY >= 0 && newY < yLength {
				pos := Coordinates{X: newX, Y: newY}
				if !antinodeAlreadyExists(pos, *antinodes) {
					*antinodes = append(*antinodes, pos)
				}
			}
		}
	}
}

func antinodeAlreadyExists(pos Coordinates, antinodes []Coordinates) bool {
	for _, node := range antinodes {
		if node.X == pos.X && node.Y == pos.Y {
			return true
		}
	}
	return false
}
