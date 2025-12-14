package main

import "strconv"

func main() {
	return
}

func Rotations(rotations []string) (password int, err error) {
	password = 0
	var currentPosition = 50

	for _, rotation := range rotations {
		var direction = rotation[:1]
		var rotationAmount, _ = strconv.Atoi(rotation[1:])

		if direction == "L" {
			currentPosition = (currentPosition - rotationAmount) % 100
		} else {
			currentPosition = (currentPosition + rotationAmount) % 100
		}

		if currentPosition == 0 {
			password++
		}
	}

	return
}
