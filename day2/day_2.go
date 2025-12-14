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

	// If both start and end range are odd, return immediately as doubles are always even length
	if len(start) == len(end) && len(start)%2 == 1 {
		return
	}

	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)
	for i := startInt; i <= endInt; i++ {
		iString := strconv.Itoa(i)

		if len(iString)%2 == 0 {
			halvedLength := len(iString) / 2
			if iString[:halvedLength] == iString[halvedLength:] {
				sum += i
			}
		}
	}

	return
}
