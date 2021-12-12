package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

type Node struct {
	Name        string
	Connections []*Node
	IsSmall     bool
	IsStart     bool
}

func (n *Node) IsConnectedTo(neigh string) bool {
	for _, conn := range n.Connections {
		if conn.Name == neigh {
			return true
		}
	}

	return false
}

func (n *Node) IsEnd() bool {
	return n.Name == "end"
}

func CheckShouldVisit(path []*Node, next *Node, canVisitSmallCavesTwice bool) bool {

	var visitCount = make(map[*Node]int)

	if next.IsStart {
		return false
	}

	for _, node := range path {
		if _, ok := visitCount[node]; ok {
			visitCount[node]++
		} else {
			visitCount[node] = 1
		}
	}

	if _, ok := visitCount[next]; ok {
		visitCount[next]++
	} else {
		visitCount[next] = 1
	}

	if !canVisitSmallCavesTwice {
		for key, value := range visitCount {
			if key.IsSmall && value >= 2 {
				return false
			}
		}

		return true
	}

	exhaustedCave := false
	for key, value := range visitCount {
		if key.IsSmall && value > 2 {
			return false
		}

		if key.IsSmall && value == 2 && !exhaustedCave {
			exhaustedCave = true
			continue
		}

		if key.IsSmall && value == 2 && exhaustedCave {
			return false
		}

	}

	return true
}

func ExploreUntilTheEnd(prefix []*Node, canVisitSmallCavesTwice bool) [][]*Node {
	var ret [][]*Node

	if prefix[len(prefix)-1].IsEnd() {
		ret = append(ret, prefix)

		return ret
	}

	for _, next := range prefix[len(prefix)-1].Connections {
		if !CheckShouldVisit(prefix, next, canVisitSmallCavesTwice) {
			continue
		}

		tmpPrefix := append(prefix, next)
		contd := ExploreUntilTheEnd(tmpPrefix, canVisitSmallCavesTwice)

		ret = append(ret, contd...)
	}

	return ret

}

func CountPaths(start *Node, canVisitSmallCavesTwice bool) int {
	var CheckedPaths [][]*Node

	CheckedPaths = ExploreUntilTheEnd([]*Node{start}, canVisitSmallCavesTwice)

	return len(CheckedPaths)
}

func ParseConnections(scanner *bufio.Scanner) *Node {
	var allNodes = make(map[string]*Node)
	for scanner.Scan() {
		line := scanner.Text()

		conn := strings.Split(line, "-")

		zero, zeroExists := allNodes[conn[0]]
		one, oneExists := allNodes[conn[1]]

		if !zeroExists {
			allNodes[conn[0]] = &Node{conn[0], []*Node{}, strings.ToLower(conn[0]) == conn[0], conn[0] == "start"}
			zero = allNodes[conn[0]]
		}

		if !oneExists {
			allNodes[conn[1]] = &Node{conn[1], []*Node{}, strings.ToLower(conn[1]) == conn[1], conn[1] == "start"}
			one = allNodes[conn[1]]
		}

		if !zero.IsConnectedTo(conn[1]) {
			zero.Connections = append(zero.Connections, one)
		}

		if !one.IsConnectedTo(conn[0]) {
			one.Connections = append(one.Connections, zero)
		}
	}

	return allNodes["start"]
}

func Part1(start *Node) {
	log.Printf("[PATH] Starting conns: %d", len(start.Connections))
	log.Printf("[PATH] Found paths: %d", CountPaths(start, false))
}

func Part2(start *Node) {
	log.Printf("[PATH] Starting conns: %d", len(start.Connections))
	log.Printf("[PATH] Found paths: %d", CountPaths(start, true))
}

func main() {
	inputPath := flag.String("input", "day12/input.txt", "Input file for puzzle")
	partTwo := flag.Bool("part2", false, "Move onto puzzle 2")

	flag.Parse()

	file, err := os.Open(*inputPath)

	if err != nil {
		log.Panicln("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	if !*partTwo {
		Part1(ParseConnections(scanner))
	} else {
		Part2(ParseConnections(scanner))
	}
}
