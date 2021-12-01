package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

//file, error := os.Open(*inputPath)
//
//	if error != nil {
//		log.Panicln("Invalid input file!")
//	}
//
//	scanner := bufio.NewScanner(file)

func TestPart2Uneven(t *testing.T) {
	file, error := os.Open("test_data/not_multiple_of_3.txt")

	if error != nil {
		t.Fatalf("Can't open test data")
	}

	scanner := bufio.NewScanner(file)

	result := part2(scanner)

	pint, error := strconv.ParseInt(result, 10, 64)

	if error != nil {
		t.Fatalf("Expected int but got something else: %s", result)
	}

	if pint != 1 {
		t.Fatalf("Got invalid number: %d [expected =1]", pint)
	}
}

func TestPart2StandardFirstBigger(t *testing.T) {
	file, error := os.Open("test_data/standard_1.txt")

	if error != nil {
		t.Fatalf("Can't open test data")
	}

	scanner := bufio.NewScanner(file)

	result := part2(scanner)

	pint, error := strconv.ParseInt(result, 10, 64)

	if error != nil {
		t.Fatalf("Expected int but got something else: %s", result)
	}

	if pint != 0 {
		t.Fatalf("Got invalid number: %d [expected =0]", pint)
	}
}

func TestPart2StandardSecondBigger(t *testing.T) {
	file, error := os.Open("test_data/standard_2.txt")

	if error != nil {
		t.Fatalf("Can't open test data")
	}

	scanner := bufio.NewScanner(file)

	result := part2(scanner)

	pint, error := strconv.ParseInt(result, 10, 64)

	if error != nil {
		t.Fatalf("Expected int but got something else: %s", result)
	}

	if pint != 1 {
		t.Fatalf("Got invalid number: %d [expected =1]", pint)
	}
}

func TestPart2Insufficient(t *testing.T) {
	file, error := os.Open("test_data/insufficient_data.txt")

	if error != nil {
		t.Fatalf("Can't open test data")
	}

	scanner := bufio.NewScanner(file)

	result := part2(scanner)

	pint, error := strconv.ParseInt(result, 10, 64)

	if error != nil {
		t.Fatalf("Expected int but got something else: %s", result)
	}

	if pint != 0 {
		t.Fatalf("Got invalid number: %d [expected =0]", pint)
	}
}
