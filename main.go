package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Result of Advent #1: %v", advent1())
}

func advent1() int {
	input, err := readLines("./input1.txt")
	output := []string{}

	if err != nil {
		log.Fatalf("An error occurred %v", err)
	}

	for _, line := range input {
		re := regexp.MustCompile("[0-9]+")
		reResult := re.FindAllString(line, -1)
		allInts := []string{}
		allInts = append(allInts, reResult...)
		splits := splitDigits(allInts)

		first := splits[0]
		last := splits[len(splits)-1]

		output = append(output, first+last)
	}

	return addTogetherDigits(output)
}

func advent1Secnd(input []string) int {
	fmt.Print(input)
	return 0
}

func addTogetherDigits(digits []string) int {
	convertedValues := []int{}
	result := 0

	for _, e := range digits {
		i, err := strconv.Atoi(e)
		if err != nil {
			// ... handle error
			panic(err)
		}
		convertedValues = append(convertedValues, i)
	}

	for _, e := range convertedValues {
		result = result + e
	}

	return result
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func splitDigits(slice []string) []string {
	var result []string

	for _, numStr := range slice {
		// Split the string into individual digits
		digits := strings.Split(numStr, "")

		// Append each digit as a separate string
		result = append(result, digits...)
	}

	return result
}
