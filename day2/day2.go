package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2(scanner *bufio.Scanner) (int64, int64) {
	var horizontal int64 = 0
	var depth int64 = 0
	var aim int64 = 0

	for scanner.Scan() {
		line := scanner.Text()
		segments := strings.Split(line, " ")
		num, err := strconv.ParseInt(segments[1], 10, 64)
		if err != nil {
			log.Fatalln("Invalid number in file")
		}

		switch segments[0] {
		case "forward":
			horizontal += num
			depth += aim * num
		case "down":
			//depth += num
			aim += num
		case "up":
			//depth -= num
			aim -= num
		}
	}

	return horizontal, depth
}

func part1(scanner *bufio.Scanner) (int64, int64) {
	var horizontal int64 = 0
	var depth int64 = 0

	for scanner.Scan() {
		line := scanner.Text()
		segments := strings.Split(line, " ")
		num, err := strconv.ParseInt(segments[1], 10, 64)
		if err != nil {
			log.Fatalln("Invalid number in file")
		}

		switch segments[0] {
		case "forward":
			horizontal += num
		case "down":
			depth += num
		case "up":
			depth -= num
		}
	}

	return horizontal, depth
}

func main() {
	inputPath := flag.String("input", "input.txt", "Input file for puzzle")
	partTwo := flag.Bool("part2", false, "Move onto puzzle 2")

	flag.Parse()

	file, err := os.Open(*inputPath)

	if err != nil {
		log.Panicln("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	if !*partTwo {
		h, d := part1(scanner)
		log.Printf("[Puzzle 1] Horizontal: %d, depth: %d [solution=%d]\n", h, d, h*d)
	} else {
		h, d := part2(scanner)
		log.Printf("[Puzzle 2] Horizontal: %d, depth: %d [solution=%d]\n", h, d, h*d)
	}
}
