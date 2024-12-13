package main

import (
	"fmt"
	"slices"
	"stokalas/advent-of-code/commonUtils"
)

type Plot struct {
	X int
	Y int
}

type Field struct {
	PlantType rune
	Plots     []Plot
}

func main() {
	data, err := commonUtils.ReadData("data.txt")

	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}

	fields := make([]Field, 0, 100)

	for y, row := range *data {
		plots := []Plot{}
		currentType := rune(row[0])
		for x, plotType := range row {
			if plotType == currentType {
				plots = append(plots, Plot{X: x, Y: y})
			} else {
				assignPlotsToField(&fields, plots, currentType)
				plots = []Plot{{X: x, Y: y}}
				currentType = plotType
			}
		}
		assignPlotsToField(&fields, plots, currentType)
	}
	fields = *reconcileFields(&fields)

	result := 0
	result2nd := 0
	for _, field := range fields {
		area := calculateFieldArea(field)
		perimeter := calculateFieldPerimeter(field)
		sides := calculateFieldSides(field)

		result += area * perimeter
		result2nd += area * sides
	}

	fmt.Println(result)
	fmt.Println(result2nd)
}

func calculateFieldPerimeter(field Field) int {
	result := 0

	for _, plot := range field.Plots {
		//top
		if !containsPlot(field.Plots, plot.X, plot.Y-1) {
			result++
		}
		// bottom
		if !containsPlot(field.Plots, plot.X, plot.Y+1) {
			result++
		}
		// left
		if !containsPlot(field.Plots, plot.X-1, plot.Y) {
			result++
		}
		// right
		if !containsPlot(field.Plots, plot.X+1, plot.Y) {
			result++
		}
	}
	return result
}

func calculateFieldSides(field Field) int {
	result := 0
	right := []Plot{}
	left := []Plot{}
	top := []Plot{}
	bottom := []Plot{}

	for _, plot := range field.Plots {
		//top
		if !containsPlot(field.Plots, plot.X, plot.Y-1) {
			top = append(top, Plot{plot.X, plot.Y})
		}
		// bottom
		if !containsPlot(field.Plots, plot.X, plot.Y+1) {
			bottom = append(bottom, Plot{plot.X, plot.Y})
		}
		// left
		if !containsPlot(field.Plots, plot.X-1, plot.Y) {
			left = append(left, Plot{plot.X, plot.Y})
		}
		// right
		if !containsPlot(field.Plots, plot.X+1, plot.Y) {
			right = append(right, Plot{plot.X, plot.Y})
		}
	}

	result += handleVertical(left)
	result += handleVertical(right)
	result += handleHorizontal(top)
	result += handleHorizontal(bottom)

	return result
}

func handleVertical(plots []Plot) int {
	result := 0
	slices.SortFunc(plots, func(a, b Plot) int {
		if a.Y > b.Y {
			return 1
		}
		if a.Y == b.Y {
			return 0
		}

		return -1
	})

	for index, plot := range plots {
		found := false
		for j := index + 1; j < len(plots); j++ {
			if plot.X != plots[j].X {
				continue
			}

			if plots[j].Y == plot.Y+1 || plots[j].Y == plot.Y-1 {
				found = true
				break
			}
		}
		if !found {
			result++
		}
	}

	return result
}

func handleHorizontal(plots []Plot) int {
	result := 0
	slices.SortFunc(plots, func(a, b Plot) int {
		if a.X > b.X {
			return 1
		}
		if a.X == b.X {
			return 0
		}

		return -1
	})
	for index, plot := range plots {
		found := false
		for j := index + 1; j < len(plots); j++ {
			if plot.Y != plots[j].Y {
				continue
			}

			if plots[j].X == plot.X+1 || plots[j].X == plot.X-1 {
				found = true
				break
			}
		}
		if !found {
			result++
		}
	}

	return result
}

func containsPlot(plots []Plot, x, y int) bool {
	for _, plot := range plots {
		if plot.X == x && plot.Y == y {
			return true
		}
	}

	return false
}

func calculateFieldArea(field Field) int {
	return len(field.Plots)
}

func assignPlotsToField(fields *[]Field, plots []Plot, plotType rune) {
	for i, item := range *fields {
		if plotType != item.PlantType {
			continue
		}

		for _, currPlot := range item.Plots {
			for _, plot := range plots {
				distX := absInt(currPlot.X - plot.X)
				distY := absInt(currPlot.Y - plot.Y)
				if (distX == 1 && distY == 0) || (distX == 0 && distY == 1) {
					(*fields)[i].Plots = append(item.Plots, plots...)
					return
				}
			}

		}
	}

	*fields = append(*fields, Field{PlantType: plotType, Plots: plots})
}

func reconcileFields(fields *[]Field) *[]Field {
	result := make([]Field, 0, len(*fields)/2)
	for _, field := range *fields {
		assignPlotsToField(&result, field.Plots, field.PlantType)
	}

	return &result
}

func absInt(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
}
