package main

import (
	"log"
	"os"

	"github.com/jhpyyk/tdd-game-of-life/pattern"
	"github.com/jhpyyk/tdd-game-of-life/rle_parser"
)

func main() {
	path := os.Args[1]
	raw, err := rle_parser.ParseRleFile(path)
	if err != nil {
		log.Fatal(err)
	}
	pat, err := pattern.ParsePattern(raw.X, raw.Y, raw.PatternString)
	if err != nil {
		log.Fatal(err)
	}
	println(pat.ToString())
}
