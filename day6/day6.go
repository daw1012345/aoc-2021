package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func runSimulation(fish [9]uint64) [9]uint64 {
	res := fish
	totalNewFish := fish[0]

	for i, e := range res {
		if i-1 < 0 {
			continue
		}

		res[i-1] = e
	}

	res[8] = totalNewFish
	res[6] += totalNewFish

	return res

}

func sum(fish [9]uint64) uint64 {
	var total uint64 = 0

	for _, e := range fish {
		total += e
	}

	return total
}

func simulate(scanner *bufio.Scanner, num int) {
	scanner.Scan()
	line := scanner.Text()
	rawLineComponents := strings.Split(line, ",")

	var initState []int

	for _, e := range rawLineComponents {
		parsed, _ := strconv.ParseInt(e, 10, 64)
		initState = append(initState, int(parsed))
	}

	var fish [9]uint64

	for _, e := range initState {
		fish[e]++
	}

	for i := 1; i <= num; i++ {
		fish = runSimulation(fish)
		log.Printf("[FISH] Day = %d, Count = %d\n", i, sum(fish))
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
		simulate(scanner, 80)
	} else {
		simulate(scanner, 256)
	}
}
