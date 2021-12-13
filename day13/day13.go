package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coords = [][2]int

type Fold struct {
	axis string
	line int
}

func Parse(scanner *bufio.Scanner) (Coords, []Fold) {
	var AllFolds []Fold
	var AllCoords Coords
	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.Trim(line, "\n ")) <= 0 {
			continue
		}

		// This is a fold
		if strings.Contains(line, "fold") {
			components := strings.Split(line, " ")

			foldParams := strings.Split(components[len(components)-1], "=")
			parsedLine, _ := strconv.ParseInt(foldParams[1], 10, 64)

			AllFolds = append(AllFolds, Fold{foldParams[0], int(parsedLine)})
			continue
		}

		// Otherwise, this is a coordinate

		unparsedCoords := strings.Split(line, ",")

		parsedX, _ := strconv.ParseInt(unparsedCoords[0], 10, 64)
		parsedY, _ := strconv.ParseInt(unparsedCoords[1], 10, 64)

		AllCoords = append(AllCoords, [2]int{int(parsedX), int(parsedY)})

	}

	return AllCoords, AllFolds
}

func DoPrintBoard(board Coords) {

	// Such a shitty way to do this... but it is the easiest

	maxX := 0
	maxY := 0
	for _, e := range board {
		if maxX < e[0] {
			maxX = e[0]
		}

		if maxY < e[1] {
			maxY = e[1]
		}
	}

	var matrix = make([][]bool, maxY+1)

	for i, _ := range matrix {
		matrix[i] = make([]bool, maxX+1)
	}

	for _, e := range board {
		matrix[e[1]][e[0]] = true
	}

	for _, row := range matrix {
		for _, col := range row {
			if col {
				fmt.Print("#")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}

}

func DoFold(board Coords, fold Fold) Coords {
	var final Coords

	for _, e := range board {
		if fold.axis == "y" {
			if fold.line > e[1] {
				final = append(final, e)
				continue
			}

			final = append(final, [2]int{e[0], fold.line + (fold.line - e[1])})
		} else {
			if fold.line > e[0] {
				final = append(final, e)
				continue
			}

			final = append(final, [2]int{fold.line + (fold.line - e[0]), e[1]})
		}
	}

	return final
}

func Dedup(board Coords) Coords {
	var deduped Coords

	for _, dup := range board {
		IsDuped := false
		for _, dedup := range deduped {
			if dedup[0] == dup[0] && dedup[1] == dup[1] {
				IsDuped = true
				break
			}
		}

		if !IsDuped {
			deduped = append(deduped, dup)
		}
	}

	return deduped

}

func Part1(coords Coords, folds []Fold) {
	answer := DoFold(coords, folds[0])

	log.Printf("[FOLD P1] Number visible: %d\n", len(Dedup(answer)))
}

func Part2(coords Coords, folds []Fold) {
	var answer = coords
	for _, fold := range folds {
		log.Printf("[FOLD P1] Doing fold: %s=%d", fold.axis, fold.line)
		answer = DoFold(answer, fold)
		log.Printf("[FOLD P1] Number visible: %d\n", len(answer))
	}

	DoPrintBoard(Dedup(answer))

}

func main() {
	inputPath := flag.String("input", "day13/input.txt", "Input file for puzzle")
	partTwo := flag.Bool("part2", false, "Move onto puzzle 2")

	flag.Parse()

	file, err := os.Open(*inputPath)

	if err != nil {
		log.Panicln("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	c, f := Parse(scanner)

	if !*partTwo {
		Part1(c, f)
	} else {
		Part2(c, f)
	}
}
