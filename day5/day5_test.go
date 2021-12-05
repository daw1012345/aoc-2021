package main

import (
	"bufio"
	"os"
	"testing"
)

func TestBasicHvV(t *testing.T) {
	file, err := os.Open("test_data/vertical_vs_horizontal.txt")

	if err != nil {
		t.Fatalf("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	res := Part1(scanner)

	if res != 1 {
		t.Fatalf("Expected 1 intersection!")
	}
}

func TestVerticalNoIntersection(t *testing.T) {
	file, err := os.Open("test_data/vertical_no_intersection.txt")

	if err != nil {
		t.Fatalf("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	res := Part1(scanner)

	if res != 0 {
		t.Fatalf("Expected 0 intersections!")
	}
}

func TestHorizontalNoIntersection(t *testing.T) {
	file, err := os.Open("test_data/horiz_no_int.txt")

	if err != nil {
		t.Fatalf("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	res := Part1(scanner)

	if res != 0 {
		t.Fatalf("Expected 0 intersections!")
	}
}

func TestHorizontalOverlap(t *testing.T) {
	file, err := os.Open("test_data/horiz_overlap.txt")

	if err != nil {
		t.Fatalf("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	res := Part1(scanner)

	if res != 2 {
		t.Fatalf("Expected 2 intersections!")
	}
}

func TestVerticalOverlap(t *testing.T) {
	file, err := os.Open("test_data/vert_overlap.txt")

	if err != nil {
		t.Fatalf("Invalid input file!")
	}

	scanner := bufio.NewScanner(file)

	res := Part1(scanner)

	if res != 11 {
		t.Fatalf("Expected 2 intersections!")
	}
}
