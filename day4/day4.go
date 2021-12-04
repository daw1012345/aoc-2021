package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board [5][5]int

func LinearSearch(arr []int, item int) bool {
	for _, e := range arr {
		if e == item {
			return true
		}
	}

	return false
}

func (b Board) CheckHasBingo(pulls []int) bool {

	// Check if there's bingo in column
	for _, column := range b {
		matches := 0
		for _, e := range column {
			if LinearSearch(pulls, e) {
				matches++
			} else {
				break
			}
		}

		if matches == 5 {
			return true
		}
	}

	for row := 0; row < 5; row++ {
		matches := 0
		for col := 0; col < 5; col++ {
			if LinearSearch(pulls, b[col][row]) {
				matches++
			} else {
				break
			}
		}

		if matches == 5 {
			return true
		}
	}

	return false
}

func GetBoards(scanner *bufio.Scanner) []Board {
	var allBoards []Board
	var thisBoard Board

	offset := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 {
			continue
		}

		row := strings.Split(line, " ")

		i := 0
		for _, c := range row {
			if len(string(c)) <= 0 {
				continue
			}

			val, _ := strconv.ParseInt(string(c), 10, 64)
			thisBoard[i][offset] = int(val)
			i++
		}

		offset++

		if offset >= 5 {
			offset = 0
			allBoards = append(allBoards, thisBoard)
		}
	}

	return allBoards
}

func PareStrArray(arr []string) []int {
	var result []int

	for _, e := range arr {
		parsed, _ := strconv.ParseInt(e, 10, 64)
		result = append(result, int(parsed))
	}

	return result
}

func CalculateScore(board Board, pulls []int) int {
	last := pulls[len(pulls)-1]

	totalUnchecked := 0

	for _, col := range board {
		for _, e := range col {
			if !LinearSearch(pulls, e) {
				totalUnchecked += e
			}
		}
	}

	return last * totalUnchecked

}

func partOne(scanner *bufio.Scanner) {
	scanner.Scan()
	pullsTxt := scanner.Text()

	allPulls := strings.Split(pullsTxt, ",")
	allBoards := GetBoards(scanner)

	for i := 1; i <= len(allPulls); i++ {
		pullsSlice := allPulls[0:i]
		parsedPulls := PareStrArray(pullsSlice)

		for i, board := range allBoards {
			if board.CheckHasBingo(parsedPulls) {
				score := CalculateScore(board, parsedPulls)
				log.Printf("[BINGO] Board = %d, Last = %d, Score = %d", i, parsedPulls[len(parsedPulls)-1], score)
				return
			}
		}
	}
}

func Part2(scanner *bufio.Scanner) {
	scanner.Scan()
	pullsTxt := scanner.Text()

	allPulls := strings.Split(pullsTxt, ",")
	allBoards := GetBoards(scanner)
	acceptable := allBoards

	for i := 1; i <= len(allPulls); i++ {
		pullsSlice := allPulls[0:i]
		parsedPulls := PareStrArray(pullsSlice)

		var noBingo []Board
		for _, board := range acceptable {
			if !board.CheckHasBingo(parsedPulls) {
				noBingo = append(noBingo, board)
			} else {
				if len(acceptable) == 1 {
					log.Printf("[BINGO LAST] Score = %d\n", CalculateScore(acceptable[0], parsedPulls))
					return
				}
			}
		}

		acceptable = noBingo
	}
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
		partOne(scanner)
	} else {
		Part2(scanner)
	}
}
