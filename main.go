package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jhpyyk/tdd-game-of-life/pattern"
	"github.com/jhpyyk/tdd-game-of-life/rle_parser"
	"github.com/jhpyyk/tdd-game-of-life/visualizer"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing filepath")
	}
	if len(os.Args) < 3 {
		log.Fatal("missing generations")
	}
	draw := false
	if len(os.Args) == 4 && os.Args[3] == "--draw" {
		draw = true
	}
	path := os.Args[1]
	raw, err := rle_parser.ParseRleFile(path)
	if err != nil {
		log.Fatal(err)
	}
	pat, err := pattern.ParsePatternFromRLEPatternString(raw.X, raw.Y, raw.PatternString)
	if err != nil {
		log.Fatal(err)
	}
	generations, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	if draw {
		visualizer.EnterAltScreen()
		defer visualizer.ExitAltScreen()
	}
	for range generations {
		pat = pat.GetNextGeneration()
		if pat.X < 1 || pat.Y < 0 {
			break
		}
		if draw {
			patString, err := pat.ToStringFixedSize(40)
			if err != nil {
				log.Fatal(err)
			}
			frame := visualizer.Frame{
				Header:  fmt.Sprintf("Generation %d", pat.Generation),
				Content: patString,
			}
			visualizer.DrawFrame(os.Stdout, frame)
			time.Sleep(100 * time.Millisecond)
		}
	}

	rle, err := pat.ToRLE()
	if err != nil {
		log.Fatal(err)
	}
	println(rle)
}
