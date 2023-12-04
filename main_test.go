package main

import (
	"testing"
)

func TestAdvent1(test *testing.T) {
	// arrange
	expected := 55208

	// act
	result := advent1()
	// expect
	if result != expected {
		test.Errorf("got %v but expected %v", result, expected)
	}
}

func TestAdvent1Secnd(test *testing.T) {
	// arrange
	testInput := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expected := 276

	// act
	result := advent1Secnd(testInput)

	if result != expected {
		test.Errorf("got %v, but expect %v", result, expected)
	}
}
