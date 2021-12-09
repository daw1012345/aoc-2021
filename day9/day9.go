package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
)

type Coords struct {
	row int
	col int
}

func WalkHorizontal(start Coords, direction int, board [][]int, exclusions *[]Coords) int {
	sum := 0

	if CheckCanVisit(start, *exclusions) {
		*exclusions = append(*exclusions, start)
		sum++
	}

	currentCol := start.col + direction

	for currentCol < len(board[start.row]) && currentCol >= 0 && board[start.row][currentCol] != 9 && CheckCanVisit(Coords{start.row, currentCol}, *exclusions) {
		sum++

		*exclusions = append(*exclusions, Coords{start.row, currentCol})

		if start.row-1 >= 0 && board[start.row-1][currentCol] != 9 {
			sum += WalkVertical(Coords{start.row, currentCol}, -1, board, exclusions)
		}

		if start.row+1 < len(board) && board[start.row+1][currentCol] != 9 {
			sum += WalkVertical(Coords{start.row, currentCol}, 1, board, exclusions)
		}

		currentCol += direction
	}

	return sum
}

func CheckCanVisit(field Coords, exclusions []Coords) bool {
	for _, e := range exclusions {
		if field.row == e.row && field.col == e.col {
			return false
		}
	}

	return true
}

func WalkVertical(start Coords, direction int, board [][]int, exclusions *[]Coords) int {
	sum := 0

	if CheckCanVisit(start, *exclusions) {
		*exclusions = append(*exclusions, start)
		sum++
	}

	currentRow := start.row + direction

	for currentRow < len(board) && currentRow >= 0 && board[currentRow][start.col] != 9 && CheckCanVisit(Coords{currentRow, start.col}, *exclusions) {
		sum++

		*exclusions = append(*exclusions, Coords{currentRow, start.col})

		if start.col-1 >= 0 && board[currentRow][start.col-1] != 9 {
			sum += WalkHorizontal(Coords{currentRow, start.col}, -1, board, exclusions)
		}

		if start.col+1 < len(board[currentRow]) && board[currentRow][start.col+1] != 9 {
			sum += WalkHorizontal(Coords{currentRow, start.col}, 1, board, exclusions)
		}

		currentRow += direction
	}

	return sum

}

func GetTopThree(basins []int) int {

	sort.Ints(basins)
	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func GetBasinScore(data [][]int, lows []Coords) int {

	var basins []int
	exclusions := lows

	for _, low := range lows {
		total := 0
		total += WalkHorizontal(low, 1, data, &exclusions) + WalkHorizontal(low, -1, data, &exclusions) + WalkVertical(low, 1, data, &exclusions) + WalkVertical(low, -1, data, &exclusions) + 1
		basins = append(basins, total)
	}

	return GetTopThree(basins)
}

func FindLowPoints(data [][]int) []Coords {
	var res []Coords

	for rowi, row := range data {
		for coli, val := range row {
			lowerThanAbove := rowi-1 < 0 || rowi-1 >= 0 && data[rowi-1][coli] > val
			lowerThanBelow := rowi+1 >= len(data) || rowi+1 < len(data) && data[rowi+1][coli] > val
			lowerThanRight := coli+1 >= len(data[rowi]) || coli+1 < len(data[rowi]) && data[rowi][coli+1] > val
			lowerThanLeft := coli-1 < 0 || coli-1 >= 0 && data[rowi][coli-1] > val

			if lowerThanAbove && lowerThanBelow && lowerThanRight && lowerThanLeft {
				res = append(res, Coords{rowi, coli})
			}
		}
	}

	return res
}

func part1(scanner *bufio.Scanner) {
	var data [][]int // [row][column]

	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		data = append(data, make([]int, 0))
		for _, field := range line {
			pfield, _ := strconv.ParseInt(string(field), 10, 64)
			data[row] = append(data[row], int(pfield))
		}

		row++
	}

	lows := FindLowPoints(data)

	sum := 0

	for _, low := range lows {
		sum += data[low.row][low.col] + 1
		//log.Printf("[LOWS] Row = %d, Col = %d", low.row, low.col)
	}

	log.Printf("[TUBES] Score = %d", sum)
}

func part2(scanner *bufio.Scanner) {
	var data [][]int // [row][column]

	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		data = append(data, make([]int, 0))
		for _, field := range line {
			pfield, _ := strconv.ParseInt(string(field), 10, 64)
			data[row] = append(data[row], int(pfield))
		}

		row++
	}

	lows := FindLowPoints(data)

	score := GetBasinScore(data, lows)

	log.Printf("[TUBES2] Score = %d", score)

}

func main() {
	inputPath := flag.String("input", "day9/input.txt", "Input file for puzzle")
	partTwo := flag.Bool("part2", false, "Move onto puzzle 2")

	flag.Parse()

	file, err := os.Open(*inputPath)

	if err != nil {
		log.Panicln("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	if !*partTwo {
		part1(scanner)
	} else {
		part2(scanner)
	}
}
