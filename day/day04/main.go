package main

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"os"
)

func main() {
	f, err := os.Open("inputs/day04.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %w", err))
	}

	defer f.Close()

	s := parse(f)
	fmt.Printf("Part 1: %d\n", part1(s))
	fmt.Printf("Part 2: %d\n", part2(s))

}

func part1(s *space) int {
	var starts []image.Point
	acc := 0

	for pt, val := range s.points {
		if val == 'X' {
			starts = append(starts, pt)
		}
	}

	for _, pt := range starts {
		acc += s.tracedAt("MAS", pt)
	}

	return acc
}

func part2(s *space) int {
	acc := 0

	for pt, val := range s.points {
		if val != 'A' {
			continue
		}
		if s.detect_cross_mas(pt) {
			acc++
		}
	}

	return acc
}

func (s *space) detect_cross_mas(pt image.Point) bool {
	nw := s.points[pt.Add(image.Pt(-1, -1))]
	sw := s.points[pt.Add(image.Pt(-1, 1))]
	ne := s.points[pt.Add(image.Pt(1, -1))]
	se := s.points[pt.Add(image.Pt(1, 1))]

	// ensure that all the characters are either 'M' or 'S'
	for _, v := range []rune{nw, sw, ne, se} {
		if v != 'M' && v != 'S' {
			return false
		}
	}

	// if either diagonal starts and ends with the same character
	// it's not a match
	if nw == se || ne == sw {
		return false
	}

	return true
}

type space struct {
	points     map[image.Point]rune
	rows, cols int
}

func parse(r io.Reader) *space {
	s := &space{
		points: make(map[image.Point]rune),
	}

	b := bufio.NewScanner(r)

	for b.Scan() {
		s.cols = len(b.Text())
		for col, c := range b.Text() {
			s.points[image.Pt(col, s.rows)] = c
		}
		s.rows++
	}

	return s
}

func (s *space) tracedAt(target string, pt image.Point) int {
	directions := []image.Point{
		{1, 0},   // right
		{-1, 0},  // left
		{0, 1},   // down
		{0, -1},  // up
		{1, 1},   // down-right
		{-1, -1}, // up-left
		{1, -1},  // up-right
		{-1, 1},  // down-left
	}
	acc := 0
	for _, d := range directions {
		if s.traceDirection(target, pt, d) {
			acc++
		}
	}
	return acc
}

func (s *space) traceDirection(target string, pt, direction image.Point) bool {
	x := pt.Add(direction)
	for _, val := range target {
		if s.points[x] != val {
			return false
		}
		x = x.Add(direction)
	}
	return true
}
