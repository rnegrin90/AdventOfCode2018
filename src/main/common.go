package main

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
)

func readDailyInput(day int) ([]string, error) {
	baseFilePath := "C:/Users/rnegr/Documents/repos/AdventOfCode2018/src/day%02[1]d/input%02[1]d.txt"
	file, err := os.Open(fmt.Sprintf(baseFilePath, day))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

func manhattanDistance(start image.Point, end image.Point) float64 {
	return math.Abs(float64(end.X-start.X)) + math.Abs(float64(end.Y-start.Y))
}
