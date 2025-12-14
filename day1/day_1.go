package day1

import (
	"strconv"
)

func Rotations(rotations []string) (password int, err error) {
	password = 0
	var currentPosition = 50

	for _, rotation := range rotations {
		var direction = rotation[:1]
		var rotationAmount, _ = strconv.Atoi(rotation[1:])
		var nextPosition int

		// When value is 100 or more, get the number of times the dial passes through 0 immediately
		timesPassedZero := rotationAmount / 100
		remainderRotation := rotationAmount % 100

		if direction == "L" {
			nextPosition = currentPosition - remainderRotation
		} else {
			nextPosition = currentPosition + remainderRotation
		}

		// Dial is 0 to 99, more than 100 indicates it will have passed through 0
		if nextPosition >= 100 {
			currentPosition = nextPosition % 100
			timesPassedZero++
		} else if nextPosition < 0 {
			// Only increment when the starting position is not zero
			// Handles edge case where dial is at 0 currently, and L1 => stops at 99 without going past 0 during rotation.
			if currentPosition != 0 {
				timesPassedZero++
			}
			currentPosition = 100 + nextPosition
		} else {
			// Only increment when the starting position is not zero
			// Handles situation where dial is started at 0 and L1000 to prevent double counting
			if nextPosition == 0 && currentPosition != 0 {
				timesPassedZero++
			}
			currentPosition = nextPosition
		}

		password = password + timesPassedZero
	}

	return
}
