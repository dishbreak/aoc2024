package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day01.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %w", err))
	}
	defer f.Close()
	left, right := parse(f)

	fmt.Printf("Part 1: %d\n", part1(left, right))
	fmt.Printf("Part 2: %d\n", part2(left, right))
}

func parse(r io.Reader) (left, right []int) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		parts := strings.Fields(s.Text())
		l, _ := strconv.Atoi(parts[0])
		left = append(left, l)
		r, _ := strconv.Atoi(parts[1])
		right = append(right, r)
	}

	return
}

func part1(left, right []int) int {
	acc := 0

	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		acc += abs(left[i] - right[i])
	}

	return acc
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func part2(left, right []int) int {
	hits := make(map[int]int)
	for _, val := range right {
		hits[val] += 1
	}

	acc := 0
	for _, val := range left {
		acc += val * hits[val]
	}

	return acc
}
