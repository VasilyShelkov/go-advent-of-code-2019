package main

import "testing"

func TestIntcodeOptCode1(t *testing.T) {
	input := []int{1,0,0,0,99}
	expectedOutput := []int{2,0,0,0,99}
	result := calculateIntcode(input)
	if (!Equal(result, expectedOutput)) {
		t.Errorf(`calculateIntcode is not equal to %v, got %v`, expectedOutput, result)
	}	
}

func TestIntcodeOptCode2(t *testing.T) {
	input := []int{2,3,0,3,99}
	expectedOutput := []int{2,3,0,6,99}
	result := calculateIntcode(input)
	if (!Equal(result, expectedOutput)) {
		t.Errorf(`calculateIntcode is not equal to %v, got %v`, expectedOutput, result)
	}	
}

func TestIntcodeOptCodeReplaceHalt(t *testing.T) {
	input := []int{1,1,1,4,99,5,6,0,99}
	expectedOutput := []int{30,1,1,4,2,5,6,0,99}
	result := calculateIntcode(input)
	if (!Equal(result, expectedOutput)) {
		t.Errorf(`calculateIntcode is not equal to %v, got %v`, expectedOutput, result)
	}	
}

func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
	}

    return true
}