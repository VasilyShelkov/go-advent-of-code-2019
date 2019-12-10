package main

import "fmt"

func hasAdjacentNumbers(number int) bool {
	previousDigit := number % 10
    for (previousDigit > 0) {
		previousDigit /= 10
		nextDigit := previousDigit % 10
		if (previousDigit == nextDigit) {
			return true
		}

		previousDigit = nextDigit
	}
	
    return false
}

func doDigitsAlwaysIncrease(number int) bool {
	var digits []int
	for (number > 0) {
		digits = append(digits, number % 10)
        number /= 10
	}

	// reverse the digit slice order
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	previousDigitToCheck := digits[0]
	for _, nextDigit := range digits[1:] {
		if (previousDigitToCheck > nextDigit) {
			return false
		}
		previousDigitToCheck = nextDigit
	}

	return true
}

func isValidPassword(passwordAttempt int) bool {
	if (hasAdjacentNumbers(passwordAttempt) && doDigitsAlwaysIncrease(passwordAttempt)) {
		return true
	}
	return false
}

func main() {
	minPassword := 307237
	maxPassowrd := 769058

	validpasswordCount := 0
	for i := minPassword;  i <= maxPassowrd; i++ {
		if (isValidPassword(i)) {
			fmt.Printf("%d...", i)
			validpasswordCount++
		}
	}
	fmt.Printf("Number of valid passwords: %d", validpasswordCount)
}