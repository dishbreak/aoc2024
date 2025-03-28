package main

import (
	"fmt"
	"image"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestPart1(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	s := parse(strings.NewReader(input))
	assert.Equal(t, 18, part1(s))
}

func TestPart2(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	s := parse(strings.NewReader(input))
	assert.Equal(t, 9, part2(s))
}

type TestCases struct {
	TestCrossMas []struct {
		Result bool   `yaml:"result"`
		Input  string `yaml:"input"`
	} `yaml:"test_cross_mas"`
}

func TestCrossMas(t *testing.T) {
	f, err := os.Open("testdata/test_cases.yaml")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	d := yaml.NewDecoder(f)
	c := &TestCases{}
	d.Decode(c)

	for i, testCase := range c.TestCrossMas {
		t.Run(fmt.Sprint("test case", i), func(t *testing.T) {
			space := parse(strings.NewReader(testCase.Input))
			assert.Equal(t, testCase.Result, space.detect_cross_mas(image.Pt(1, 1)))
		})
	}

}
