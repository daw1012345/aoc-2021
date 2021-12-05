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

type Segment struct {
	start [2]int
	end   [2]int
}

func (s Segment) WalkLine(countDiagonal bool) [][2]int {
	isStraight := false
	for i := 0; i < 2; i++ {
		isStraight = isStraight || s.start[i] == s.end[i]
	}

	if isStraight {
		return s.WalkStraightLine()
	} else if countDiagonal {
		return s.WalkDiagonalLine()
	} else {
		return make([][2]int, 0)
	}

}

func (s Segment) WalkDiagonalLine() [][2]int {
	var res [][2]int

	m := (s.end[1] - s.start[1]) / (s.end[0] - s.start[0]) // 1 or -1

	startX := int(math.Min(float64(s.start[0]), float64(s.end[0])))
	endX := int(math.Max(float64(s.start[0]), float64(s.end[0])))

	for i := startX; i != endX+1; i++ {
		var tmp [2]int

		currY := m*(i-s.start[0]) + s.start[1]

		tmp[0] = i
		tmp[1] = currY

		res = append(res, tmp)
	}

	return res
}

func (s Segment) WalkStraightLine() [][2]int {
	var res [][2]int

	if s.start[0] == s.end[0] {
		var offset int = (s.end[1] - s.start[1]) / int(math.Abs(float64(s.end[1]-s.start[1])))
		x := s.start[0]

		for i := s.start[1]; i != s.end[1]+offset; i += offset {
			var tmp [2]int
			tmp[0] = x
			tmp[1] = i
			res = append(res, tmp)
		}
	} else {
		var offset int = (s.end[0] - s.start[0]) / int(math.Abs(float64(s.end[0]-s.start[0])))
		y := s.start[1]

		for i := s.start[0]; i != s.end[0]+offset; i += offset {
			var tmp [2]int
			tmp[0] = i
			tmp[1] = y
			res = append(res, tmp)

		}
	}

	return res

}

func ParseSegmentFromRawString(in string) Segment {
	rawCoords := strings.Split(in, "->")
	var segment Segment

	isEnd := false
	for _, e := range rawCoords {
		coords := strings.Split(strings.TrimSpace(e), ",")
		for i, e := range coords {
			parsed, _ := strconv.ParseInt(e, 10, 64)
			if !isEnd {
				segment.start[i] = int(parsed)
			} else {
				segment.end[i] = int(parsed)
			}
		}

		isEnd = true
	}

	return segment
}

func ParseSegmentsFromFile(scanner *bufio.Scanner) []Segment {
	var allSegments []Segment

	for scanner.Scan() {
		line := scanner.Text()

		allSegments = append(allSegments, ParseSegmentFromRawString(line))

	}

	return allSegments
}

func CountIntersects(occupied [][2]int) int {
	remainder := occupied
	var intersects [][2]int

	for len(remainder) > 1 {
		hasInter := false
		hasDup := false

		compare := remainder[0]
		remainder = remainder[1:]

		for _, inter := range remainder { // Check if intersects
			if inter[0] == compare[0] && inter[1] == compare[1] {
				hasInter = true
			}
		}

		for _, inter := range intersects { // Check if dup
			if inter[0] == compare[0] && inter[1] == compare[1] {
				hasDup = true
			}
		}

		if hasInter && !hasDup {
			intersects = append(intersects, compare)
		}
	}

	return len(intersects)
}

func GetOccupiedSpaces(segments []Segment, countDiagonal bool) [][2]int {

	var occupiedSpaces [][2]int

	for _, seg := range segments {
		occupiedSpaces = append(occupiedSpaces, seg.WalkLine(countDiagonal)...)
	}

	return occupiedSpaces
}

func Part1(scanner *bufio.Scanner) int {
	allSegments := ParseSegmentsFromFile(scanner)

	return CountIntersects(GetOccupiedSpaces(allSegments, false))
}

func Part2(scanner *bufio.Scanner) int {
	allSegments := ParseSegmentsFromFile(scanner)

	return CountIntersects(GetOccupiedSpaces(allSegments, true))
}

func main() {
	inputPath := flag.String("input", "day5/input.txt", "Input file for puzzle")
	partTwo := flag.Bool("part2", false, "Move onto puzzle 2")

	flag.Parse()

	file, err := os.Open(*inputPath)

	if err != nil {
		log.Panicln("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	var res int
	if !*partTwo {
		res = Part1(scanner)
	} else {
		res = Part2(scanner)
	}
	log.Printf("Intersections = %d\n", res)
}
