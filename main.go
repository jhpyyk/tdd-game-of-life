package main

import (
	"fmt"
	"os"
	"strings"
)

type Pattern struct {
	x     int
	y     int
	cells [][]string
}

func (pattern *Pattern) toString() string {
	var sb strings.Builder

	for _, row := range pattern.cells {
		for _, col := range row {
			sb.WriteString(col)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func ParseRleFile(path string) Pattern {
	_, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file. Exiting... %v", err)
		os.Exit(1)
	}

	fakePattern := Pattern{
		x: 2,
		y: 2,
		cells: [][]string{
			{"#", "#"},
			{"#", "#"},
		},
	}

	return fakePattern
}

func main() {
	println("hello world")
}
