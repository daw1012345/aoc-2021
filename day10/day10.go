package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
)

var closingChars = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var scoresInvalid = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var scoresAutocomplete = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func Part2(scanner *bufio.Scanner) {

	var points []int

	for scanner.Scan() {
		line := scanner.Text()
		var parseList []string
		localTotal := 0

		for _, c := range line {
			// If it's an opening character
			if _, ok := closingChars[string(c)]; ok {
				parseList = append(parseList, string(c))
				continue
			}

			// It's a closing character
			if closingChars[parseList[len(parseList)-1]] != string(c) {
				log.Printf("[PARSER P2] Discarding invalid line")
				parseList = parseList[:0]
				break
			} else {
				parseList = parseList[:len(parseList)-1]
			}
		}

		if len(parseList) == 0 { // Valid line
			continue
		}

		for i := len(parseList) - 1; i >= 0; i-- {
			correctClosingChar := closingChars[parseList[i]]

			localTotal *= 5
			localTotal += scoresAutocomplete[correctClosingChar]
		}

		points = append(points, localTotal)
	}

	sort.Ints(points)
	log.Printf("[PARSER P1] Score = %d", points[len(points)/2])
}

func Part1(scanner *bufio.Scanner) {

	points := 0

	for scanner.Scan() {
		line := scanner.Text()
		var parseList []string

		for _, c := range line {
			// If it's an opening character
			if _, ok := closingChars[string(c)]; ok {
				parseList = append(parseList, string(c))
				continue
			}
			// It's a closing character
			if closingChars[parseList[len(parseList)-1]] != string(c) {
				log.Printf("[PARSER P1] Found %s but expected %s ", string(c), closingChars[parseList[len(parseList)-1]])
				points += scoresInvalid[string(c)]
				break
			} else {
				parseList = parseList[:len(parseList)-1]
			}
		}
	}

	log.Printf("[PARSER P1] Score = %d", points)
}

func main() {
	inputPath := flag.String("input", "day10/input.txt", "Input file for puzzle")
	partTwo := flag.Bool("part2", false, "Move onto puzzle 2")

	flag.Parse()

	file, err := os.Open(*inputPath)

	if err != nil {
		log.Panicln("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	if !*partTwo {
		Part1(scanner)
	} else {
		Part2(scanner)
	}
}
