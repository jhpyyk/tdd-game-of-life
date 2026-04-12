package rle_parser

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
	X             int
	Y             int
	PatternString string
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
	pat.X = xdim
	pat.Y = ydim

	patternString := parsePatternString(lines[headerIndex+1:])
	pat.PatternString = patternString
	return pat
}

func parsePatternString(lines []string) string {
	var sb strings.Builder
	for _, line := range lines {

		trimmedLine := strings.TrimSpace(line)
		lineWithoutExclamation := strings.Trim(trimmedLine, "!")
		sb.WriteString(lineWithoutExclamation)
	}
	return sb.String()
}

func parseDimensionsFromHeader(header string) (int, int, error) {
	splitted := strings.Split(header, ",")
	xdimString, err := parseIntFromString(splitted[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Error parsing x dimension from file %v. Exiting... %v", xdimString, err)
	}
	xdim, err := strconv.Atoi(xdimString)
	if err != nil {
		return 0, 0, fmt.Errorf("Error parsing x dimension from file %v. Exiting... %v", xdim, err)
	}
	ydimString, err := parseIntFromString(splitted[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Error parsing y dimension from file %v. Exiting... %v", ydimString, err)
	}
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

func parseIntFromString(in string) (string, error) {
	reg, err := regexp.Compile("[^0-9 ]+")
	if err != nil {
		return "", fmt.Errorf("Failed to parse int from string %v", err)
	}
	numeric := reg.ReplaceAllString(in, "")
	trimmed := strings.TrimSpace(numeric)
	return trimmed, nil
}
