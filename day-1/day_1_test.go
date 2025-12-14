package main_test

import (
	Day_One "aoc-2025/day-1"
	"bufio"
	"log"
	"os"
	"testing"
)

func TestRotateLeftOnceToZero(t *testing.T) {
	rotations := []string{"L50"}

	result, err := Day_One.Rotations(rotations)

	if result != 1 || err != nil {
		t.Errorf("Expected 1, got %d and error %v", result, err)
	}
}

func TestRotateRightOnceToZero(t *testing.T) {
	rotations := []string{"R50"}

	result, err := Day_One.Rotations(rotations)

	if result != 1 || err != nil {
		t.Errorf("Expected 1, got %d and error %v", result, err)
	}
}

func TestRotateReachZeroTwice(t *testing.T) {
	rotations := []string{"R50", "L50", "R50"}

	result, err := Day_One.Rotations(rotations)

	if result != 2 || err != nil {
		t.Errorf("Expected 2, got %d and error %v", result, err)
	}
}

func TestRotateLeftAndRight(t *testing.T) {
	rotations := []string{"R39", "L61", "L28"}

	result, err := Day_One.Rotations(rotations)

	if result != 1 || err != nil {
		t.Errorf("Expected 1, got %d and error %v", result, err)
	}
}

func TestTenRoundsInOneGo(t *testing.T) {
	rotations := []string{"R50", "L1000"}

	result, err := Day_One.Rotations(rotations)

	if result != 11 || err != nil {
		t.Errorf("Expected 11, got %d and error %v", result, err)
	}
}

func TestExpectTenFromRotatingRight(t *testing.T) {
	rotations := []string{"R950"}
	result, err := Day_One.Rotations(rotations)

	if result != 10 || err != nil {
		t.Errorf("Expected 1, got %d and error %v", result, err)
	}
}

func TestExpectTenFromRotatingLeft(t *testing.T) {
	rotations := []string{"L950"}
	result, err := Day_One.Rotations(rotations)

	if result != 10 || err != nil {
		t.Errorf("Expected 1, got %d and error %v", result, err)
	}
}

// Reads from the AOC Day 1 input
func TestRotateFromFile(t *testing.T) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	var rotations []string
	for fileScanner.Scan() {
		rotations = append(rotations, fileScanner.Text())
	}

	result, err := Day_One.Rotations(rotations)

	if result != 5872 || err != nil {
		t.Errorf("Expected 5872, got %d and error %v", result, err)
	}
}
