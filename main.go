package main

import (
	"bufio"
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
	var sb strings.Builder
	var patternStartsOnNextLine bool

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
		var lastLine bool

		trimmedLine = strings.TrimSpace(line)

		if trimmedLine[len(trimmedLine)-1] == '!' {
			lastLine = true
			trimmedLine = strings.Trim(trimmedLine, "!")
		}

		if patternStartsOnNextLine {
			sb.WriteString(trimmedLine)
		}

		if trimmedLine[0] == 'x' {
			patternStartsOnNextLine = true
			splitted := strings.Split(trimmedLine, ",")
			xdimString := stripRegex(splitted[0])
			xdim, err := strconv.Atoi(xdimString)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing x dimension from file %v. Exiting... %v", xdim, err)
				os.Exit(1)
			}
			ydimString := stripRegex(splitted[1])
			ydim, err := strconv.Atoi(ydimString)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing y dimension from file %v. Exiting... %v", ydim, err)
				os.Exit(1)
			}
			pat.x = xdim
			pat.y = ydim
		}

		if lastLine {
			pat.patternString = sb.String()
			break
		}
	}
	return pat
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
