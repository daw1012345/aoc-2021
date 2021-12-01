package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	inputPath := flag.String("input", "input.txt", "Input file for puzzle")
	partTwo := flag.Bool("part2", false, "Move onto puzzle 2")

	flag.Parse()

	file, error := os.Open(*inputPath)

	if error != nil {
		log.Panicln("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)
}
