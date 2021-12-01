package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(scanner *bufio.Scanner) string {
	hasFirstMeasurment := false
	var previous int64 = 0
	var totalIncrease int64 = 0

	for scanner.Scan() {
		thisLine := scanner.Text()

		parsedInt, error := strconv.ParseInt(thisLine, 10, 64)

		if error != nil {
			log.Panicln("Data not an int!")
		}

		if previous < parsedInt && hasFirstMeasurment {
			totalIncrease++
		}

		previous = parsedInt
		hasFirstMeasurment = true
	}

	return fmt.Sprintf("%d", totalIncrease)
}

func shiftDown(arr *[4]int64) {
	for i, num := range arr {
		if i-1 < 0 {
			continue
		}

		arr[i-1] = num
	}
}

func sum(arr []int64) int64 {
	var total int64 = 0

	for _, num := range arr {
		total += num
	}

	return total
}

func part2(scanner *bufio.Scanner) string {
	var fillCnt = 0
	var shiftingBuffer [4]int64

	var total int64 = 0

	for scanner.Scan() {
		thisLine := scanner.Text()

		parsedInt, error := strconv.ParseInt(thisLine, 10, 64)
		if error != nil {
			log.Fatalln("Data not an int!")
		}

		shiftingBuffer[fillCnt] = parsedInt
		fillCnt++

		if fillCnt > 3 {
			if sum(shiftingBuffer[0:3]) < sum(shiftingBuffer[1:4]) {
				total++
			}

			shiftDown(&shiftingBuffer)
			fillCnt = 3
		}

	}

	return fmt.Sprintf("%d", total)
}

func main() {
	inputPath := flag.String("input", "input.txt", "Input file for puzzle")
	partTwo := flag.Bool("part2", false, "Move onto puzzle 2")

	flag.Parse()

	file, error := os.Open(*inputPath)

	if error != nil {
		log.Fatalln("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	var result string

	if !*partTwo {
		log.Println("Solving part 1")
		result = part1(scanner)
	} else {
		log.Println("Solving part 2")
		result = part2(scanner)
	}

	log.Printf("Puzzle solution: %s", result)
}
