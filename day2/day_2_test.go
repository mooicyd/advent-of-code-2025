package day2_test

import (
	"aoc-2025/day2"
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestSumOfInvalidIds_EmptyList_ExpectZero(t *testing.T) {
	var idList []string

	result, err := day2.SumOfInvalidIds(idList)

	if result != 0 || err != nil {
		t.Errorf("Expected 0, got %d and error %v", result, err)
	}
}

func TestSumOfInvalidIds_RangeBetween11And12_Expect11(t *testing.T) {
	idList := []string{"11-12"}

	result, err := day2.SumOfInvalidIds(idList)

	if result != 11 || err != nil {
		t.Errorf("Expected 11, got %d and error %v", result, err)
	}
}

func TestSumOfInvalidIds_RangeBetween11And22_Expect33(t *testing.T) {
	idList := []string{"11-22"}

	result, err := day2.SumOfInvalidIds(idList)

	if result != 33 || err != nil {
		t.Errorf("Expected 11, got %d and error %v", result, err)
	}
}

func TestSumOfInvalidIds_MultipleRanges_Expect132(t *testing.T) {
	idList := []string{"11-22", "99-115"}

	result, err := day2.SumOfInvalidIds(idList)

	if result != 132 || err != nil {
		t.Errorf("Expected 11, got %d and error %v", result, err)
	}
}

func TestSumOfInvalidIds_RangeBetween11And99_Expect495(t *testing.T) {
	idList := []string{"11-99"}

	result, err := day2.SumOfInvalidIds(idList)

	if result != 495 || err != nil {
		t.Errorf("Expected 495, got %d and error %v", result, err)
	}
}

func TestSumOfInvalidIds_RangesWithOddNumberOfDigits_ExpectZero(t *testing.T) {
	idList := []string{"11111-99999", "1-9", "111-999", "1111111-9999999", "111111111-999999999"}

	result, err := day2.SumOfInvalidIds(idList)

	if result != 0 || err != nil {
		t.Errorf("Expected %d, got %d and error %v", 0, result, err)
	}
}

func TestSumOfInvalidIds_RangeBetween27And47_Expect77(t *testing.T) {
	idList := []string{"27-47"}

	result, err := day2.SumOfInvalidIds(idList)

	if result != 77 || err != nil {
		t.Errorf("Expected %d, got %d and error %v", 77, result, err)
	}
}

func TestSumOfInvalidIds_RangeBetween1And99_Expect495(t *testing.T) {
	idList := []string{"1-99"}

	result, err := day2.SumOfInvalidIds(idList)

	if result != 495 || err != nil {
		t.Errorf("Expected 495, got %d and error %v", result, err)
	}
}

func TestSumOfInvalidIds_First3ElementsOfExample_Expect1142(t *testing.T) {
	idList := []string{"11-22", "99-115", "998-1015"}

	result, err := day2.SumOfInvalidIds(idList)

	if result != 1142 || err != nil {
		t.Errorf("Expected %d, got %d and error %v", 1142, result, err)
	}
}

func TestSumOfInvalidIds_ReadFromFile(t *testing.T) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	var idList []string
	for fileScanner.Scan() {
		idList = strings.Split(fileScanner.Text(), ",")
	}

	result, err := day2.SumOfInvalidIds(idList)

	if result != 15873079081 || err != nil {
		t.Errorf("Expected 5872, got %d and error %v", result, err)
	}
}
