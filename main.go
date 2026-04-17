package main

import (
	"log"
	"os"
	"strconv"

	"github.com/jhpyyk/tdd-game-of-life/pattern"
	"github.com/jhpyyk/tdd-game-of-life/rle_parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing filepath")
	}
	if len(os.Args) < 3 {
		log.Fatal("missing generations")
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
	for range generations {
		pat = pat.GetNextGeneration()
	}
	rle, err := pat.ToRLE()
	if err != nil {
		log.Fatal(err)
	}
	println(rle)
}
