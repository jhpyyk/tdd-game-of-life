package main

import (
	"log"
	"os"

	"github.com/jhpyyk/tdd-game-of-life/pattern_parser"
	"github.com/jhpyyk/tdd-game-of-life/rle_parser"
)

func main() {
	path := os.Args[1]
	raw := rle_parser.ParseRleFile(path)
	pat, err := pattern_parser.ParsePattern(raw.X, raw.Y, raw.PatternString)
	if err != nil {
		log.Fatal(err)
	}
	println(pat.ToString())
}
