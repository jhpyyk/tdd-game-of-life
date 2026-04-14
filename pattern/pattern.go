package pattern

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/jhpyyk/tdd-game-of-life/utils"
)

type Pattern struct {
	x          int
	y          int
	generation int
	cells      [][]int
}

func (pattern *Pattern) GetNextGeneration() Pattern {
	padded1 := utils.PadMatrix(pattern.cells)
	padded2 := utils.PadMatrix(padded1)

	result := utils.ZeroMatrix(len(padded1), len(padded1[0]))

	kernel := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	for i := range padded1 {
		for j := range len(padded1[0]) {
			threeByThreeSlice := utils.GetSubMatrix(padded2, i, j, 3, 3)
			neighbors := utils.FrobeniusInnerProduct(threeByThreeSlice, kernel)

			if padded1[i][j] == 1 {
				switch {
				case neighbors < 2:
					result[i][j] = 0
				case neighbors == 2, neighbors == 3:
					result[i][j] = 1
				case neighbors > 3:
					result[i][j] = 0
				}
			}
			if padded1[i][j] == 0 && neighbors == 3 {
				result[i][j] = 1
			}
		}
	}
	result = utils.TrimPadding(result)

	nextGen := Pattern{
		x:          len(result),
		y:          len(result[0]),
		generation: pattern.generation + 1,
		cells:      result,
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

func FromString(x int, y int, patternString string) (Pattern, error) {
	stripped := utils.StripPattern(patternString)
	length := len(stripped)
	if length < x*y {
		return Pattern{}, fmt.Errorf("Error: pattern length %v does not match dimensions x: %v, y: %v", length, x, y)
	}

	cells := [][]int{}
	row := []int{}
	for i, c := range stripped {
		if string(c) == "#" {
			row = append(row, 1)
		}
		if string(c) == "." {
			row = append(row, 0)
		}
		if (i+1)%x == 0 {
			cells = append(cells, row)
			row = []int{}
		}
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
