package day02

import (
	"reflect"
	"testing"
)

func TestLetterCounter(t *testing.T) {
	tables := []struct {
		input  string
		output map[string]int
	}{
		{"abcdef", map[string]int{"a": 1, "b": 1, "c": 1, "d": 1, "e": 1, "f": 1}},
		{"bababc", map[string]int{"a": 2, "b": 3, "c": 1}},
		{"abbcde", map[string]int{"a": 1, "b": 2, "c": 1, "d": 1, "e": 1}},
		{"abcccd", map[string]int{"a": 1, "b": 1, "c": 3, "d": 1}},
		{"aabcdd", map[string]int{"a": 2, "b": 1, "c": 1, "d": 2}},
		{"abcdee", map[string]int{"a": 1, "b": 1, "c": 1, "d": 1, "e": 2}},
		{"ababab", map[string]int{"a": 3, "b": 3}},
	}

	for _, table := range tables {
		result := LetterCounter(table.input)

		eq := reflect.DeepEqual(result, table.output)
		if !eq {
			t.Errorf("Got %d, expected %d", result, table.output)
		}
	}
}

func TestChecksumCalculator(t *testing.T) {
	tables := []struct {
		input  []string
		output int
	}{
		{
		[]string {
				"abcdef",
				"bababc",
				"abbcde",
				"abcccd",
				"aabcdd",
				"abcdee",
				"ababab",
			},
		12,
		},
	}

	for _, table := range tables {
		result := ChecksumCalculator(table.input)

		if result != table.output {
			t.Errorf("Got %d, expected %d", result, table.output)
		}
	}
}

func TestDifferentCharCount(t *testing.T) {
	tables := []struct {
		firstString  string
		secondString string
		output int
	}{
		{"abcde", "axcye", 2},
		{"fghij", "fguij", 1},
	}

	for _, table := range tables {
		result := DifferentCharCount(table.firstString, table.secondString)

		if result != table.output {
			t.Errorf("Got %d, expected %d", result, table.output)
		}
	}
}

func TestCommonLetters(t *testing.T) {
	tables := []struct {
		firstString  string
		secondString string
		output string
	}{
		{"fghij", "fguij", "fgij"},
	}

	for _, table := range tables {
		result := CommonLetters(table.firstString, table.secondString)

		if result != table.output {
			t.Errorf("Got %s, expected %s", result, table.output)
		}
	}
}

func TestBoxIdComparer(t *testing.T) {
	tables := []struct {
		input  []string
		output string
	}{
		{
			[]string{
				"abcde",
				"fghij",
				"klmno",
				"pqrst",
				"fguij",
				"axcye",
				"wvxyz",
			},
			"fgij",
		},
	}

	for _, table := range tables {
		result := BoxIdComparer(table.input)

		if result != table.output {
			t.Errorf("Got %s, expected %s", result, table.output)
		}
	}
}
