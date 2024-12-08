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

	xLength := len((*data)[0])
	yLength := len(*data)

	antennasMap := getAntennasMap(data)

	antinodes := make([]Coordinates, 0, 200)
	for _, v := range antennasMap {
		processType(v, &antinodes, xLength, yLength, false)
	}

	fmt.Println("First task result:", len(antinodes))

	antinodesWithHarmonics := make([]Coordinates, 0, 500)
	for _, v := range antennasMap {
		processType(v, &antinodesWithHarmonics, xLength, yLength, true)
	}

	fmt.Println("Second task result:", len(antinodesWithHarmonics))
}

func processType(locations []Coordinates, antinodes *[]Coordinates, xLength, yLength int, includeHarmonics bool) {
	for i := 0; i < len(locations)-1; i++ {
		for j := i + 1; j < len(locations); j++ {
			xDiff := locations[j].X - locations[i].X
			yDiff := locations[j].Y - locations[i].Y

			var funcToUse func(Coordinates, *[]Coordinates, int, int, int, int) bool

			if includeHarmonics {
				funcToUse = processAntinodeWithHarmonics
			} else {
				funcToUse = processAntinode
			}

			funcToUse(locations[j], antinodes, xDiff, yDiff, xLength, yLength)

			xDiff *= -1
			yDiff *= -1

			funcToUse(locations[i], antinodes, xDiff, yDiff, xLength, yLength)
		}
	}
}

func processAntinodeWithHarmonics(antenna Coordinates, antinodes *[]Coordinates, xDiff, yDiff, xLen, yLen int) bool {
	appendAntinode(antenna, antinodes)

	xDiffAdj := xDiff
	yDiffAdj := yDiff
	for processAntinode(antenna, antinodes, xDiffAdj, yDiffAdj, xLen, yLen) {
		xDiffAdj += xDiff
		yDiffAdj += yDiff
	}
	return true
}

func processAntinode(antenna Coordinates, antinodes *[]Coordinates, xDiff, yDiff, xLen, yLen int) bool {
	newX := antenna.X + xDiff
	newY := antenna.Y + yDiff

	if newX >= 0 && newX < xLen && newY >= 0 && newY < yLen {
		pos := Coordinates{X: newX, Y: newY}
		appendAntinode(pos, antinodes)
		return true
	}

	return false
}

func appendAntinode(pos Coordinates, antinodes *[]Coordinates) {
	for _, node := range *antinodes {
		if node.X == pos.X && node.Y == pos.Y {
			return
		}
	}
	*antinodes = append(*antinodes, pos)
}

func getAntennasMap(data *[]string) map[rune][]Coordinates {
	result := make(map[rune][]Coordinates, 32)

	for yIndex, row := range *data {
		for xIndex, char := range row {
			if char != '.' {
				result[char] = append(result[char], Coordinates{X: xIndex, Y: yIndex})
			}
		}
	}

	return result
}
