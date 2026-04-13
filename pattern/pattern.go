package pattern

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Pattern struct {
	x          int
	y          int
	generation int
	cells      [][]int
}

func (pattern *Pattern) GetNextGeneration() Pattern {
	nextGen := Pattern{
		pattern.x,
		pattern.y,
		pattern.generation + 1,
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

func stripPattern(pattern string) string {
	var sb strings.Builder
	for _, c := range pattern {
		if c == '.' || c == '#' {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func FromString(x int, y int, patternString string) (Pattern, error) {
	stripped := stripPattern(patternString)
	length := len(stripped)
	if length < x*y {
		return Pattern{}, fmt.Errorf("Error: pattern length %v does not match dimensions x: %v, y: %v", length, x, y)
	}

	cells := [][]int{}
	for i := range x {
		row := []int{}
		for j := range y {
			if stripped[(i+1)*(j+1)-1] == '.' {
				row = append(row, 0)
			}
			if stripped[(i+1)*(j+1)-1] == '#' {
				row = append(row, 1)
			}
		}
		cells = append(cells, row)
	}
	pat := Pattern{
		x,
		y,
		0,
		cells,
	}
	return pat, nil
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
		x:          x,
		y:          y,
		generation: 0,
		cells:      cells,
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
