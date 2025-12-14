package day3_test

import (
	"aoc-2025/day3"
	"bufio"
	"log"
	"os"
	"testing"
)

func TestMaxJoltage_SingleRowSingleBattery_Expect1(t *testing.T) {
	rows := []string{"1"}
	expected := 1

	result := day3.MaxJoltage(rows)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMaxJoltage_SingleRowSameJoltage_Expect11(t *testing.T) {
	rows := []string{"111111"}
	expected := 11

	result := day3.MaxJoltage(rows)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMaxJoltage_SingleRowAscendingJoltage_Expect19(t *testing.T) {
	rows := []string{"123456789"}
	expected := 89

	result := day3.MaxJoltage(rows)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMaxJoltage_SingleRowDescendingJoltage_Expect98(t *testing.T) {
	rows := []string{"987654321"}
	expected := 98

	result := day3.MaxJoltage(rows)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMaxJoltage_SingleRowWeirdSequenceJoltage_Expect31(t *testing.T) {
	rows := []string{"1312"}
	expected := 32

	result := day3.MaxJoltage(rows)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMaxJoltage_MultiRowExample_Expect357(t *testing.T) {
	rows := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	expected := 357

	result := day3.MaxJoltage(rows)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSumOfInvalidIds_ReadFromFile(t *testing.T) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	var rowsOfBatteries []string
	for fileScanner.Scan() {
		rowsOfBatteries = append(rowsOfBatteries, fileScanner.Text())
	}

	expected := 16946

	result := day3.MaxJoltage(rowsOfBatteries)

	if result != expected || err != nil {
		t.Errorf("Expected %d, got %d and error %v", expected, result, err)
	}
}
