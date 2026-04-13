package pattern

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Pattern struct {
	x     int
	y     int
	cells [][]int
}

func (pattern *Pattern) GetNextGeneration() Pattern {
	nextGen := Pattern{
		pattern.x,
		pattern.y,
		pattern.cells,
	}
	return nextGen
}

func (pattern *Pattern) ToString() string {
	var sb strings.Builder
	for _, row := range pattern.cells {
		for _, cell := range row {
			if cell == 1 {
				sb.WriteRune('#')
			}
			if cell == 0 {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func ParsePatternFromRLEPatternString(x int, y int, pattern string) (Pattern, error) {
	cells := [][]int{}
	repeat := 0
	row := []int{}
	for _, runeValue := range pattern {
		if unicode.IsDigit(runeValue) {
			number, err := strconv.Atoi(string(runeValue))
			if err != nil {
				return Pattern{}, fmt.Errorf("Error parsing pattern %v", err)
			}
			repeat = 10*repeat + number
			continue
		}
		if repeat == 0 {
			repeat = 1
		}
		switch runeValue {
		case 'o':
			for range repeat {
				row = append(row, 1)
			}
		case 'b':
			for range repeat {
				row = append(row, 0)
			}
		case '$', '!':
			if len(row) != x {
				row = addPaddingToRow(x, row)
			}
			cells = append(cells, row)
			row = []int{}
		}
		repeat = 0
	}
	pat := Pattern{
		x:     x,
		y:     y,
		cells: cells,
	}

	return pat, nil
}

func addPaddingToRow(x int, row []int) []int {
	missingCells := x - len(row)
	for range missingCells {
		row = append(row, 0)
	}
	return row
}
