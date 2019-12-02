package main

import "testing"

func TestOuputNounVerb(t *testing.T) {
	input := 4714701
	expectedOutput := 1202
	result := getOutput(input)
	if (result != expectedOutput) {
		t.Errorf(`calculateIntcode is not equal to %v, got %v`, expectedOutput, result)
	}
}