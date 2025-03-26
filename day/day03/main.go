package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day03.txt")
	if err != nil {
		panic(fmt.Errorf("Failed to open input file: %w", err))
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(fmt.Errorf("failed to read input file: %w", err))
	}
	input := string(bytes)
	input = strings.ReplaceAll(input, "\n", "")

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))

}

func part1(input string) int {
	return extract(input)
}

func part2(input string) int {
	return extract(clean(input))
}

var extractor = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func extract(input string) int {
	acc := 0

	hits := extractor.FindAllStringSubmatch(input, -1)
	for _, hit := range hits {
		a, _ := strconv.Atoi(hit[1])
		b, _ := strconv.Atoi(hit[2])
		acc += (a * b)
	}

	return acc
}

func clean(input string) string {
	var sb strings.Builder

	// start by splitting the input into strings that start with do()
	for _, part := range strings.Split(input, "do()") {
		// for each split, we know that there are 0 or more don't() instructions
		// and zero do() instrcutions
		// that means no instructions after the first don't() are effective
		// if we split the split again on don't() and discard all but the first split,
		// we will have only the effective instructions.
		sb.WriteString(strings.Split(part, "don't()")[0])
	}

	return sb.String()
}
