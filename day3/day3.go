package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
)

type Node struct {
	zero int
	one  int
}

func buildDataFromArray(arr *[]string) []Node {
	var list = make([]Node, 12)
	for _, e := range *arr {
		for i, c := range e {
			if string(c) == "1" {
				list[i].one++
			} else {
				list[i].zero++
			}
		}
	}

	return list
}

func buildData(scanner *bufio.Scanner) ([]Node, []string) {
	var list = make([]Node, 12)
	var all []string

	for scanner.Scan() {
		line := scanner.Text()
		all = append(all, line)
		for i, c := range line {
			if string(c) == "1" {
				list[i].one++
			} else {
				list[i].zero++
			}
		}
	}

	return list, all
}

func partOne(scanner *bufio.Scanner) (string, string) {
	treeRoot, _ := buildData(scanner)
	gamma := ""
	epsilon := ""

	for _, e := range treeRoot {
		if e.one < e.zero {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	return gamma, epsilon
}

func eliminateWithoutPrefix(arr *[]string, prefix string) []string {
	var result []string

	for _, e := range *arr {
		if e[0:len(prefix)] == prefix {
			result = append(result, e)
		}
	}

	return result
}

func part2(scanner *bufio.Scanner) (string, string) {
	treeRoot, all := buildData(scanner)
	cotwoCandidates := all
	oxygenCandidates := all

	cotwoRoot := treeRoot
	oxygenRoot := treeRoot

	cotwo := ""
	oxygen := ""

	for i := 0; i < 12; i++ {
		if len(cotwoCandidates) > 1 {
			if cotwoRoot[i].zero <= cotwoRoot[i].one {
				cotwo += "0"
			} else {
				cotwo += "1"
			}

			cotwoCandidates = eliminateWithoutPrefix(&cotwoCandidates, cotwo)
			cotwoRoot = buildDataFromArray(&cotwoCandidates)
		}

		if len(oxygenCandidates) > 1 {
			if oxygenRoot[i].zero <= oxygenRoot[i].one {
				oxygen += "1"
			} else {
				oxygen += "0"
			}

			oxygenCandidates = eliminateWithoutPrefix(&oxygenCandidates, oxygen)
			oxygenRoot = buildDataFromArray(&oxygenCandidates)
		}
	}

	return cotwoCandidates[0], oxygenCandidates[0]
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
		gs, es := partOne(scanner)

		gamma, err := strconv.ParseInt(gs, 2, 64)
		if err != nil {
			log.Fatalln("Gamma not a number!")
		}

		epsilon, err := strconv.ParseInt(es, 2, 64)
		if err != nil {
			log.Fatalln("Epsilon not an int!")
		}

		log.Printf("Gamma = %d, Epsilon = %d [Consumption = %d]", gamma, epsilon, gamma*epsilon)
	} else {
		co, ox := part2(scanner)

		cotwo, err := strconv.ParseInt(co, 2, 64)
		if err != nil {
			log.Fatalln("CO not a number!")
		}

		oxygen, err := strconv.ParseInt(ox, 2, 64)
		if err != nil {
			log.Fatalln("Oxygen not an int!")
		}

		log.Printf("CO2 = %d, Oxygen = %d [Score = %d]", cotwo, oxygen, cotwo*oxygen)
	}
}
