package pattern_parser

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Pattern struct {
	x     int
	y     int
	cells [][]rune
}

func (pattern *Pattern) ToString() string {
	var sb strings.Builder
	for _, row := range pattern.cells {
		for _, cell := range row {
			sb.WriteRune(cell)
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func ParsePattern(x int, y int, pattern string) (Pattern, error) {
	cells := [][]rune{}
	repeat := 0
	row := []rune{}
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
		if runeValue == 'o' {
			for range repeat {
				row = append(row, '#')
			}
			repeat = 0
			continue
		}
		if runeValue == 'b' {
			for range repeat {
				row = append(row, '.')
			}
			repeat = 0
			continue
		}
		if runeValue == '$' || runeValue == '!' {
			cells = append(cells, row)
			row = []rune{}
			repeat = 0
			continue
		}
	}
	pat := Pattern{
		x:     x,
		y:     y,
		cells: cells,
	}

	return pat, nil
}
