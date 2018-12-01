package day01

import (
	"log"
	"testing"
)

func TestFrequencyCalculator(t *testing.T) {
	tables := []struct {
		input  []string
		output int
	}{
		{[]string{"+1", "-2", "+3", "+1"}, 3},
		{[]string{"+1", "+1", "+1"}, 3},
		{[]string{"+1", "+1", "-2"}, 0},
		{[]string{"-1", "-2", "-3"}, -6},
	}

	for _, table := range tables {
		result, err := FrequencyCalculator(table.input)
		if err != nil {
			log.Fatal(err)
			return
		}
		if result != table.output {
			t.Errorf("Got %d, expected %d", result, table.output)
		}
	}
}

func TestRepeatedFrequenciesCalculator(t *testing.T) {
	tables := []struct {
		input  []string
		output int
	}{
		{[]string{"+1", "-2", "+3", "+1"}, 2},
		{[]string{"+1", "-1"}, 0},
		{[]string{"+3", "+3", "+4", "-2", "-4"}, 10},
		{[]string{"-6", "+3", "+8", "+5", "-6"}, 5},
		{[]string{"+7", "+7", "-2", "-7", "-4"}, 14},
	}

	for _, table := range tables {
		result, err := RepeatedFrequenciesCalculator(table.input)
		if err != nil {
			log.Fatal(err)
			return
		}
		if result != table.output {
			t.Errorf("Got %d, expected %d", result, table.output)
		}
	}
}
