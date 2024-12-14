package main

import (
	"errors"
	"fmt"
	"stokalas/advent-of-code/commonUtils"
	"strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

type Game struct {
	A      Coordinates
	B      Coordinates
	Result Coordinates
}

func main() {
	data, err := commonUtils.ReadData("data.txt")

	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	games, err := parseData(data)
	if err != nil {
		fmt.Println("Failed to parse data:", err)
		return
	}

	result := 0
	for index, game := range *games {
		aPresses, bPresses, err := handleEquation(game)
		if err != nil {
			fmt.Println(index, err)
			continue
		}
		tokens := calculateTokens(aPresses, bPresses)
		fmt.Println(index, tokens)
		result += tokens
	}

	fmt.Println(result)
}

func calculateTokens(a, b int) int {
	return a*3 + b*1
}

func handleEquation(game Game) (int, int, error) {
	// nA.X + mB.X = Rez.X
	// nA.y + mB.Y = Rez.Y
	D := game.A.X*game.B.Y - game.B.X*game.A.Y
	if D == 0 {
		return 0, 0, errors.New("unsolvable equation")
	}
	Dx := game.Result.X*game.B.Y - game.B.X*game.Result.Y
	Dy := game.A.X*game.Result.Y - game.Result.X*game.A.Y

	if Dx%D != 0 || Dy%D != 0 {
		return 0, 0, errors.New("non-integer result")
	}
	return Dx / D, Dy / D, nil
}

func parseData(data *[]string) (*[]Game, error) {
	gameCount := (len(*data) + 1) / 4

	result := make([]Game, 0, gameCount)

	for i := 0; i < gameCount; i++ {
		aButt, err := parseButton((*data)[4*i])
		if err != nil {
			return nil, err
		}

		bButt, err := parseButton((*data)[4*i+1])
		if err != nil {
			return nil, err
		}

		prize, err := parsePrize((*data)[4*i+2])
		if err != nil {
			return nil, err
		}

		result = append(result, Game{A: aButt, B: bButt, Result: prize})
	}

	return &result, nil
}

func parseButton(data string) (Coordinates, error) {
	fields := strings.Fields(data)

	xCoord, err1 := strconv.Atoi(fields[2][2 : len(fields[2])-1])
	yCoord, err2 := strconv.Atoi(fields[3][2:])

	if err1 != nil || err2 != nil {
		return Coordinates{}, errors.New("failed to parse button integer")
	}

	return Coordinates{X: xCoord, Y: yCoord}, nil
}

func parsePrize(data string) (Coordinates, error) {
	fields := strings.Fields(data)

	xCoord, err1 := strconv.Atoi(fields[1][2 : len(fields[1])-1])
	yCoord, err2 := strconv.Atoi(fields[2][2:])

	if err1 != nil || err2 != nil {
		return Coordinates{}, errors.New("failed to parse prize integer")
	}

	return Coordinates{X: xCoord, Y: yCoord}, nil
}
