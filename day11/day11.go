package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
)

type Board = [10][10]int

func ParseBoard(scanner *bufio.Scanner) Board {
	row := 0
	var board Board
	for scanner.Scan() {
		line := scanner.Text()

		for i, c := range line {
			num, _ := strconv.ParseInt(string(c), 10, 64)
			board[row][i] = int(num)
		}

		row++
	}

	return board
}

func SimulationStep(board *Board) int {
	var flashes [][2]int
	lastLen := 0

	// Increase all
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			board[r][c]++
		}
	}

	// Go until stop flashing
	for {
		lastLen = len(flashes)

		for r, re := range board {
			for c, ce := range re {
				if CheckFlashed([2]int{r, c}, &flashes) {
					continue
				}

				if ce > 9 {
					IncreaseNeighbors([2]int{r, c}, board)
					flashes = append(flashes, [2]int{r, c})
				}

			}
		}

		if lastLen == len(flashes) {
			break
		}
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if board[r][c] > 9 {
				board[r][c] = 0
			}
		}
	}

	return len(flashes)

}

func CheckFlashed(trgt [2]int, flashes *[][2]int) bool {
	for _, flash := range *flashes {
		if trgt[0] == flash[0] && trgt[1] == flash[1] {
			return true
		}
	}

	return false
}

func IncreaseNeighbors(trgt [2]int, board *Board) {
	variations := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

	for _, variation := range variations {
		newCoord := [2]int{trgt[0] - variation[0], trgt[1] - variation[1]}

		if !(newCoord[0] >= 0 && newCoord[0] < len(board)) {
			continue
		}

		if !(newCoord[1] >= 0 && newCoord[1] < len(board[newCoord[0]])) {
			continue
		}

		board[newCoord[0]][newCoord[1]]++
	}
}

func Part2(board Board) {
	for i := 1; ; i++ {
		tmp := SimulationStep(&board)
		if tmp == 10*10 {
			log.Printf("[OCTOPI P2] Synchronised at %d\n", i)
			break
		}
	}
}

func Part1(board Board) {
	total := 0
	for i := 1; i <= 100; i++ {
		tmp := SimulationStep(&board)
		total += tmp
		log.Printf("[OCTOPI P1] Flashes after step %d: %d\n", i, tmp)
	}

	log.Printf("[OCTOPI P1] Total flashes = %d\n", total)

}

func main() {
	inputPath := flag.String("input", "day11/input.txt", "Input file for puzzle")
	partTwo := flag.Bool("part2", false, "Move onto puzzle 2")

	flag.Parse()

	file, err := os.Open(*inputPath)

	if err != nil {
		log.Panicln("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)
	board := ParseBoard(scanner)

	if !*partTwo {
		Part1(board)
	} else {
		Part2(board)
	}
}
