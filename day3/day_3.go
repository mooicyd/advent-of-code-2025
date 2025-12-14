package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func MaxJoltage(rowsOfBatteries []string) int {
	sum := 0

	for _, row := range rowsOfBatteries {
		var currentPair []string
		for _, j := range row {
			joltage := string(j)
			if len(currentPair) < 2 {
				currentPair = append(currentPair, joltage)
			} else {
				if currentPair[1] > currentPair[0] {
					currentPair[0] = currentPair[1]
					currentPair[1] = joltage
				} else if joltage > currentPair[1] {
					currentPair[1] = joltage
				}
			}
		}

		result, _ := strconv.Atoi(strings.Join(currentPair, ""))
		fmt.Println(result)
		sum += result
	}
	return sum
}
