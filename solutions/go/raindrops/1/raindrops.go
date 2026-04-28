package raindrops

import "strconv"

func divisibleBy(dividend int, divider int) bool {
	return dividend%divider == 0
}

func Convert(number int) string {
	if !divisibleBy(number, 3) &&
		!divisibleBy(number, 5) &&
		!divisibleBy(number, 7) {
		return strconv.Itoa(number)
	}

	finalValue := ""

	if divisibleBy(number, 3) {
		finalValue += "Pling"
	}

	if divisibleBy(number, 5) {
		finalValue += "Plang"
	}

	if divisibleBy(number, 7) {
		finalValue += "Plong"
	}

	return finalValue
}