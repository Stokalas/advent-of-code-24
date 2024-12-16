package main

import (
	"errors"
	"fmt"
	"stokalas/advent-of-code/commonUtils"
	"strconv"
	"strings"
)

type Robot struct {
	Position Coordinates
	Velocity Coordinates
}

type Coordinates struct {
	X int
	Y int
}

const WIDTH, LENGTH int = 101, 103

func main() {
	data, err := commonUtils.ReadData("data.txt")

	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	robots, err := parseRobots(data)

	if err != nil {
		fmt.Println("Failed to parse robots:", err)
		return
	}

	for i := 0; i < 100; i++ {
		for r := 0; r < len(*robots); r++ {
			moveRobot(&(*robots)[r])
		}
	}

	result := 1
	result *= countRobotsInQuadrant(0, robots)
	result *= countRobotsInQuadrant(1, robots)
	result *= countRobotsInQuadrant(2, robots)
	result *= countRobotsInQuadrant(3, robots)

	fmt.Println(result)
}

func countRobotsInQuadrant(quadrant int, robots *[]Robot) int {
	minX := 0
	maxX := (WIDTH / 2) - 1
	if quadrant == 1 || quadrant == 3 {
		minX = (WIDTH / 2) + 1
		maxX = WIDTH - 1
	}

	minY := 0
	maxY := (LENGTH / 2) - 1
	if quadrant == 2 || quadrant == 3 {
		minY = (LENGTH / 2) + 1
		maxY = LENGTH - 1
	}

	result := 0
	for _, r := range *robots {
		if r.Position.X >= minX && r.Position.X <= maxX {
			if r.Position.Y >= minY && r.Position.Y <= maxY {
				result++
			}
		}
	}

	return result
}

func moveRobot(robot *Robot) {
	robot.Position.X = calculateNewCoord(robot.Position.X, robot.Velocity.X, WIDTH)
	robot.Position.Y = calculateNewCoord(robot.Position.Y, robot.Velocity.Y, LENGTH)
}

func calculateNewCoord(current, step, size int) int {
	newCoord := current + step
	if newCoord < 0 {
		newCoord = size + newCoord
	}

	if newCoord >= size {
		newCoord = 0 + newCoord - size
	}

	return newCoord
}

func parseRobots(data *[]string) (*[]Robot, error) {
	result := make([]Robot, 0, len(*data))
	for _, row := range *data {
		fields := strings.Fields(row)
		if len(fields) != 2 {
			return nil, errors.New("unexpected number of fields while parsing data")
		}

		location, err := parseCoordinates(fields[0])
		if err != nil {
			return nil, err
		}

		velocity, err := parseCoordinates(fields[1])
		if err != nil {
			return nil, err
		}

		result = append(result, Robot{Position: location, Velocity: velocity})
	}

	return &result, nil
}

func parseCoordinates(input string) (Coordinates, error) {
	split := strings.Split(input, ",")

	if len(split) != 2 {
		return Coordinates{}, errors.New("unexpeted number of fields while parsing coordinates")
	}

	x, err1 := strconv.Atoi(split[0][2:])
	y, err2 := strconv.Atoi(split[1])

	if err1 != nil || err2 != nil {
		return Coordinates{}, errors.New("failed to parse number while parsing coordinates")
	}

	return Coordinates{X: x, Y: y}, nil
}
