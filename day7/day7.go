package main

import (
	"bufio"
	"flag"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func CalculateScore(points []int, target int, useCrabFormula bool) uint64 {
	var score uint64 = 0

	for _, e := range points {
		normalScore := math.Abs(float64(e) - float64(target))
		if !useCrabFormula {
			score += uint64(normalScore)
		} else {
			score += uint64((2 + normalScore - 1) * (normalScore / 2))
		}
	}

	return score
}

func TakeAverageDistance(points []int) int {
	n := 0
	sum := 0
	remainder := points

	for len(remainder) > 1 {
		rhs := remainder[0]
		remainder = remainder[1:]
		for _, e := range remainder {
			sum += int(math.Abs(float64(rhs) - float64(e)))
			n++
		}
	}

	return sum / n
}

func MaxPoint(points []int) int {
	max := points[0]

	for _, e := range points {
		if max < e {
			max = e
		}
	}

	return max
}

func MinPoint(points []int) int {
	min := points[0]

	for _, e := range points {
		if min > e {
			min = e
		}
	}

	return min
}

func WalkUp(points []int, target int, bound int, useCrabFormula bool) int {
	lastScore := CalculateScore(points, target, useCrabFormula)
	localTarget := target

	for localTarget <= bound {
		newScore := CalculateScore(points, localTarget+1, useCrabFormula)
		log.Printf("[UP] Target = %d, Score = %d\n", localTarget+1, newScore)
		if newScore <= lastScore {
			localTarget++
			lastScore = newScore
		} else {
			break
		}
	}

	return localTarget
}

func WalkDown(points []int, target int, bound int, useCrabFormula bool) int {
	lastScore := CalculateScore(points, target, useCrabFormula)
	localTarget := target

	for localTarget >= bound {
		newScore := CalculateScore(points, localTarget-1, useCrabFormula)
		log.Printf("[DOWN] Target = %d, Score = %d\n", localTarget-1, newScore)

		if newScore <= lastScore {
			localTarget--
			lastScore = newScore
		} else {
			break
		}
	}

	return localTarget
}

func WalkGradient(points []int, target int, useCrabFormula bool) int {
	max := MaxPoint(points)
	min := MinPoint(points)

	proposedAbove := WalkUp(points, target, max, useCrabFormula)
	proposedBelow := WalkDown(points, target, min, useCrabFormula)

	scoreAbove := CalculateScore(points, proposedAbove, useCrabFormula)
	scoreBelow := CalculateScore(points, proposedBelow, useCrabFormula)

	if scoreAbove <= scoreBelow {
		return proposedAbove
	} else {
		return proposedBelow
	}

}

func part2(scanner *bufio.Scanner) {
	scanner.Scan()
	rawLine := scanner.Text()
	rawInts := strings.Split(rawLine, ",")

	var points []int

	for _, e := range rawInts {
		point, _ := strconv.ParseInt(e, 10, 64)

		points = append(points, int(point))
	}

	avg := TakeAverageDistance(points)
	min := MinPoint(points)
	max := MaxPoint(points)

	startingPoint := ((max - avg) + (min + avg)) / 2

	log.Printf("Starting at: %d", startingPoint)
	proposedMinimum := WalkGradient(points, startingPoint, true)

	log.Printf("Proposed = %d, score = %d\n", proposedMinimum, CalculateScore(points, proposedMinimum, true))
}

func part1(scanner *bufio.Scanner) {
	scanner.Scan()
	rawLine := scanner.Text()
	rawInts := strings.Split(rawLine, ",")

	var points []int

	for _, e := range rawInts {
		point, _ := strconv.ParseInt(e, 10, 64)

		points = append(points, int(point))
	}

	avg := TakeAverageDistance(points)
	min := MinPoint(points)
	max := MaxPoint(points)

	startingPoint := (max - avg) + (min+avg)/2

	log.Printf("Starting at: %d", startingPoint)
	proposedMinimum := WalkGradient(points, startingPoint, false)

	log.Printf("Proposed = %d, score = %d\n", proposedMinimum, CalculateScore(points, proposedMinimum, false))

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
		part1(scanner)
	} else {
		part2(scanner)
	}
}
