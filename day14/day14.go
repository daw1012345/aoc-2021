package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

type Mapping = map[string]string
type Count = map[string]int

func Parse(scanner *bufio.Scanner) (Count, Mapping, string) {
	scanner.Scan()

	begin := scanner.Text()
	var ass = make(Mapping)
	var board = make(Count)

	i := 0
	for i+1 < len(begin) {
		key := begin[i : i+2]
		if _, ok := board[key]; ok {
			board[key]++
		} else {
			board[key] = 1
		}

		i++
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) <= 0 {
			continue
		}

		rawMapping := strings.Split(line, "->")

		ass[strings.TrimSpace(rawMapping[0])] = strings.TrimSpace(rawMapping[1])
	}

	return board, ass, string(begin[len(begin)-1])
}

func DoSim(llist Count, ass Mapping, last string, count int) {
	i := 0

	final := llist
	for i < count {
		log.Printf("[POLYMER] Step = %d", i)
		var additional = make(Count)

		for key, value := range final { // For all existing components
			if value <= 0 {
				continue
			}

			if _, ok := ass[key]; ok { // If it exists in mapping
				combinations := [2]string{string(key[0]) + ass[key], ass[key] + string(key[1])}
				for _, e := range combinations {
					additional[e] += value
				}
			} else {
				additional[key] = value
			}
		}

		final = additional

		i++
	}

	letterMap := make(Count)

	for key, value := range final {
		letterMap[string(key[0])] += value
	}

	letterMap[last] += 1

	min := -1
	max := -1

	for _, value := range letterMap {
		if max < value {
			max = value
		}

		if value < min || min == -1 {
			min = value
		}
	}

	log.Printf("[POLYMER] After %d steps, score = %d\n", count, max-min)
}

func main() {
	inputPath := flag.String("input", "day14/input.txt", "Input file for puzzle")
	partTwo := flag.Bool("part2", false, "Move onto puzzle 2")

	flag.Parse()

	file, err := os.Open(*inputPath)

	if err != nil {
		log.Panicln("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	list, ass, last := Parse(scanner)
	if !*partTwo {
		DoSim(list, ass, last, 10)
	} else {
		DoSim(list, ass, last, 40)
	}
}
