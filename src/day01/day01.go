package day01

import (
	"container/ring"
	"strconv"
)

func FrequencyCalculator(input []string) (int, error) {
	frequencies, err := readFrequencies(input)
	if err != nil {
		return 0, err
	}

	currentFrequency := 0
	for _, element := range frequencies {
		currentFrequency += element
	}
	return currentFrequency, nil
}

func RepeatedFrequenciesCalculator(input []string) (int, error) {
	frequencies, err := readFrequencies(input)
	if err != nil {
		return 0, err
	}
	ringFreq := ring.New(len(frequencies))
	for i := 0; i < ringFreq.Len(); i++ {
		ringFreq.Value = frequencies[i]
		ringFreq = ringFreq.Next()
	}

	currentFrequency := 0
	var freqStore = map[int]int{
		0: 1,
	}
	for {
		currentFrequency += ringFreq.Value.(int)
		if freqStore[currentFrequency] >= 1 {
			return currentFrequency, nil
		}
		freqStore[currentFrequency] += 1
		ringFreq = ringFreq.Next()
	}
}

func readFrequencies(input []string) ([]int, error) {
	var symbol = map[string]int{
		"+": 1,
		"-": -1,
	}

	var frequencies []int
	for _, element := range input {
		multiplier := symbol[string(element[0])]
		value, err := strconv.Atoi(element[1:])
		if err != nil {
			return nil, err
		}

		frequencies = append(frequencies, value*multiplier)
	}
	return frequencies, nil
}
