package main

import (
	"bufio"
	"fmt"
	"io"
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

func ParseRleFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file. Exiting... %v", err)
		os.Exit(1)
	}
	defer file.Close()

	var sb strings.Builder

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "Error reading file. Exiting... %v", err)
			os.Exit(1)
		}
		var trimmedLine string
		trimmedLine = strings.TrimSpace(line)
		var lastLine bool

		if trimmedLine[len(trimmedLine)-1] == '!' {
			lastLine = true
			trimmedLine = strings.Trim(trimmedLine, "!")
		}

		if trimmedLine[0] == '2' {
			sb.WriteString(trimmedLine)
		}

		if lastLine {
			break
		}
	}
	return sb.String()
}

func main() {
	println("hello world")
}
