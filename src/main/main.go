package main

import (
	"day01"
	"fmt"
)

func runDay01() {
	fmt.Println("Read input 1!")
	input, err := readDailyInput(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := day01.FrequencyCalculator(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result for part 1 is: %d", result) // returns 582

	fmt.Println()

	result, err = day01.RepeatedFrequenciesCalculator(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result for part 2 is: %d", result) // returns 488
}

func main() {
	runDay01()
}
