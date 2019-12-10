package main

import "testing"

func TestHasAdjacentNumbersFalse(t *testing.T) {
	input := 10
	expectedOutput := false
	result := hasAdjacentNumbers(input)
	if (!(result == expectedOutput)) {
		t.Errorf(`HasAdjacentNumbers is not equal to %v, got %v`, expectedOutput, result)
	}	
}

func TestHasAdjacentNumbersTrue(t *testing.T) {
	input := 11
	expectedOutput := true
	result := hasAdjacentNumbers(input)
	if (!(result == expectedOutput)) {
		t.Errorf(`HasAdjacentNumbers is not equal to %v, got %v`, expectedOutput, result)
	}	
}

func TestDoDigitsAlwaysIncreaseFalse(t *testing.T) {
	input := 321
	expectedOutput := false
	result := doDigitsAlwaysIncrease(input)
	if (!(result == expectedOutput)) {
		t.Errorf(`doDigitsAlwaysIncrease is not equal to %v, got %v`, expectedOutput, result)
	}	
}

func TestDoDigitsAlwaysIncreaseTrue(t *testing.T) {
	input := 123
	expectedOutput := true
	result := doDigitsAlwaysIncrease(input)
	if (!(result == expectedOutput)) {
		t.Errorf(`doDigitsAlwaysIncrease is not equal to %v, got %v`, expectedOutput, result)
	}	
}

func TestDoDigitsAlwaysIncreaseWhenSameTrue(t *testing.T) {
	input := 111
	expectedOutput := true
	result := doDigitsAlwaysIncrease(input)
	if (!(result == expectedOutput)) {
		t.Errorf(`doDigitsAlwaysIncrease is not equal to %v, got %v`, expectedOutput, result)
	}	
}

func TestValidPasswordTrue(t *testing.T) {
	input := 234566
	expectedOutput := true
	result := isValidPassword(input)
	if (!(result == expectedOutput)) {
		t.Errorf(`isValidPassword is not equal to %v, got %v`, expectedOutput, result)
	}	
}

func TestValidPasswordTrueTwo(t *testing.T) {
	input := 111111
	expectedOutput := true
	result := isValidPassword(input)
	if (!(result == expectedOutput)) {
		t.Errorf(`isValidPassword is not equal to %v, got %v`, expectedOutput, result)
	}	
}
