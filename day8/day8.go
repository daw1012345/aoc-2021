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

type NumberComponent int

const (
	TOP               NumberComponent = 0
	BOTTOM                            = 1
	MIDDLE                            = 2
	SIDE_TOP_LEFT                     = 3
	SIDE_TOP_RIGHT                    = 4
	SIDE_BOTTOM_LEFT                  = 5
	SIDE_BOTTOM_RIGHT                 = 6
)

type Number struct {
	intValue   int
	components []NumberComponent
	unique     bool
}

var allNumbers [10]Number = [10]Number{
	{0, []NumberComponent{TOP, SIDE_TOP_RIGHT, SIDE_BOTTOM_RIGHT, BOTTOM, SIDE_BOTTOM_LEFT, SIDE_TOP_LEFT}, false},
	{1, []NumberComponent{SIDE_TOP_RIGHT, SIDE_BOTTOM_RIGHT}, true},
	{2, []NumberComponent{TOP, SIDE_TOP_RIGHT, MIDDLE, SIDE_BOTTOM_LEFT, BOTTOM}, false},
	{3, []NumberComponent{TOP, SIDE_TOP_RIGHT, MIDDLE, SIDE_BOTTOM_RIGHT, BOTTOM}, false},
	{4, []NumberComponent{SIDE_TOP_LEFT, SIDE_TOP_RIGHT, MIDDLE, SIDE_BOTTOM_RIGHT}, true},
	{5, []NumberComponent{TOP, SIDE_TOP_LEFT, MIDDLE, SIDE_BOTTOM_RIGHT, BOTTOM}, false},
	{6, []NumberComponent{TOP, SIDE_TOP_LEFT, MIDDLE, SIDE_BOTTOM_LEFT, BOTTOM, SIDE_BOTTOM_RIGHT}, false},
	{7, []NumberComponent{TOP, SIDE_TOP_RIGHT, SIDE_BOTTOM_RIGHT}, true},
	{8, []NumberComponent{TOP, SIDE_TOP_RIGHT, SIDE_TOP_LEFT, MIDDLE, SIDE_BOTTOM_RIGHT, SIDE_BOTTOM_LEFT, BOTTOM}, true},
	{9, []NumberComponent{TOP, SIDE_TOP_RIGHT, SIDE_TOP_LEFT, MIDDLE, SIDE_BOTTOM_RIGHT, BOTTOM}, false}}

func Difference(target []NumberComponent, known []NumberComponent) []NumberComponent {
	var difference []NumberComponent

	for _, te := range target {
		found := false
		for _, ke := range known {
			if ke == te {
				found = true
			}
		}

		if !found {
			difference = append(difference, te)
		}
	}

	return difference
}

func RemoveAllFromString(target string, cutset string) string {
	res := target
	for _, c := range cutset {
		res = strings.ReplaceAll(res, string(c), "")
	}

	return res
}

func IdentifyRemainingUniqCombinations(nums *map[int]string, allWires []string) {
	// We know 1, 4, 7, 8 from the start. We need 2, 3, 5, 6, 9. We can identify them based on `difference against known number`

	for _, definedNumber := range allNumbers {
		if definedNumber.unique {
			continue
		}

		expectedDifferenceWithOne := len(Difference(definedNumber.components, allNumbers[1].components))
		expectedDifferenceWithFour := len(Difference(definedNumber.components, allNumbers[4].components))
		expectedDifferenceWithSeven := len(Difference(definedNumber.components, allNumbers[7].components))
		expectedDifferenceWithEight := len(Difference(definedNumber.components, allNumbers[8].components))

		for _, wire := range allWires {
			actualDifferenceWithOne := len(RemoveAllFromString(wire, (*nums)[1]))
			actualDifferenceWithFour := len(RemoveAllFromString(wire, (*nums)[4]))
			actualDifferenceWithSeven := len(RemoveAllFromString(wire, (*nums)[7]))
			actualDifferenceWithEight := len(RemoveAllFromString(wire, (*nums)[8]))

			if len(wire) == len(definedNumber.components) && actualDifferenceWithOne == expectedDifferenceWithOne && actualDifferenceWithFour == expectedDifferenceWithFour && actualDifferenceWithSeven == expectedDifferenceWithSeven && actualDifferenceWithEight == expectedDifferenceWithEight {
				(*nums)[definedNumber.intValue] = wire
				break
			}
		}

	}
}

func HasAllComponents(lhs string, rhs string) bool {
	hasAll := len(lhs) == len(rhs)

	for _, v := range rhs {
		hasAll = hasAll && strings.Contains(lhs, string(v))
	}

	return hasAll
}

func part2(scanner *bufio.Scanner) {
	var outputs []int

	for scanner.Scan() {
		nums := make(map[int]string)

		line := scanner.Text()
		rawInfo := strings.Split(line, "|")

		rawDisplayedNums := strings.Split(rawInfo[1], " ")
		rawWireConfig := strings.Split(rawInfo[0], " ")

		for _, dNum := range rawWireConfig {
			for _, kNum := range allNumbers {
				if len(dNum) == len(kNum.components) && kNum.unique {
					nums[kNum.intValue] = dNum
				}
			}
		}

		IdentifyRemainingUniqCombinations(&nums, rawWireConfig)

		tmpString := ""
		for _, e := range rawDisplayedNums {
			if len(e) == 0 {
				continue
			}

			for key, value := range nums {
				if HasAllComponents(value, e) {
					tmpString += fmt.Sprintf("%d", key)
					break
				}
			}
		}
		n, _ := strconv.ParseInt(tmpString, 10, 64)

		outputs = append(outputs, int(n))
	}

	var sum uint64 = 0
	for _, e := range outputs {
		sum += uint64(e)
	}

	log.Printf("[NUMS] [P2] Sum = %d", sum)

}

func part1(scanner *bufio.Scanner) {
	var discoveredNums [10]int

	for scanner.Scan() {
		line := scanner.Text()
		rawInfo := strings.Split(line, "|")

		rawDisplayedNums := strings.Split(rawInfo[1], " ")

		for _, dNum := range rawDisplayedNums {
			for _, kNum := range allNumbers {
				if len(dNum) == len(kNum.components) && kNum.unique {
					discoveredNums[kNum.intValue]++
				}
			}
		}
	}

	sum := 0
	for i, n := range discoveredNums {
		log.Printf("[NUMS] Num = %d, Count = %d\n", i, n)
		sum += n
	}

	log.Printf("[NUMS] Sum = %d", sum)
}

func main() {
	inputPath := flag.String("input", "day8/input.txt", "Input file for puzzle")
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
