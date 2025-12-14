package day2

import (
	"strconv"
	"strings"
)

func SumOfInvalidIds(idRanges []string) (sum int, err error) {
	sum = 0
	for _, idRange := range idRanges {
		split := strings.Split(idRange, "-")
		sum = sum + sumOfInvalidIdInRange(split[0], split[1])
	}

	return
}

func sumOfInvalidIdInRange(start, end string) (sum int) {
	sum = 0
	set := make(map[string]struct{})

	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)
	for i := startInt; i <= endInt; i++ {
		iString := strconv.Itoa(i)

		if len(iString) == 1 {
			continue
		}

		maxSubStringLength := len(iString) / 2
		for j := maxSubStringLength; j >= 1; j-- {
			if hasRepeatedValues(iString, j) {
				set[iString] = struct{}{}
			}
		}

	}

	for element := range set {
		integer, _ := strconv.Atoi(element)
		sum += integer
	}

	return
}

func hasRepeatedValues(evenLengthInteger string, numOfCharacters int) bool {
	length := len(evenLengthInteger)
	remainder := length % numOfCharacters
	if remainder != 0 {
		return false
	}
	firstNCharacters := evenLengthInteger[:numOfCharacters]
	return evenLengthInteger == strings.Repeat(firstNCharacters, length/numOfCharacters)
}
