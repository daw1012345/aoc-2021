package main

import (
	"bufio"
	"os"
	"testing"
)

func TestWithExampleP2(t *testing.T) {
	file, err := os.Open("test_data/example.txt")

	if err != nil {
		t.Fatal("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	h, d := part2(scanner)

	if h != 15 || d != 60 {
		t.Fatalf("Invalid h (%d) and d (%d) [expeted h=15, d=60]", h, d)
	}

}

func TestWithExampleP1(t *testing.T) {
	file, err := os.Open("test_data/example.txt")

	if err != nil {
		t.Fatal("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	h, d := part1(scanner)

	if h != 15 || d != 10 {
		t.Fatalf("Invalid h (%d) and d (%d) [expeted h=15, d=10]", h, d)
	}

}
