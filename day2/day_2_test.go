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

func TestSumOfInvalidIds_MultipleRanges_Expect243(t *testing.T) {
	idList := []string{"11-22", "99-115"}

	result, err := day2.SumOfInvalidIds(idList)

	if result != 243 || err != nil {
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

	if result != 493939393380 || err != nil {
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
	expected := 1142 + 111 + 999

	result, err := day2.SumOfInvalidIds(idList)

	if result != expected || err != nil {
		t.Errorf("Expected %d, got %d and error %v", expected, result, err)
	}
}

func TestSumOfInvalidIds_Example(t *testing.T) {
	idList := []string{"11-22", "99-115", "998-1015", "1188511880-1188511890", "222220-222224", "1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659", "824824821-824824827", "2121212118-2121212124"}

	expected := 4174379265

	result, err := day2.SumOfInvalidIds(idList)

	if result != expected || err != nil {
		t.Errorf("Expected %d, got %d and error %v", expected, result, err)
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

	if result != 22617871034 || err != nil {
		t.Errorf("Expected 5872, got %d and error %v", result, err)
	}
}
