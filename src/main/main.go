package main

import (
	"day01"
	"day02"
	"fmt"
	"log"
	"time"
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

func runDay02() {
	fmt.Println("Read input 2!")
	input, err := readDailyInput(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := day02.ChecksumCalculator(input)
	fmt.Printf("Result for part 1 is: %d", result) // returns 6696

	fmt.Println()

	secondResult := day02.BoxIdComparer(input)
	fmt.Printf("Result for part 2 is: %s", secondResult) // returns bvnfawcnyoeyudzrpgslimtkj
}

func main() {
	start := time.Now()
	// runDay01()
	runDay02()
	elapsed := time.Since(start)
	fmt.Println()
	log.Printf("Took %s", elapsed)
}
