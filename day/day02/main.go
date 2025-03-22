package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day02.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open file: %w", err))
	}
	defer f.Close()
	reports := parse(f)

	fmt.Printf("Part 1: %d\n", part1(reports))
	fmt.Printf("Part 2: %d\n", part2(reports))

}

func part1(reports [][]int) int {
	acc := 0

	for _, report := range reports {
		if safe_report(report) {
			acc++
		}
	}

	return acc
}

func safe_report(report []int) (safe bool) {
	safe = all(report, increasing) || all(report, decreasing)
	safe = safe && all(report, safe_jump_v1)
	return safe
}

func part2(reports [][]int) int {
	acc := 0

	for _, report := range reports {
		if safe_report(report) {
			acc++
			continue
		}
		for i := range report {
			var r []int
			r = append(r, report[:i]...)
			r = append(r, report[i+1:]...)
			if safe_report(r) {
				acc++
				break
			}
		}
	}

	return acc
}

func parse(r io.Reader) (result [][]int) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		var l []int
		for _, pt := range strings.Fields(s.Text()) {
			k, _ := strconv.Atoi(pt)
			l = append(l, k)
		}
		result = append(result, l)
	}

	return
}

type rule func(int, int) bool

func increasing(a, b int) bool {
	return a <= b
}

func decreasing(a, b int) bool {
	return a >= b
}

func safe_jump_v1(a, b int) bool {
	d := abs(b - a)
	return d >= 1 && d <= 3
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func all(s []int, rules ...rule) bool {
	if len(s) < 2 {
		return false
	}

	for _, r := range rules {
		for i := 1; i < len(s); i++ {
			if !r(s[i-1], s[i]) {
				return false
			}
		}
	}

	return true
}
