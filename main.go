package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RawPattern struct {
	x             int
	y             int
	patternString string
}

func ParseRleFile(path string) RawPattern {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file. Exiting... %v", err)
		os.Exit(1)
	}
	defer file.Close()

	pat := RawPattern{}

	lines := readLines(file)
	headerIndex, err := findHeaderLineIndex(lines)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v. Exiting...", err)
		os.Exit(1)
	}

	xdim, ydim, err := parseDimensionsFromHeader(lines[headerIndex])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v. Exiting...", err)
		os.Exit(1)
	}
	pat.x = xdim
	pat.y = ydim

	patternString := parsePatternString(lines[headerIndex+1:])
	pat.patternString = patternString
	return pat
}

func parsePatternString(lines []string) string {
	var sb strings.Builder
	for _, line := range lines {

		trimmedLine := strings.TrimSpace(line)

		if isLastLine(trimmedLine) {
			lineWithoutExclamation := strings.Trim(trimmedLine, "!")
			sb.WriteString(lineWithoutExclamation)
			break
		}
	}
	return sb.String()
}

func parseDimensionsFromHeader(header string) (int, int, error) {
	splitted := strings.Split(header, ",")
	xdimString := stripRegex(splitted[0])
	xdim, err := strconv.Atoi(xdimString)
	if err != nil {
		return 0, 0, fmt.Errorf("Error parsing x dimension from file %v. Exiting... %v", xdim, err)
	}
	ydimString := stripRegex(splitted[1])
	ydim, err := strconv.Atoi(ydimString)
	if err != nil {
		return 0, 0, fmt.Errorf("Error parsing y dimension from file %v. Exiting... %v", ydim, err)
	}
	return xdim, ydim, nil
}

func findHeaderLineIndex(lines []string) (int, error) {
	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine[0] == 'x' {
			return i, nil
		}
	}
	return 0, errors.New("Cannot parse header")
}

func isLastLine(line string) bool {
	return strings.Contains(line, "!")
}

func readLines(file *os.File) []string {
	lines := []string{}
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
		lines = append(lines, line)
	}
	return lines
}

func stripRegex(in string) string {
	reg, _ := regexp.Compile("[^0-9 ]+")
	numeric := reg.ReplaceAllString(in, "")
	trimmed := strings.TrimSpace(numeric)
	return trimmed
}

func main() {
	println("hello world")
}
