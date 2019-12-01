package main

import "testing"

func TestCalculateFuelRequired(t *testing.T) {
	var result int 
	result = calculateFuelRequired(12)
	if (result != 2) {
		t.Errorf(`calculateFuelRequired(12) is not equal to %d, got %d`, 2, result)
	}	

	result = calculateFuelRequired(14)
	if (result != 2) {
		t.Errorf(`calculateFuelRequired(14) is not equal to %d, got %d`, 2, result)
	}	

	result = calculateFuelRequired(1969)
	if (result != 966) {
		t.Errorf(`calculateFuelRequired(1969) is not equal to %d, got %d`, 966, result)
	}	

	result = calculateFuelRequired(100756)
	if (result != 50346) {
		t.Errorf(`calculateFuelRequired(33583) is not equal to %d, got %d`, 50346, result)
	}	
}